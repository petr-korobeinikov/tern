#!/usr/bin/env bats

@test "tern bool mode 0" {
    res=$(tern -b 0 yes no)
    [ "$res" == "no" ]
}

@test "tern bool mode false" {
    res=$(tern -b false yes no)
    [ "$res" == "no" ]
}

@test "tern bool mode FALSE" {
    res=$(tern -b FALSE yes no)
    [ "$res" == "no" ]
}

@test 'tern bool mode "false"' {
    res=$(tern -b "false" yes no)
    [ "$res" == "no" ]
}

@test "tern bool mode 'false'" {
    res=$(tern -b 'false' yes no)
    [ "$res" == "no" ]
}
