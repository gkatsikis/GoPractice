package main

import {
	"math/rand"
	"time"
}

func shuffle(d deck) deck {
	rand.Seed(time.Now().UnixNano())

	numberOfCards := len(d)

	rand.Shuffle(numberOfCards, func(i, j int)) {
		d[i], d[j] = d[j], d[i]
	}

	return d
}