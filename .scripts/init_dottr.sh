#!/bin/sh

#
# This File will init everything required to test dotter
#

touch $HOME/.test-dottr-file
mkdir $HOME/.test-dottr
touch $HOME/.test-dottr/.gitkeep

mv $HOME/.dottr.yaml $HOME/.dottr.yaml.backup
cp ./.samples/config.yaml $HOME/.dottr.yaml
