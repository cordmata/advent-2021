package main

import (
	"log"
	"strings"

	"github.com/cordmata/advent-2021/utils"
)

func main() {
	log.Println("Part1 example:", part1(exampleInput), "== 4512")
	log.Println("Part1 winner:", part1(actualInput), "== 58412")
	log.Println("Part2 example:", part2(exampleInput), "== 1924")
	log.Println("Part2 winner:", part2(actualInput), "== 10030")
}

func part1(in string) int {
	hopper, boards := parseInput(in)
	for _, h := range hopper {
		for _, b := range boards {
			b.playNumber(h)
			if b.isWinner() {
				return b.score() * h
			}
		}
	}
	panic("no winner")
}

func part2(in string) int {
	hopper, boards := parseInput(in)
	boardsWon := make(map[*board]bool)
	for _, h := range hopper {
		for _, b := range boards {
			b.playNumber(h)
			if _, ok := boardsWon[b]; !ok && b.isWinner() {
				boardsWon[b] = true
				if len(boardsWon) == len(boards) {
					return b.score() * h
				}
			}
		}
	}
	panic("no loser")
}

type cell struct {
	val    int
	marked bool
}

const boardSize = 5

type board [boardSize][boardSize]cell

func (b *board) playNumber(n int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cell := &b[i][j]
			if cell.val == n {
				cell.marked = true
			}
		}
	}
}

func (b *board) isWinner() bool {
	var colMatches [boardSize]int
	for i := 0; i < boardSize; i++ {
		var rowMatches int
		for j := 0; j < boardSize; j++ {
			cell := &b[i][j]
			if cell.marked {
				rowMatches++
				colMatches[j]++
			}
		}
		if rowMatches == boardSize {
			return true
		}
	}
	for i := 0; i < boardSize; i++ {
		if colMatches[i] == boardSize {
			return true
		}
	}
	return false
}

func (b *board) score() int {
	var score int
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cell := &b[i][j]
			if !cell.marked {
				score += cell.val
			}
		}
	}
	return score
}

func parseInput(i string) ([]int, []*board) {
	parts := strings.Split(i, "\n\n")
	hopper := utils.StringSliceToIntSlice(strings.Split(parts[0], ","))
	var boards []*board
	for _, b := range parts[1:] {
		boards = append(boards, parseBoard(b))
	}
	return hopper, boards
}

func parseBoard(i string) *board {
	board := &board{}
	for i, row := range strings.Split(i, "\n") {
		for j, c := range utils.StringSliceToIntSlice(strings.Fields(row)) {
			board[i][j] = cell{val: c}
		}
	}
	return board
}

var exampleInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

3 15  0  2 22
9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`

var actualInput = `15,61,32,33,87,17,56,73,27,83,0,18,43,8,86,37,40,6,93,25,14,68,64,57,39,46,55,13,21,72,51,81,78,79,52,65,36,92,97,28,9,24,22,69,70,42,3,4,63,50,91,16,41,94,77,85,49,12,76,67,11,62,99,54,95,1,74,34,88,89,82,48,84,98,58,44,5,53,7,19,29,30,35,26,31,10,60,59,80,71,45,38,20,66,47,2,23,96,90,75

26 68  3 95 59
40 88 50 22 48
75 67  8 64  6
29  2 73 78  5
49 25 80 89 96

57 26 21 56 70
38 48 78 40 54
82 71 22 24  4
16  9 65 42 79
43 94 39 12 67

70 97 26 39 22
65 11 69  6 93
71 74 72 57 59
61 21 89 86 17
66 15 94 79 85

46  4 55  2 92
 7  8 53 65 42
49 35 99 77  0
82 28 25 43 33
79 12 58 81 71

82  5 63 98 48
78 61 29 68  9
19  6 69 73 89
20 81  8 17 11
24 90 59 21 91

76 87 99 12 78
14 97 19 70  9
66 44 88 30 63
58 85 55 24 36
49 10  6  2 92

33 63 13 28 70
73 67 69 27 91
85 71 98 37 36
46 80 21 97 75
54 92 44 11  2

88  4 55 76 75
29 79 30 26 18
45 31 72 77 86
 7 14 74 94 98
93 78  8 66 16

75 34 74 94 59
31 41 11 85 68
12 24 19 80 29
 7 97 77 73 14
56 27 26 72 35

15 35 65 77 76
31  9 89 73 63
54 18  4 39 12
22 75 30 67 33
79 84  6 64 92

75 58 28 83 37
81 82 62  7 25
31 55 51 42 41
53 38 85 69 63
 0 89 20 96 79

43 29 62 59 27
 6  1 66 77 32
70 37 57 82 38
22 60 19 68 88
99 93 40 28 47

76 47 68 60 28
57 34 35 83 11
 2 51  0 49  9
31 65 14 90 46
96 81 32 63 24

17 38 59 27 54
26 78 91 40 80
74 98 77 14 53
13 49  3 63 46
45 34 64 66 37

 9 62 22 36 32
40 56  2 98 74
71 45 58 31 42
14  1 51 39 66
86 18 90 88 20

98 61  0 78 22
 6 42 83 74 30
17 39 45 52 75
49 18  4 67 91
12 84 89 10 87

12 40 69 11 87
70 73 14 92 78
99 67 34 50 94
43 49 82 68  1
86 20 16 18 63

26 13 68 33 88
41 73 17 78 12
10 82 62 19  8
58  3 84  5 70
97 91  2 46 15

48 61 90 54 73
46 77 47 37 49
41 81 79 88  0
23 86 84 65 60
87 91 34 78 38

 4  9 58 62 99
35 11  6 13 25
68 67 78 28 19
95 44 63 92 79
20 61 54 88 81

78 75 59 79 44
 6  3  7 33 10
41  2 90  9 81
20  4 55 54 27
28 51 64 95 21

28 55 99 76 58
61 18 71 67 65
80 84  6 63 46
72 81 29 35 10
64 94 15 26 79

57 18 32 13 70
81  3 84 91 75
10 15 63 16 87
97 86 95 58 50
83  2  0 12 89

48 37  6 61 43
49 89 80 93 50
64 12 23 82 99
38 36 69 67 63
25 66 72 92 95

 4 20 11 85 27
88 29 59 79 48
16  2 64 55 90
73 82  1 56 61
34  8 23 32  6

24 12 80 18 14
54 86 89 66 38
 0 62 46 72 17
77 37 74 79 85
 1 29 87 42 15

54 88 75 16 33
76  9 19 69  2
14 98 46 87 67
32  6 62 45 25
30 38 52  0 51

15 32 85 22 81
49 59 72 74 23
54 33 34 84 89
37 12 55 83  2
11 25 51 24 38

72 20 87 38 24
74 28 76 42 94
23 50 75 80 18
15 53 85 95 68
88  0 34 59 27

15 69  3 45 94
 8 90 27 95 48
37  9  2 74 59
43 46 25 98 83
68 73 24 31 72

63 65 61  7 57
95 51 66 64 17
94 93 44 58 59
30 90 97 23 10
47 92 56 20 70

26 30 45 99 31
81 41  3 69 83
44 89 74 72 65
86 20 29 67  4
64 84  9 33 18

49  7 38 88 94
98 54 36 89 23
34 51 70 79 61
19 15  4 27 73
12 13 35 44 95

36 49 60 27 53
44 81 23 50 73
12 31 48 63 61
91 32 47  1 87
88 96 70 90 10

80 13 46 51 61
20 12 71 73 54
27 86 21 62 37
30 25 67 14  6
55 78 88 32 57

99 71 10 58 52
26  5 36 44 86
56 46 16 90  3
94 89 53 23 31
77 14 34 32 54

59 60 15  0 45
55 78 48 14 33
37 32 20 68 19
 8 24 26 52 23
82 70 40 21  7

61 70 95 96 18
26 89 43 33 92
50 88 64 54 76
78 19 56  7 68
29 59 77 71 63

48 28 79 68 64
61 38 52 78 59
31 41 10 39 77
34 15 43 98 73
17 54 55 75 27

55 49  3  5 71
85 48 75 95 45
88  2 56 29 31
83 70 53 52 66
74 27 89 50 91

32 61 13  1 37
 9 92 35 26 39
98 25  7  5 12
 0 99 17 82 27
19 42  8 21 30

33 77 59 37 42
80 64 61 73 79
15 41 58 45 93
86  2 92 57 83
20 24 17 13 23

98 57 16 14 99
70 43 97 94 52
19 89 88 54 17
44 41 15 60  4
31 71 33 96 68

 9 69 10 79 43
84 19  5 48 71
40 22 66 89 82
36 62  6 76 81
51 18 30 93 75

96 42 32 35 36
28 81 53 87 92
51 34 91 80 15
23 29 62 98  4
25  2 83 41 46

55 79 98 90 27
33 75 48 38 39
21 65 52 63  2
85 53  5 88 15
50 96 70  3  8

22 45 77 61 69
15 81 71 59 26
74 12 30  2 27
25 78 24 70 65
99 66 35 16 57

44 85 88 14 64
67 55 47 98 99
87 57 10 84 27
42 28 39 81 56
46 76  8 75 95

96  4 48 28 81
67 29  6 30  8
20 24 64  7 12
16 71 59 19 99
82  9  0 62 87

70 98 10 97 92
 7 94 67 20 26
77 13 69 61 51
 4 71  3 28 91
85 11 27 56 54

 3 98 31 47 75
34 95  8 27 42
74 49 80 79 11
15 17 89 85 33
55 52 32 36 45

 4 78 93 29 63
13 24 40 17 75
12 92 48 82 60
26 54  8 47 37
41 57 36 32 99

 1 27  8 47 41
86 20 53 61 96
45 21 80 58 64
 5 95 23 13 10
81 87 49 24 50

27 78 50 18 43
23 75 77 38 29
71 93 64  5 56
34 84 67 52 79
90 95 19 46 88

83 41 79 67 69
60 47  2 43 85
12 17 28 89 81
16 18 98 35 62
 7 45 25 40 58

79 21  7 85 76
55  8 14  3 72
25 30 62  6 82
38 16 32 95 59
27 99 33 75 98

65 76 69 98 78
94 55 31 73 77
10 14 79 58 22
26 34 16 87 29
 2 24 30 27 91

31 42 38 93 88
36 68  7 66 59
23 71 45 72 94
52 81 84 27 41
18 49 76 82 70

50 58 95 52 35
21 46 68 71 59
34 84 76 62 57
94 41 99 77 55
69 48 97 78 73

53 86 48  8 26
 3 72 57 27 23
99  4 71 21 50
39 18 54 41 82
 7 46  1 65 96

63 98 33 80 56
89 75 15 22 59
69 36  0 86 12
21 41 55 49 74
 7 90 76  5 44

78 40  2 61 76
25  5 42 88 35
 1 41 28 71 85
 3 34 22 72 23
15 56 67 38 68

11 89 28 48 87
57 80  1 42 33
59 18  7 24 65
30 79 12 68 83
44 82  2 53 58

 7 14 22 23 29
53 37 48 86  3
56 25 54 82 43
 0 91  6 17 49
33 95 63 94 12

86 62  0 47 69
80 91 37 15 46
50 28 75 83 31
65  5 39 23 55
88 84 72 70 74

36 31 82 32 78
30 18 11 29 38
55 84  9 33 57
16 51 48 77 58
73 22 79 85 54

70 66 89 40 55
75 17 36 88 28
22 97 92 43 72
25 27  3 18 45
13 14 54 12 74

11 43 96 92 51
 7 59  2 32 69
79  0 46 68 80
18 95 88 39 60
84 14 58 36 22

82 88 64 85 51
45 95 50 27 99
15 13 21 69  9
53 36 79 22 68
83  8 92 65 32

31 37 89 28 26
17 25 99 20  5
97 49 21 60 83
55 57 16 40  6
45 39 33  0 65

86 51  0 31  7
95 33 85 87 14
32 48 91 46 36
60 90 88 38 15
52 75 40 23 84

25 35 57  8 86
 3 59 46 96 13
 0 41 45 76 79
97 36 60 26 53
33 74 64 66 93

59  3 96 84 71
39 90 61 77 19
92 38  6 32 54
12  5 62 86 75
43 98 23 82 33

 1 67 51  6 94
57 44 53 90  2
19 89 80 30 45
42 88 62 98 33
20 63 78 56 83

47 21 70 31 75
19 38 91 85 73
22 27 54 86 13
 8 49  7 89 37
32 25 17 16  0

64 77  7 23 83
56  2 17 65 60
43 98 68 67 18
22 96 72 69 86
20 26  4 84 16

30 31 95 98 48
36 11 92 60  5
 0 76 73 27 14
50 46 38 53 33
12 97 59 61 51

93 45 66 91 63
80 75 52 55  1
31 68 76 24 79
15  2 42 70 20
89 90 21 25 48

36 99 49 83 57
24 79 89 91 63
58 47 27 74 38
90 54 39 40 98
 7  2 77 14 86

44 33 12 86  8
65 92 74 52 55
20  3 78 28 47
80 17 11 41 29
62 18 39 48  7

63 52 87 81 14
91 56  4 84 27
 9 24 68 18 47
57 44 26  0 37
40 75 11 88 20

61  4 91 31 79
67  7  6 34 95
19 23 62 99 50
43  1 37 16 74
38 94 47 10 25

90 51 37  7 16
68 61 28 65  1
58 80 49 11 23
24  8 12  6  4
30 75 19 63 53

70 90 34  4 97
73 26 87 61 88
38  2  0 71 28
57 69 18 15 60
80 39 78 33 36

44 10  3 46 31
43 57 12 29 92
 0 61 54 23 52
 5 55 27 93 11
24 14 30 87 99

 5 20 33 37 97
78 83 50 93 65
30 59 74 68 27
 4 32 90 16 79
52 22 76 45 41

29 53  9 20 15
17 61 94 52 83
43 82 97 14 57
18  2 16 95 72
30 39 79 65 25

22  5 42 15 73
32 16 29 36 77
 9 53 98 69 18
97 56 79 66 88
90 99  3 10 84

71 92 19  1 80
17 21  4 54 61
27 66 20 63 49
18 74 11 70 39
97 98 64 34 10

66 83 73 54 57
68 10  8 17 22
53 87 71 18 40
43  4 65 89 59
27 35 47 15 46

87 76 88 54  2
42 68 47 44 17
16 70 10 53 43
 7 78 12 39 83
15 65 96 85 24

73 28 78  4 98
97 56 16 69  6
46 90 18 63 81
26 95 19 30 31
59 32 49 21 13

68 48  7 85 12
58 95 41 59 78
 1 28 53 51  9
10 93 97 91 65
61 75 63 23 57

26 82 40 53 11
 0 22 68 99 96
64 45 74  5 92
84 33 13 34 73
47 54 81 77 46

12 83 25 82 72
 8  0 95  6 40
17 64 27 23 91
14 73 70 55 44
69 76 92 78 56

20 15 45 44 52
94 26 61 38 64
84 67 16 23 21
73 71  5 10 36
62 65  9 24 58

59 75 60  0 97
41 94 73 86 51
 8 89 22 45 18
 3 63 85 57 16
42 44 10 23 93

 3 68 80 19 59
41  6 92 58 28
94 57 81  5 71
90 54  9  8 14
32 96 30 37 10

15 16 14 10 52
51 26 54 24 84
45 90 28 36 96
56 70 86 94 32
67 81 13 29 27

67 30 89 43  3
86 10 38 90 77
70 78 97 94 33
29  8 85 69 56
40 80 47 12 17`
