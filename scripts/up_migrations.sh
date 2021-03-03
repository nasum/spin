#!/bin/bash
migrate -database ${DATABASE_URL} -path migrations up $1
