# Examen mutante

# Tecnologia

- Go go1.15.5
- echo v4
- Postgres13

## Setup

### Requirimientos

- go version go1.15.5 
- Postgres13
- docker

### Instalacion

- Clonar https://github.com/andrew221293/examenMutante.git
- Ejecutar comando `go get -v` para instalar dependencias
- levantar una instancia docker mediante un archivo `docker-compose.yml` ejemplo en el repositorio
- Ejecutuar comando en la ruta del archivo `docker-compose up -d`
- verificar que el servicio este activo `docker ps`
- Entrar a la instancia `docker exec -it <name instance> psql -U postgres -d postgres`
- Crear Database marvel `CREATE DATABASE marvel;`
- Crear User `create role <nameRol> with login password 'some passord';`
- Dar permisos de superUsuario `ALTER ROLE <nameUser> WITH SUPERUSER;`
- Dar todos los privilegios al usuario en la BD `GRANT ALL PRIVILEGES ON DATABASE "nameBD" to nameUser;`
- Levantar instnacia `go run main.go` y esta ejecutara los archivos de migracion para crear las tablas

## Uso

### Postman

#### Saber si es mutante o humano

- URL: http://localhost:8080/mutant
