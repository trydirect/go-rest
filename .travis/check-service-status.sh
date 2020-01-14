#!/usr/bin/env bash

STATUS=$(curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/swagger/index.html)
[ $STATUS -eq 200 ]
