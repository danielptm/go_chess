package piece

type Piece struct {
	Name            string
	CurrentPosition string
	CurrentY        int
	CurrentX        int
	CossibleMoves   []string
}

func (p Piece) GetName() string {
	if p.Name == "" {
		return " "
	} else {
		return p.Name
	}
}

//TODO: Add a unique and possible move to the possibleMoves array
func (p *Piece) AddMove() {

}

//TODO: Generate a theoretically possible move. The board will determine if it is ok. Make this an interface method that all pieces can implement.
func (p Piece) GenerateMove() {

}
