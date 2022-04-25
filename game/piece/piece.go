package piece

type Piece struct {
	Name            string
	CurrentPosition string
	CurrentY        int
	CurrentX        int
	HasMoved        bool
	PossibleMoves   []string
}

func (p Piece) GetName() string {
	if p.Name == "" {
		return " "
	} else {
		return p.Name
	}
}

//func Build(name string, currentX int, currentY int) Piece{
//
//}