# article-api

# Setup project and run the project

1. clone this repo
2. create database with postgreSQL
3. create file .env for setup database
   example:
   HOST=localhost
   PORTDB=5432
   DBUSER=postgres
   DBPASSWORD=postgres
   DBNAME=article-app
   Or you can see in file .env-template
4. run go mod tidy in terminal
5. go run main.go / make nrun (use nodemon)

NOTE: Table automatically created because use gorm automigrate

# Documentasi Api

create article
http://localhost:8080/articles (POST)

request body :
{
"author":"Yusuke Murata",
"title":"Mantap Men 2",
"body":"1 pukulan menghajar semuanya"
}

get all article
http://localhost:8080/articles (GET)
query param:
search (string)
author (string)

get one article by id
http://localhost:8080/articles/2 (GET)

delete articel by id
http://localhost:8080/articles/5 (DELETE)
(this will be soft delete)

update article
http://localhost:8080/articles/2 (PATCH)
request body :
{
"author":"Yusuke Murata",
"title":"Mantap Men 2",
"body":"1 pukulan menghajar semuanya"
}
