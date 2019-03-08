#!/bin/bash

declare GREEN='\033[0;32m'
declare NC='\033[0m'

print () {
    echo -e "${GREEN}$1${NC}"
}

test_race_conditions() {
    print "Running tests with the race condition detector"
    
    test_targets=$(go list ./...)
    env GORACE="history_size=7 halt_on_error=1" go test -count=1 -v -p 1 --test.parallel 10 -race $test_targets
}

test_race_conditions
