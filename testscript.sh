#!/bin/bash

while true
do 
  expr $RANDOM % 20 + 1
  sleep $(($RANDOM%10 / 5))
done
