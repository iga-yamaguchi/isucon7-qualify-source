.PHONY: mysql check-slow

mysql:
	sudo touch /var/log/mysql/slow_query.log
	sudo rm /var/log/mysql/slow_query.log
	sudo cp /home/isucon/isubata/files/db/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf
	sudo systemctl restart mysql

check-slow:
	sudo mysqldumpslow -s t /var/log/mysql/slow_query.log
