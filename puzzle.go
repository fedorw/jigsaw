package main
import "fmt"


type Board struct {
	pieces [16]*Piece
	available [16]*Piece
	last int 
}

func (b *Board) clone() *Board {
	bc:=new(Board)
	for i:=0;i<16;i++ {
		bc.pieces[i]=b.pieces[i]
		bc.available[i]=b.available[i]
	}
	bc.last=b.last
	return bc
}
func (b *Board) clean() {
	for n:=0;n<16;n++ {
		b.pieces[n]=nil
	}
}
func (board *Board) Draw()  {
	fmt.Println("-----------------------------------------------------------------------")
	for i:=0;i<16;i+=4 {
		a1,a2,a3,a4:=
			board.available[i+0],board.available[i+1],
			board.available[i+2],board.available[i+3];
		p1,p2,p3,p4:=
			board.pieces[i+0],board.pieces[i+1],
			board.pieces[i+2],board.pieces[i+3];
		fmt.Print("  " + GetSymbol(0, p1) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(0, p2) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(0, p3) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(0, p4) +" ")
		fmt.Print("   |  " )
		fmt.Print("  " + GetSymbol(0, a1) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(0, a2) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(0, a3) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(0, a4) +" ")

		fmt.Print("\n")
		center1,center2,center3,center4:=" "," "," "," "
		if p1!=nil&&p1.flipped {
			center1= "."
		}
		if p2!=nil&&p2.flipped {
			center2= "."
		}
		if p3!=nil&&p3.flipped {
			center3= "."
		}
		if p4!=nil&&p4.flipped {
			center4= "."
		}
		fmt.Print(GetSymbol(3, p1) + center1+ GetSymbol(1, p1));
		fmt.Print(GetSymbol(3, p2) + center2+ GetSymbol(1, p2));
		fmt.Print(GetSymbol(3, p3) + center3+ GetSymbol(1, p3));
		fmt.Print(GetSymbol(3, p4) + center4+ GetSymbol(1, p4));

		fmt.Print("  |  " )
		fmt.Print(GetSymbol(3, a1) + " "+ GetSymbol(1, a1));
		fmt.Print(GetSymbol(3, a2) + " "+ GetSymbol(1, a2));
		fmt.Print(GetSymbol(3, a3) + " "+ GetSymbol(1, a3));
		fmt.Print(GetSymbol(3, a4) + " "+ GetSymbol(1, a4));
		fmt.Print("\n")

		fmt.Print("  " + GetSymbol(2, p1) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(2, p2) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(2, p3) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(2, p4) +" ")

		fmt.Print("   |  " )
		fmt.Print("  " + GetSymbol(2, a1) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(2, a2) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(2, a3) +" ")
		fmt.Print(" ")
		fmt.Print("  " + GetSymbol(2, a4) +" ")


		fmt.Print("\n\n")

	}
}
func (b *Board) getNeighbours(pos int) []*Piece {
	n:=make([]*Piece, 4)
	if pos<=3 {
		n[0]=nil
	} else {
		n[0]=b.pieces[pos-4]
	}
	if pos==3||pos==7||pos==11||pos==15 {
		n[1]=nil
	} else {
		n[1]=b.pieces[pos+1]
	}
	if pos>=12 && pos <=15 {
		n[2]=nil
	} else{
		n[2]=b.pieces[pos+4]
	}
	if pos==0||pos==4||pos==8||pos==12 {
		n[3]=nil
	} else {
		n[3]=b.pieces[pos-1]
	}
	return n
}

type Piece struct {
	rot int
	con [4]int
	flipped bool
}

func (p *Piece) getConn(side int) int {
	return p.con[side]
}

func (p *Piece) flip() (bool) {
	var j=p.con[1]
	p.con[1]=p.con[3]
	p.con[3]=j
	p.flipped=true
	return true
}
func (p *Piece) rotate() (bool) {
	var j=p.con[3]
	p.con[3]=p.con[2]
	p.con[2]=p.con[1]
	p.con[1]=p.con[0]
	p.con[0]=j
	p.rot++
	if (p.rot==4) {
		p.rot=0
		return false
	}
	return true 
}
func NewPiece(c1, c2, c3, c4 int) *Piece {
	p:=new(Piece)
	p.rot=0
	p.flipped=false
	p.con=[4]int{c1,c2,c3,c4}
	return p;
}
func GetSymbol(pos int, p *Piece) string {
	if p==nil {
		return "   "
	}
	sym:=p.con[pos] 
	inside:=sym<0
	if (inside) { sym=-sym }
	str:=""

	A,B:=1,4
	if (inside) {
		B,A=1,4
	}
	switch sym {
		case 2: str="o" 
		case 3: str="+"
	case A:
		switch pos {
			case 0: str="v"
			case 1: str="<"
			case 2: str="^"
			case 3: str=">"
		}
	case B:
		switch pos {
			case 0: str="^"
			case 1: str=">"
			case 2: str="v"
			case 3: str="<"
		}

	}
	if (inside) {
		return "("+str+")"
	}
	return " "+str+" "
}

func setup() []*Piece {
	pieces:=[]*Piece{
		NewPiece(1,2,-4,-1),NewPiece(1,1,-3,-4),
		NewPiece(1,2,-3,-2),NewPiece(2,4,-4,-2),

		NewPiece(4,1,-2,-4),NewPiece(4,4,-1,-2),
		NewPiece(3,2,-2,-3),NewPiece(3,1,-1,-4),

		NewPiece(4,3,-2,-4),NewPiece(3,2,-2,-1),
		NewPiece(1,1,-3,-2),NewPiece(3,2,-4,-4),

		NewPiece(1,2,-1,-3),NewPiece(2,2,-4,-2),
		NewPiece(2,2,-4,-2),NewPiece(2,4,-3,-2),
	}
	return pieces
}
func (_b *Board) place(p *Piece) *Board {
	if p== nil {
		return nil
	}

	b:=_b.clone()
	pos:=b.last+1
	if b.pieces[0]==nil {
		b.pieces[0]=p
		b.last=0
		for j:=0;j<16;j++ {
			if p==b.available[j] {
				b.available[j]=nil
			}
		}
		return b
	}

	neighbours:=b.getNeighbours(pos)

	if neighbours[0]!=nil && neighbours[0].getConn(2) + p.getConn(0) !=0 {
		return nil
	}
	if neighbours[1]!=nil  && neighbours[1].getConn(3) + p.getConn(1) !=0 {
		return nil
	}
	if neighbours[2]!=nil && neighbours[2].getConn(0) + p.getConn(2) !=0 {
		return nil
	}
	if neighbours[3]!=nil && neighbours[3].getConn(1) + p.getConn(3) !=0 {
		return nil
	}
	b.pieces[pos]=p
	b.last=pos
	for j:=0;j<16;j++ {
		if p==b.available[j] {
			b.available[j]=nil
		}
	}
	return b
}
func solve(board *Board, flip bool) {
	if board.last==15 { // done
		board.Draw()
		return
	}
	loop:for i:=0;i<16;i++ {
		if board.available[i]==nil { 
			continue loop
		}

		res:=board.place(board.available[i])
		if res!=nil { solve(res,flip) }

		for r:=0;r<3;r++ {
			board.available[i].rotate()
			res=board.place(board.available[i])
			if res!=nil { solve(res,flip) }
		}

		if (!flip) {
			continue
		}

		board.available[i].flip()

		res=board.place(board.available[i])
		if res!=nil { solve(res,flip) }

		for r:=0;r<3;r++ {
			board.available[i].rotate()
			res=board.place(board.available[i])
			if res!=nil { solve(res,flip) }
		}

	}
}

func main() {
	pieces:=setup()
	board:=new(Board)
	board.last=-1
	for j:=0;j<16;j++ {
		board.available[j]=pieces[j]
	}

	tmp:=new(Board)
	for j:=0;j<16;j++ {
		tmp.pieces[j]=pieces[j]
	}
	tmp.Draw()

	solve(board, false /*flip*/)

}


/*

1 arrow in      -1  arrow out
2 ball          -2  ball
3 plus          -3  plus
4 arrow out     -4  arrow in

1,2,-4,-1
1,1,-3,-4
1,2,-3,-2
2,4,-4,-2

4,1,-2,-4
4,4,-1,-2
3,2,-2,-3
3,1,-1,-4

4,3,-2,-4
3,2,-2,-1
1,1,-3,-2
3,2,-4,-4

1,2,-1,-3
2,2,-4,-2
2,2,-4,-2
2,4,-3,-2


*/
