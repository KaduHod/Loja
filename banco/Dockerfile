# Usar a imagem oficial do PostgreSQL
FROM postgres:latest

# Definir variáveis de ambiente para o usuário, senha e nome do banco de dados
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_DB=loja
COPY ./dumps/dump.sql /dump.sql
RUN mkdir /var/log/banco
RUN touch /var/log/banco/postgresql.log
RUN chmod -R 777 /var/log/banco
EXPOSE 5432

CMD ["postgres", "-c", "logging_collector=on", "-c", "log_directory=/var/log/banco", "-c", "log_filename=postgresql.log", "-c", "log_statement=all"]

