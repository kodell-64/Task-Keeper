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
├── mysql
│   └── db_export.txt
```

## Setup

Create a mysql database and user. The code currently expects a 'jgdb' database name and 'jg' database username. Default password used is 'jgpass'. Change as desired.

Import mysql/db_export.txt into your mysql db with mysql -u jg jgdb -p < mysql/db_export.txt

Build the Golang binary with 'go build process.go'. This will produce the runnable binary named 'process'.

If you do not have the gorilla-mux or mysql driver packages installed, install them per your O/S. Debian/Ubuntu uses the follow aptitude commands for installation:

'apt-get install golang-github-gorilla-mux-dev'

'apt-get install golang-github-go-sql-driver-mysql-dev'

## Run

Start the application by executing './process' within your shell. Browse to http://localhost:8000 to begin using Task Keeper.

## Usage

One can create, edit, mark-as-complete and delete tasks.

## Testing

One can run a small suite of unit tests against the service by executing 'go test'. Have a look at 'process_test.go' to review the tests that have been implemented.