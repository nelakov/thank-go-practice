// Epic Battle
//
// Your task is to understand the code and win the greatest battle of our time.
//
// The task is not too serious: you can do nothing at all and win by the power
// of statistics and pressing "submit". Or you can slightly change the odds
// in your hero's favor.
//
// Sample Output:
//
//   Вы победили

package main

import "math/rand"

// how long the battle will last
const maxRounds = 500

// Monster size.
const (
	sizeSmall = iota
	sizeMedium
	sizeLarge
)

// Hero weapon.
const (
	weaponLightsaber = iota
	weaponSword
	weaponDagger
)

// Special traits of hero/monster.
const (
	traitLucky  = 0b0001
	traitFast   = 0b0010
	traitStrong = 0b0100
	traitRegen  = 0b1000
)

// Actor is a battle participant.
type Actor struct {
	name   string
	attp   int // attack
	defp   int // defense
	hp     int // health
	traits int
}

// Defend calculates damage to the actor,
// considering their defense and special traits.
func (a Actor) Defend(damage int) int {
	effectiveDamage := damage - a.defp
	if a.HasTrait(traitFast) {
		effectiveDamage -= 10
	}
	if a.HasTrait(traitLucky) && checkLuck() {
		effectiveDamage -= 10
	}
	if a.HasTrait(traitRegen) {
		effectiveDamage /= 2
	}
	if effectiveDamage < 0 {
		return 0
	}
	return effectiveDamage
}

// HasTrait checks if the actor has the specified trait.
func (a Actor) HasTrait(trait int) bool {
	return a.traits&trait != 0
}

// Hero is the player character.
type Hero struct {
	Actor
	weapon int
}

// Attack returns the damage dealt by the hero,
// considering weapon and special traits.
func (h Hero) Attack() int {
	damage := h.attp
	switch h.weapon {
	case weaponSword:
		damage += 5
	case weaponLightsaber:
		damage += 10
	}
	if h.HasTrait(traitStrong) {
		damage += 10
	}
	if h.HasTrait(traitLucky) && checkLuck() {
		damage += 10
	}
	return damage
}

// Monster is the enemy.
type Monster struct {
	Actor
	size int
}

// Attack returns the damage dealt by the monster,
// considering size and special traits.
func (m Monster) Attack() int {
	damage := m.attp
	switch m.size {
	case sizeMedium:
		damage += 5
	case sizeLarge:
		damage += 10
	}
	if m.HasTrait(traitStrong) {
		damage += 10
	}
	if m.HasTrait(traitLucky) && checkLuck() {
		damage += 10
	}
	return damage
}

// generateHero generates a random hero.
func generateHero(name string) Hero {
	return Hero{
		Actor: Actor{
			name:   name,
			attp:   (8 + rand.Intn(5)) * 20,
			defp:   (8 + rand.Intn(5)) * 10,
			hp:     (8 + rand.Intn(3)) * 40,
			traits: rand.Intn(16),
		},
		weapon: rand.Intn(3),
	}
}

// generateMonster generates a random monster.
func generateMonster(name string) Monster {
	return Monster{
		Actor: Actor{
			name:   name,
			attp:   (8 + rand.Intn(5)) * 20,
			defp:   (8 + rand.Intn(5)) * 10,
			hp:     (8 + rand.Intn(3)) * 40,
			traits: rand.Intn(16),
		},
		size: rand.Intn(3),
	}
}

// checkLuck checks luck.
func checkLuck() bool {
	return rand.Intn(2) == 0
}

// hidden game code follows
