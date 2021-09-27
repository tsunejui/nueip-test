FROM postgres:11.5

# refs: https://github.com/mrts/docker-postgresql-multiple-databases
# usage:
#   docker build --tag=pgdb .
COPY create-multiple-postgresql-databases.sh /docker-entrypoint-initdb.d/
