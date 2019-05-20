package main_test

import (
  // . "github.com/gavmor/canon"?
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

type Player struct {
  Score int
  Strategy Strategy
}

type Hawk struct {
}

func (Hawk) Cooperates() bool {
  return false
}

type Dove struct {
}

func (Dove) Cooperates() bool {
  return true
}

type Strategy interface {
  Cooperates() bool
}

func Play(a Player, b Player) (Player, Player) {
  if a.Strategy.Cooperates() {
    b.Score += 2
  }
  if b.Strategy.Cooperates() {
    a.Score += 2
  }
  return a, b
}

var _ = Describe("Play", func() {
  It("lets the wookiee win", func() {
    hawk := Player{
      Score: 0,
      Strategy: Hawk{},
    }
    dove := Player{
      Score: 0,
      Strategy: Dove{},
    }

    hawk, dove = Play(hawk, dove)

    Expect(hawk.Score).To(Equal(2))
    Expect(dove.Score).To(Equal(0))
  })

  It("ignores argument order", func () {
    hawk := Player{
      Score: 0,
      Strategy: Hawk{},
    }
    dove := Player{
      Score: 0,
      Strategy: Dove{},
    }

    dove, hawk = Play(dove, hawk)

    Expect(hawk.Score).To(Equal(2))
    Expect(dove.Score).To(Equal(0))
  })
})

type World struct {
  Players []Player
}

func NewWorld() World {
  players := []Player{}
  for i := 0; i < 50; i++ {
    players = append(players, Player{Strategy: Hawk{}})
  }
  for i := 0; i < 50; i++ {
    players = append(players, Player{Strategy: Dove{}})
  }

  return World{Players: players}
}

var _ = Describe("World", func() {
  It("Has a diverse set of strategies", func() {
    Expect(NewWorld().Players).To(ContainElement(Player{Score: 0, Strategy: Hawk{}}))
    Expect(NewWorld().Players).To(ContainElement(Player{Score: 0, Strategy: Dove{}}))
  })
})

var _ = Describe("Scoreboard", func() {
  It("counts players by strategy", func() {
    Expect(NewWorld().Players).To(Equal([]float64{50, 50}))
    Expect(NewWorld().Players).To(ContainElement(Player{Score: 0, Strategy: Dove{}}))
  })
})
