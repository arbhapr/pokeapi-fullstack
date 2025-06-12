package model

import "github.com/google/uuid"

type NicknameRequest struct {
	Nickname string `json:"nickname" validate:"required"`
}

// CaughtPokemon represents a Pokémon that has been caught, possibly with a nickname
type CaughtPokemon struct {
	ID       uuid.UUID     `json:"id"`
	Pokemon  PokemonDetail `json:"pokemon"`
	Nickname string        `json:"nickname,omitempty"`
}

// MyPokemonList holds the list of caught Pokémon
var MyPokemonList []CaughtPokemon
