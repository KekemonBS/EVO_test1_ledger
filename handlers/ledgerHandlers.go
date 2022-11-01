package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/KekemonBS/ledgerTest/models"
	"github.com/KekemonBS/ledgerTest/storage/csvimport"
)

type DbImpl interface {
	Create(ctx context.Context, tr models.Transaction) error
	Read(ctx context.Context, id int64) (models.Transaction, error)
	Delete(ctx context.Context, id int64) error
	Search(ctx context.Context, v url.Values) ([]models.Transaction, error)
}

type LedgerHandlers struct {
	ctx    context.Context
	logger *log.Logger
	dbImpl DbImpl
}

func New(ctx context.Context, logger *log.Logger, dbImpl DbImpl) *LedgerHandlers {
	return &LedgerHandlers{
		ctx:    ctx,
		logger: logger,
		dbImpl: dbImpl,
	}
}

func (l *LedgerHandlers) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	models, err := l.dbImpl.Search(l.ctx, query)
	if err != nil {
		l.logger.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(models)
	if err != nil {
		l.logger.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (l *LedgerHandlers) Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 << 26) //~134M mem limit
	f, _, err := r.FormFile("uploadedFile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	csvimport.ImportCSVFile(l.ctx, l.dbImpl, "UAH", f)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
