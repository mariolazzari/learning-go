package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(t1 string, p1 int, t2 string, p2 int) {
	if p1 > p2 {
		l.Wins[t1]++
		l.Wins[t2] += 0
	} else {
		l.Wins[t2]++
	}
}

func (l League) Ranking() []string {
	keys := make([]string, 0, len(l.Teams))
	for k := range l.Wins {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return l.Wins[keys[i]] > l.Wins[keys[j]]
	})

	return keys
}

func main() {
	celtics := Team{
		Name: "Celtics",
	}
	lakers := Team{
		Name: "Lakers",
	}
	nuggets := Team{
		Name: "Nuggets",
	}
	teams := []Team{celtics, lakers, nuggets}

	nba := League{
		Teams: teams,
		Wins:  make(map[string]int, len(teams)),
	}

	nba.MatchResult(celtics.Name, 110, lakers.Name, 99)
	nba.MatchResult(celtics.Name, 100, nuggets.Name, 98)
	nba.MatchResult(lakers.Name, 80, nuggets.Name, 111)
	fmt.Println(nba.Ranking())

}
