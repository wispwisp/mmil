#!/bin/bash

docker run --rm -it -v$(pwd -P)/src:/home/user/src golang:1.19 /bin/bash
