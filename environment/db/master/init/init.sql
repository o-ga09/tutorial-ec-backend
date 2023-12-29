CREATE USER 'repl'@'api-dbsrv01' IDENTIFIED BY 'repl';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'api-dbsrv01';