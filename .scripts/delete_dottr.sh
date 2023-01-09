#!/bin/sh

#
# This will cleanup dottr files after you're done with development
#

rm $HOME/.test-dottr-file
rm $HOME/.dottr.yaml

rm -rf $HOME/.test-dottr/

mv $HOME/.dottr.yaml.backup $HOME/.dottr.yaml
