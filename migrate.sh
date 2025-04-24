#!/bin/bash

. ./build/env.sh

# Create an SSH tunnel
# ssh -f -N -L 5433:localhost:5432 alexander0200@10.10.8.98 -p 2222

# Migrate the database
# To install migrate cli read this post https://www.freecodecamp.org/news/database-migration-golang-migrate/
#
# Usage:
#   migrate.sh [-up] [-down] [-drop] [-create <name>]
#
# Options:
#   -up
#       Move the database up to the latest version.
#
#   -down
#       Move the database down to the previous version.
#
#   -drop
#       Drop the database and all of its tables.
#
#   -create <name>
#       Create a new migration file with the given name.
#
#   -fix <version>
#       Fix the migration version to the specified version.
#
# Arguments:
#   None

while [[ "$#" -gt 0 ]]; do
  case "$1" in
  -up)
    # Move the database up to the latest version.
    migrate -path migrations -database "${DATABASE_URL}" up
    shift
    ;;
  -down)
    # Move the database down to the previous version.
    yes | migrate -path migrations -database "${DATABASE_URL}" down
    shift
    ;;
  -drop)
    # Drop the database and all of its tables.
    yes | migrate -path migrations -database "${DATABASE_URL}" drop
    shift
    ;;
  -create)
    # Create a new migration file with the given name.
    if [ -n "$2" ]; then
      migrate create -ext sql -dir migrations "$2"
      shift 2
    else
      echo "Error: Missing migration name." >&2
      exit 1
    fi
    ;;
  -fix)
    # Fix the migration version to the specified version.
    if [ -n "$2" ]; then
      migrate -path migrations -database "${DATABASE_URL}" force "$2"
      shift 2
    else
      echo "Error: Missing version number for fix." >&2
      exit 1
    fi
    ;;
  *)
    # Invalid option
    echo "Invalid option: $1" >&2
    exit 1
    ;;
  esac
done
