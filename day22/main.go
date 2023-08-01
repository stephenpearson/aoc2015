package main

import "fmt"

type Game struct {
	player_hp      int
	player_mana    int
	boss_hp        int
	boss_damage    int
	shield_turns   int
	poison_turns   int
	recharge_turns int
	mana_spent     int
}

func (g *Game) apply_effects() {
	if g.poison_turns > 0 {
		g.boss_hp -= 3
		g.poison_turns -= 1
	}
	if g.recharge_turns > 0 {
		g.player_mana += 101
		g.recharge_turns -= 1
	}
	if g.shield_turns > 0 {
		g.shield_turns -= 1
	}
}

func (g *Game) boss_attack() {
	dam := g.boss_damage
	if g.shield_turns > 0 {
		dam = dam - 7
	}
	if dam < 1 {
		dam = 1
	}
	g.player_hp -= dam
}

func (g *Game) player_attack(spell int) bool {
	switch spell {
	case 0:
		if g.player_mana >= 53 {
			g.player_mana -= 53
			g.mana_spent += 53
			g.boss_hp -= 4
			return true
		}
	case 1:
		if g.player_mana >= 73 {
			g.player_mana -= 73
			g.mana_spent += 73
			g.player_hp += 2
			g.boss_hp -= 2
			return true
		}
	case 2:
		if g.player_mana >= 113 {
			g.player_mana -= 113
			g.mana_spent += 113
			g.shield_turns = 6
			return true
		}
	case 3:
		if g.player_mana >= 173 {
			g.player_mana -= 173
			g.mana_spent += 173
			g.poison_turns = 6
			return true
		}
	case 4:
		if g.player_mana >= 229 {
			g.player_mana -= 229
			g.mana_spent += 229
			g.recharge_turns = 5
			return true
		}
	}
	return false
}

func gameSearch(games []Game, hard bool) int {
	minMana := 0
	for len(games) > 0 {
		newGames := []Game{}
		for _, g := range games {
			for i := 0; i < 5; i++ {
				cur := g
				if hard {
					cur.player_hp -= 1
					if cur.player_hp <= 0 {
						// player dead
						continue
					}
				}
				cur.apply_effects()
				if !cur.player_attack(i) {
					continue
				}

				cur.apply_effects()
				if cur.boss_hp <= 0 {
					// boss dead
					if minMana == 0 || cur.mana_spent < minMana {
						minMana = cur.mana_spent
					}
				}

				cur.boss_attack()
				if cur.player_hp > 0 {
					// player not dead
					if cur.mana_spent < minMana || minMana == 0 {
						newGames = append(newGames, cur)
					}
				}
			}
		}
		games = newGames
	}
	return minMana
}

func main() {
	g := Game{50, 500, 55, 8, 0, 0, 0, 0}
	fmt.Println("Part 1 =", gameSearch([]Game{g}, false))
	fmt.Println("Part 2 =", gameSearch([]Game{g}, true))
}
