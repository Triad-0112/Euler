package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard = map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	thechosen, rankexists := cb[rank]
	if !rankexists {
		return 0
	}
	count := 0
	for _, onBoard := range thechosen {
		if onBoard {
			count += 1
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}

	Count := 0
	for _, rank := range cb {
		if rank[file-1] {
			Count += 1
		}
	}
	return Count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	squarescount := 0
	for _, rank := range cb {
		for range rank {
			squarescount += 1
		}
	}
	return squarescount
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	Occupieds := 0
	for _, rank := range cb {
		for _, OnBoard := range rank {
			if OnBoard {
				Occupieds += 1
			}
		}
	}
	return Occupieds
}
