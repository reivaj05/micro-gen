#!/bin/bash

success_status_code=0
failure_status_code=1

echo 'Running unit tests'
nosetests -c nose.cfg

if [ $? -eq $success_status_code ]
    then
        echo "The tests ran successfully"
    else
        echo "The tests failed, please fix them to continue"
        exit $failure_status_code
fi

exit $success_status_code