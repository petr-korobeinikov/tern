#!/usr/bin/env bats

@test "tern shell mode emtpy string" {
    res=$(tern "" yes no)
    [ "$res" == "no" ]
}

@test "tern shell mode unset var" {
    res=$(tern "${UNSET_VAR}" yes no)
    [ "$res" == "no" ]
}
