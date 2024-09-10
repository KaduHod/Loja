#!/bin/sh

while true
do
    # Executa o logrotate com a configuração para o PostgreSQL e define um local acessível para o arquivo de estado
    logrotate -f /etc/logrotate.d/postgresql -s /var/log/banco/logrotate.status
    # Espera 24 horas (86400 segundos) antes de rodar novamente
    sleep 86400
done
