package main

import (
	"errors"
	"fmt"
	"os"
)

// label is a unique identifier
type label string

// command is a command that can be executed in the game
type command label

// list of available commands
var (
	eat  = command("eat")
	take = command("take")
	talk = command("talk to")
)

// thing is an object that exists in the game
type thing struct {
	name    label
	actions map[command]string
}

// supports returns true if the object
// supports the command action
func (t thing) supports(action command) bool {
	_, ok := t.actions[action]
	return ok
}

// String returns a description of the object
func (t thing) String() string {
	return string(t.name)
}

// full list of objects in the game
var (
	apple = thing{"apple", map[command]string{
		eat:  "mmm, delicious!",
		take: "you have an apple now",
	}}
	bob = thing{"bob", map[command]string{
		talk: "Bob says hello",
	}}
	coin = thing{"coin", map[command]string{
		take: "you have a coin now",
	}}
	mirror = thing{"mirror", map[command]string{
		take: "you have a mirror now",
		talk: "mirror does not answer",
	}}
	mushroom = thing{"mushroom", map[command]string{
		eat:  "tastes funny",
		take: "you have a mushroom now",
	}}
)

// step describes a game step: a combination of command and object
type step struct {
	cmd command
	obj thing
}

// isValid returns true if the object
// is compatible with the command
func (s step) isValid() bool {
	return s.obj.supports(s.cmd)
}

// String returns a description of the step
func (s step) String() string {
	return fmt.Sprintf("%s %s", s.cmd, s.obj)
}

// begin solution

// invalidStepError occurs when a step command
// is not compatible with the object
type invalidStepError struct {
	st step
}

func (stErr invalidStepError) Error() string {
	return fmt.Sprintf("cannot %s", stErr.st)
}

// notEnoughObjectsError occurs when the game
// has run out of objects of a certain type
type notEnoughObjectsError struct {
	obj thing
}

func (objErr notEnoughObjectsError) Error() string {
	return fmt.Sprintf("there are no %ss left", objErr.obj)
}

// commandLimitExceededError occurs when a player
// has exceeded the limit for executing a command
type commandLimitExceededError struct {
	cmd command
}

func (cmdErr commandLimitExceededError) Error() string {
	switch cmdErr.cmd {
	case eat:
		return "you don't want to eat anymore"
	case talk:
		return "you don't want to talk anymore"
	default:
		return ""
	}
}

// objectLimitExceededError occurs when a player
// has exceeded the limit on the number of objects
// of a certain type in their inventory
type objectLimitExceededError struct {
	obj   thing
	limit int
}

func (objErr objectLimitExceededError) Error() string {
	return fmt.Sprintf("you already have a %s", objErr.obj)
}

// gameOverError is an error that occurred in the game
type gameOverError struct {
	// number of steps successfully completed
	// before the error occurred
	nSteps int
	err    error
}

func (g gameOverError) Error() string {
	return g.err.Error()
}

func (g gameOverError) Unwrap() error {
	return g.err
}

// player represents a game player
type player struct {
	// number of items eaten
	nEaten int
	// number of dialogs
	nDialogs int
	// inventory
	inventory []thing
}

// has returns true if the player
// has the object obj in their inventory
func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

// do executes command cmd on object obj
// on behalf of the player
func (p *player) do(cmd command, obj thing) error {
	// act according to the command
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{
				cmd: eat,
			}
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{
				obj:   obj,
				limit: 1,
			}
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{
				cmd: talk,
			}
		}
		p.nDialogs++
	}
	return nil
}

// newPlayer creates a new player
func newPlayer() *player {
	return &player{inventory: []thing{}}
}

// game describes the game
type game struct {
	// player
	player *player
	// objects in the game world
	things map[label]int
	// number of successfully completed steps
	nSteps int
}

// has checks whether the specified objects remain in the game world
func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

// execute performs the step st
func (g *game) execute(st step) error {
	// check command and object compatibility
	if !st.isValid() {
		return gameOverError{
			nSteps: g.nSteps,
			err:    invalidStepError{st: st},
		}
	}

	// when the player takes or eats an object,
	// it disappears from the game world
	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return gameOverError{
				nSteps: g.nSteps,
				err:    notEnoughObjectsError{obj: st.obj},
			}
		}
		g.things[st.obj.name]--
	}

	// execute the command on behalf of the player
	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{
			nSteps: g.nSteps,
			err:    err,
		}
	}

	g.nSteps++
	return nil
}

// newGame creates a new game
func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

// giveAdvice returns advice that will help
// the player avoid the error err in the future
func giveAdvice(err error) string {
	if stErr, ok := errors.AsType[invalidStepError](err); ok {
		return fmt.Sprintf("things like '%s %s' are impossible", stErr.st.cmd, stErr.st.obj)
	}

	if objErr, ok := errors.AsType[notEnoughObjectsError](err); ok {
		return fmt.Sprintf("be careful with scarce %ss", objErr.obj)
	}

	if cmdLimitErr, ok := errors.AsType[commandLimitExceededError](err); ok {
		switch cmdLimitErr.cmd {
		case eat:
			return "eat less"
		case talk:
			return "talk to less"
		}
	}

	if objLimitErr, ok := errors.AsType[objectLimitExceededError](err); ok {
		return fmt.Sprintf("don't be greedy, %d %s is enough", objLimitErr.limit, objLimitErr.obj)
	}

	return ""
}

// end solution

func main() {
	gm := newGame()
	steps := []step{
		{eat, apple},
		{talk, bob},
		{take, coin},
		{eat, mushroom},
	}

	for _, st := range steps {
		if err := tryStep(gm, st); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("You win!")
}

// tryStep executes a game step and prints the result
func tryStep(gm *game, st step) error {
	fmt.Printf("trying to %s %s... ", st.cmd, st.obj.name)
	if err := gm.execute(st); err != nil {
		fmt.Println("FAIL")
		return err
	}
	fmt.Println("OK")
	return nil
}
