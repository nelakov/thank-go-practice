// Exercise: Football Tournament
//
// Four teams (A, B, C, D) play a football tournament. Each team plays every
// other team exactly once (6 matches total). Calculate total points per team.
//
// Scoring: Win = 3 points, Draw = 1 point, Loss = 0 points.
//
// Input: a single line of space-separated triplets. Each triplet is three
// characters: 1st = team1, 2nd = team2, 3rd = result.
//   W — team1 wins
//   L — team1 loses
//   D — draw
//
// Output: teams with their total points, sorted alphabetically, space-separated.
// Format: "<Team><Points>" (e.g. "A6 B3 C1 D7").
//
// Sample Input:  ABW DCD DAW CBL BDL ACW
// Sample Output: A6 B3 C1 D7

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// result represents a match outcome
type result byte

// possible match outcomes
const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

// team represents a team
type team byte

// match represents a single match
// contains three fields:
// - first (first team)
// - second (second team)
// - result (match outcome)
// for example, the string BAW corresponds to
// first=B, second=A, result=W
type match struct {
	first  team
	second team
	result result
}

// rating represents the tournament standings —
// total points scored by each team
type rating map[team]int

// tournament represents a tournament
type tournament []match

// calcRating calculates and returns the tournament rating
func (t *tournament) calcRating() rating {
	r := rating{}
	// NOTE: += 0 lines don't change the score, but they register the team
	// as a key in the map. Without them, a team that only lost and never
	// drew or won would never appear in the rating.
	//
	// In Go, accessing a missing map key returns the zero value (0 for int),
	// but does NOT create the key — so the team would be silently missing from output.
	for _, m := range *t {
		switch m.result {
		case win:
			r[m.first] += 3
			r[m.second] += 0
		case draw:
			r[m.first]++
			r[m.second]++
		case loss:
			r[m.first] += 0
			r[m.second] += 3
		}
	}
	return r
}

// ┌──────────────────────────────────────┐
// │ do not change the code below         │
// └──────────────────────────────────────┘

// the code that parses match results is already implemented
// the code that prints the rating is already implemented
// your task is to implement the missing structs and methods above
func main() {
	src := readString()
	trn := parseTournament(src)
	rt := trn.calcRating()
	rt.print()
}

// readString reads and returns a string from os.Stdin
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

// parseTournament parses a tournament from the source string
func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

// parseMatch parses a match from a fragment of the source string
func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

// addMatch adds a match to the tournament
func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

// print outputs the tournament results
// in the format Aw Bx Cy Dz
func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}
