#!/bin/bash

echo $(curl -s -L -I -o /dev/null -w '%{url_effective}' https://en.wikipedia.org/wiki/Special:Random)