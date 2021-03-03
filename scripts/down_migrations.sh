#!/bin/bash
migrate -database ${DATABASE_URL} -path migrations down $1
