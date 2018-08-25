```
su isucon

export PATH=$HOME/local/go/bin:$HOME/go/bin:$PATH
export GOPATH=/home/isucon/isubata/webapp/go

sudo systemctl stop nginx
sudo systemctl stop isubata.php.service

sudo rm /etc/nginx/sites-enabled/nginx.conf
sudo ln -s /etc/nginx/sites-available/nginx.conf /etc/nginx/sites-enabled/nginx.conf

cd /home/isucon/isubata/webapp/go
make build

sudo systemctl start isubata.golang.service
sudo systemctl start nginx
```
