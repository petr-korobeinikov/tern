#!/usr/bin/env bats

@test "tern bool mode 1" {
    res=$(tern -b 1 yes no)
    [ "$res" == "yes" ]
}

@test "tern bool mode true" {
    res=$(tern -b true yes no)
    [ "$res" == "yes" ]
}

@test "tern bool mode TRUE" {
    res=$(tern -b TRUE yes no)
    [ "$res" == "yes" ]
}

@test 'tern bool mode "TRUE"' {
    res=$(tern -b "TRUE" yes no)
    [ "$res" == "yes" ]
}

@test "tern bool mode 'true'" {
    res=$(tern -b 'true' yes no)
    [ "$res" == "yes" ]
}
