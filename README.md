# Library REST API

This project has the goal to become  more familiar with the Go language. <br>
It has two major components: the **Database** and the **REST API** interface. The database is a relational one (MySQL). It stores the books added via the APIs. The APIs act as an interface between the client and the database and take care of data validation as well as adding, deleting and retrieving them.

### General operating scheme
![diagram](./img/diagram.svg)

### Requirements

- Go
- MySQL

### MySQL Setup
1. Create a database
```sql
CREATE DATABASE library;
````
To make sure the database has been created you can execute this command:
```sql
SHOW DATABASE;
```
The output should be similar to this and should contain **library**:
```
+--------------------+
| Database           |
+--------------------+
| information_schema |
| library            |
| mysql              |
| performance_schema |
| recordings         |
| sys                |
+--------------------+
```
Once the database is created select it using the following command:
```sql
USE library;
```
2. Create a table
```sql
CREATE TABLE book(
    isbn VARCHAR(17) PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    author VARCHAR(128) NOT NULL,
    summary TEXT NOT NULL,
    pub_year INT NOT NULL
)
```
To make sure that you created the table you can execute this command:
```sql
DESCRIBE book;
```
The output should be the following:
```
+----------+--------------+------+-----+---------+-------+
| Field    | Type         | Null | Key | Default | Extra |
+----------+--------------+------+-----+---------+-------+
| isbn     | varchar(17)  | NO   | PRI | NULL    |       |
| title    | varchar(256) | NO   |     | NULL    |       |
| author   | varchar(128) | NO   |     | NULL    |       |
| summary  | text         | NO   |     | NULL    |       |
| pub_year | int          | YES  |     | NULL    |       |
+----------+--------------+------+-----+---------+-------+
```

3. Populate the table (optional):
If you want to have same record you can execute the following command in you sql shell:
```sql
INSERT INTO book values (
    "978-0-452-26756-5",
    "Mastery: The Keys to Success and Long-Term Fulfillment",
    "George Leonard",
    "Whether you're seeking to improve your career or your intimate relationships, increase self-esteem or create harmony within yourself, this inspiring prescriptive guide will help you master anything you choose and achieve success in all areas of your life.",
    1992);
```

```sql
INSERT INTO book values (
    "978-88-541-8089-5",
    "Le avventure di Sherlock Holmes",
    "Conan Doyle",
    "Il più popolare detective di tutta la letteratua mondiale.",
    1991);
```

### Go enviroment
The source files for Go are placed in the **src** folder. It's composed by two modules:
1. databse_manager
2. handler_api

Before running the application you need to setup some environment variables:
1. DBUSER
2. DBPASS

In Unix/MacOS you can do that by:
```bash
export DBUSER=user
export DBPASS=password
```

To start the server you have to run in the root folder
```bash
go run .
```

Make sure to download all the dependencies with
```bash
go get .
```

# API Endpoints
The API endpoint are accessible without any authentication. They're organized on the basis of the data on which they operate

## Get all books
Return the list of all books stored in the database

**URL**: `/books`
**Method**: `GET`
**Auth required** : No
**Permissions required** : None

### Success Response

**Code** : `200 OK`

**Response example**

```json
[
    {
		"isbn": "978-88-541-8089-5",
		"title": "Le avventure di Sherlock Holmes",
		"author": "Conan Doyle",
		"summary": "Il più popolare detective di tutta la letteratua mondiale.",
		"pub_year": 1991
	},
	{
		"isbn": "978-0452267565",
		"title": "Mastery: The Keys to Success and Long-Term Fulfillment",
		"author": "George Leonard",
		"summary": "Whether you're seeking to improve your career or your intimate relationships, increase self-esteem or create harmony within yourself, this inspiring prescriptive guide will help you master anything you choose and achieve success in all areas of your life.",
		"pub_year": 1992
	}
]
```

## Get a book by its ISBN
Return the book with a specified ISBN

**URL**: `/books/isbn`
**Method**: `GET`
**Auth required** : No
**Permissions required** : None

### Success Response

**URL example**: `/books/978-88-541-8089-5`
**Code** : `200 OK`
**Response example**

```json
    {
		"isbn": "978-88-541-8089-5",
		"title": "Le avventure di Sherlock Holmes",
		"author": "Conan Doyle",
		"summary": "Il più popolare detective di tutta la letteratua mondiale.",
		"pub_year": 1991
	}
```

### ISBN not found

**Code** : `404 Not Found`

**Response**
```json
{
	"message": "There is no book with this ISBN"
}
```

## Add a book
Add a book to the library

**URL**: `/books`
**Method**: `POST`
**Auth required** : No
**Permissions required** : None

### Success Response

**Code** : `201 Created`
**Body example**

```json
    {
		"isbn": "978-88-541-8089-5",
		"title": "Le avventure di Sherlock Holmes",
		"author": "Conan Doyle",
		"summary": "Il più popolare detective di tutta la letteratua mondiale.",
		"pub_year": 1991
	}
```

### Add a book with an already taken ISBN

**Code** : `400 Bad Request`
**Response**
```json
{
	"message": "Duplicate entry"
}
```

### Missing required fields

**Code** : `400 Bad Request`
**Response**

```json
{
	"message": "make sure to fill of required field as well as to use the correct type"
}
```

## Delete a book
Delete a book which was previously saved

**URL**: `/books/isbn`
**Method**: `DELETE`
**Auth required** : No
**Permissions required** : None

### Success Response

**Code**: `204 No Content`
**No body is returned**

## Get all authors
Return the list of all authors

**URL**: `/authors`
**Method**: `GET`
**Auth required** : No
**Permissions required** : None

### Success Response

**Code** : `200 OK`
**Response example**

```json
[
	"Conan Doyle",
	"George Leonard"
]
```

## Get all books written by a specified author
Return the library filtered to match the specified author

**URL**: `/authors/name` (the name must be formatted to be a valid URL)
**Method**: `GET`
**Auth required** : No
**Permissions required** : None

### Success Response

**URL example**: `/authors/Conan%20Doyle`
**Code** : `200 OK`
**Response example**

```json
[
	{
		"isbn": "978-88-541-809",
		"title": "Le avventure di Sherlock Holmes",
		"author": "Conan Dosyle",
		"summary": "This is a tes",
		"pub_year": 1991
	}
]
```
## In the **src** folder there is a [file](./rsc/LibraryAPI_insomnia.json) which can be imported in Insomia
