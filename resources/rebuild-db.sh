#!/bin/bash
set -e

run() {
  rm -f ./db/todo-db.sqlt
  sqlite3 ./db/todo-db.sqlt < ./resources/sql/schema.sql
}

$*
