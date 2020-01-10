# Task-Keeper

# Version

Phase 2 release.

Task Keeper allows one to simply create and delete tasks.

Uses gorilla-mux for routing and mysql-driver for RDBMS connectivity.
Routed, API service provided by process.go written in Golang.
Release also contains unit tests in process_test.go - usage instructions below.

## File Structure

The overall file structure is as follows:

```text
/Task-Keeper/
├── public_html
│   └── index.html
│   └── process.go
│   └── process_test.go
```

## Setup

Create a mysql database and user. The code currently expects a 'jgdb' database name and 'jg' database username. Default password used is 'jgpass'. Change as desired.

Import mysql/db_export.txt into your mysql db with mysql -u jg jgdb -p < mysql/db_export.txt

Build the Golang binary with 'go build process.go'. This will produce the runnable binary named 'process'.

## Run

Start the application by executing './process' within your shell. Browse to http://localhost:8000 to begin using Task Keeper.

## Usage

One can create, edit, mark-as-complete and delete tasks.

