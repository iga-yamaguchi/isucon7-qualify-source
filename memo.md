```
sudo ip route del default
sudo ip route add default via 10.0.0.1 dev eth0

sudo chown -R ubuntu:ubuntu /home/isucon/

echo export PATH=/home/isucon/local/go/bin:\$HOME/go/bin:\$PATH >> ~/.bashrc 
echo export GOPATH=/home/isucon/isubata/webapp/go >> ~/.bashrc
source ~/.bashrc

sudo systemctl stop nginx
sudo systemctl stop isubata.php.service



sudo rm /etc/nginx/sites-enabled/nginx.conf
sudo ln -s /etc/nginx/sites-available/nginx.conf /etc/nginx/sites-enabled/nginx.conf

cd /home/isucon/isubata/webapp/go
make build

sudo systemctl start isubata.golang.service
sudo systemctl start nginx

sudo apt install htop
sudo apt install dstat

sudo touch /var/log/mysql/slow_query.log

dstat -tafm

cd /home/isucon/isubata
git remote set-url origin git@github.com:iga-yamaguchi/isucon7-qualify-source.git

git fetch
git checkout develop

```
