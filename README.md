# Task-Keeper

# Version

Phase 2 release.

Task Keeper allows one to simply create and delete tasks.

Uses gorilla-mux for routing and mysql-driver for RDMS connectivity.
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

Create a mysql database and user. The code currently expects a 'jg' database name and databse username. Change as you wish.
Import mysql/db_export.txt into your mysql db with mysql -u jg jg -p < mysql/db_export.txt
Build the Golang binary with 'go build process.go'. This will produce the runnable binary named 'process'.

## Run

Start the Golang service by running '
The application can be located at http://localhost/<path_to_index.html>.

