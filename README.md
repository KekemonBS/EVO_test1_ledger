# ledger
---
## Install and run
### Arch
    # pacman -Sy --noconfirm docker docker-compose make go
    $ go mod vendor
    $ make up

### Debian
    # apt install -y docker docker-compose make
    $ wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz
    $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
    $ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
    $ go mod vendor
    $ make up
    
### Void
    # xbps-install -Sy docker docker-compose make go
    $ go mod vendor
    $ make up

---
## Usage
    $ xdg-open http://localhost:8080/ 

## Open Endpoints
* [Login](q.md) : `GET /q?transactionid=&terminalid=&status=&paymenttype=&datepost=&paymentnarrative=коштів&quantity=10&pagenum=1`
* [Login](u.md) : `POST /u`
