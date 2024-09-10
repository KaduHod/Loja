
FROM postgres:latest

# Definir variáveis de ambiente para o usuário, senha e nome do banco de dados
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_DB=loja

# Instalar logrotate
USER root
RUN apt-get update && apt-get install -y logrotate

# Criar o diretório de logs
RUN mkdir -p /var/log/banco && \
    touch /var/log/banco/postgresql.log && \
    chown -R postgres:postgres /var/log/banco && \
    chmod -R 755 /var/log/banco

RUN mkdir -p /var/lib/logrotate && \
    chown postgres:postgres /var/lib/logrotate && \
    chmod 755 /var/lib/logrotate

# Adicionar arquivo de configuração do logrotate
COPY logrotate.conf /etc/logrotate.d/postgresql
RUN chmod 644 /etc/logrotate.d/postgresql
COPY ./dumps/dump.sql /docker-entrypoint-initdb.d/dump.sql
EXPOSE 5432
CMD ["postgres", "-c", "logging_collector=on", "-c", "log_directory=/var/log/banco", "-c", "log_filename=postgresql.log", "-c", "log_statement=all"]

