package domain



type GameState struct {
    GameID string
    Pot int
    Bets map[string]int
    SmallBlind int
    BigBlind int
    middle []Card
    folded map[string]string
    decision map[string]string
    hand map[string][]Card
    isRuning bool
    turn int
    round int
    Deck []Card
}


