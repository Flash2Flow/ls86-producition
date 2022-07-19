package main

type Error struct {
	NotFound         error
	TokenBad         error
	LoginAlreadyHave error
	EmailAlreadyHave error
}
