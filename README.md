## Prueba Tecnica - Go HTTP API - Joseph Caceres

### Requisitos

- Go v1.15+
- MySQL (see below).

### Contenido

Este proyecto se realizo con la finalidad de desarrollar la deteccion de mutantes por medio del adn 
(prueba tecnica mercado libre)

### Uso

Para ejecutar la aplicacion, simplemente ejecute
To execute the application from any lesson, just run:

    go run xmen-mutant/cmd/api/main.go 

#### MySQL & Docker

para simplificar la ejecución, se agrega un
Archivo `docker-compose.yaml` con un contenedor MySQL ya configurado.

Para ejecutarlo:

```sh
docker-compose up -d 
```
Es posible utilizar su propia instancia de MYSQL disponible ya sea
`localhost:3306`, con un nombre de base de datos llamada `persons`

Para configurar su base de datos, puede ejecutar el archivo `schema.sql`
presente en el directorio `sql`. Se carga automáticamente si
utiliza el archivo `docker-compose.yaml` proporcionado

#### Pruebas

Para ejecutar todas las pruebas, simplemente ejecute:

```sh
go test ./... 
```
