.PHONY: mysql check-slow

init:
    sudo ip route del default
    sudo ip route add default via 10.0.0.1 dev eth0

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

    sudo apt install htop
    sudo apt install dstat

mysql:
	sudo rm /var/log/mysql/slow_query.log
	sudo cp /home/isucon/isubata/files/db/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf
	sudo systemctl restart mysql

check-slow:
	sudo mysqldumpslow -s t /var/log/mysql/slow_query.log
