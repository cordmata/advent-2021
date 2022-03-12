package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cordmata/advent-2021/utils"
)

type heightmap [][]int

func (m heightmap) findNeighbors(x int, y int) []int {
	var neighbors []int
	//left
	if x > 0 {
		neighbors = append(neighbors, m[y][x-1])
	}
	//right
	if x < m.width()-1 {
		neighbors = append(neighbors, m[y][x+1])
	}
	//up
	if y > 0 {
		neighbors = append(neighbors, m[y-1][x])
	}
	//down
	if y < m.height()-1 {
		neighbors = append(neighbors, m[y+1][x])
	}
	return neighbors
}

func (m heightmap) width() int {
	return len(m[0])
}

func (m heightmap) height() int {
	return len(m)
}

func part1(input heightmap) int {
	var totalRiskLevel int
	for y := 0; y < input.height(); y++ {
		for x := 0; x < input.width(); x++ {
			val := input[y][x]
			neighbors := input.findNeighbors(x, y)
			minNeighbor := neighbors[0]
			for _, n := range neighbors[1:] {
				if n < minNeighbor {
					minNeighbor = n
				}
			}
			if val < minNeighbor {
				totalRiskLevel += val + 1
			}
		}
	}
	return totalRiskLevel
}

func part2(input heightmap) int {
	return -1
}

func processInput(s string) heightmap {
	var out heightmap
	for _, line := range strings.Split(s, "\n") {
		out = append(out, utils.StringSliceToIntSlice(strings.Split(line, "")))
	}
	return out
}

func main() {
	exampleInput := processInput(exampleInput)
	actualInput := processInput(actualInput)

	p1Example := 15
	if r := part1(exampleInput); r != p1Example {
		log.Fatalf("[PART 1]: expected %d, got %d", p1Example, r)
	}
	fmt.Println("The answer to part 1 is:", part1(actualInput))

	p2Example := 0
	if r := part2(exampleInput); r != p2Example {
		log.Fatalf("[PART 2]: expected %d, got %d", p2Example, r)
	}
	fmt.Println("The answer to part 2 is:", part2(actualInput))
}

var exampleInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

var actualInput = `9234598321279999876543212397634598789843210123456789212999878987556456999878965432459101987654567899
8965987432367898989864301986545789698765421939697894349876967896432349899769896543598929898763478998
7896796543458997698765212987676994569865439898989965498765456987321099789858789654987898789854567987
6789899654569987549876943498787893498979698787678976789886349876532988656645699769876745699965679896
5678998965878976534989894999898932987898987656567898999943234987849876543234678998765434788978798765
4556987896989765423498789899989899876567986745456789898932145798769987653123789987654323487899899954
3249976987899654312987677789976798765456985432345689767991015679878996544012999899873212376789939853
4398765898998995499876585678965459854369876541237895456789324589989988432167898798765901234998749754
5469876789987989989965434567894398765498765430125789345695435678999876543456789679989892945897659886
6798987899876875679954321236789199876599876321234599216789576789656987664667894589998789896789769987
7987698998765654568899410145699989987987976532345678924597689996545698875798923678997656789899898798
9878549976543123456798921234987978999876986543467899435789799985324569989999434569976545689999997659
9963234898794012767987892349876867998765499876578929876789898765435678990196545698987438795498986543
9852123789989123459876789498765657898654345987689210989892999876549889991987656997696545894387897651
8743245678978954598975345998654545789654234599789329999901997987669998989998779886598767895256789872
7654346789567895987654249889843437678965123489895498999899886598798997678999898775459898952145699993
8765657893458976798765398776762123489987434567976987898789765439987876587999987654345999943234988989
9878767892599987899876987654321012591296545779989986789678954329876545456789876565234689895949877565
6989878901989898965987898765434323690987856889798765454599769212985434345699965432123456789899765454
5496989329878799654398969876567434589998967897659854343489898909874323234579876551012348997798974343
1345695498767678967989656987698997678969878998798765212378987898765210123567965432123459976656793212
0123987987654569979876545698789989789654989999897654324567896569984321234878998643234567895436789393
9294599876543456895987656789899879998793297899989775435978935498765432345679879784345678976545899989
8989678998765567984398768891998767899989356998978976567899321239886545456789769895456789987656789878
7879989889878678965679879932998658999979459897869987678989990347998676768896653939869893499767898769
6567897679989789976789989899876549679868998786458898789879889456789789879965432123978932349878987656
3458987567899897899895496798765434569659987655377789898767678968999997989987543234989321298989996543
4567893468998956568901345987654325698789876544265678987956567899019876596898655455696539987899987632
5678912345987545456912346998765412389896987432103459876545458942198765445798796578789998976989998321
6789101256996433347893459899898701278965496545212398765432369893239893234789987678999877965578999210
9893212369874321236789569789987632369896397876343459854321236789945984345696598789398765624466799321
9954323498765438345699998679876543456789298997456769968534345699896976789789459899299544312345678943
8765434789876567456789877545987854667892129976567878987678456789789899899892398978987632101234589654
9876545678987876569899965432398965878921012987678989898989567898679789921999987667898543612345678969
0988656789398987878998998321239878989432123698889998769499878977545678930198765456987659543456789878
1998767895459298989787987532456989996569234569999989654323989765434567921239874366798768767568999989
9879878987679109899656796543456799987898945678989978963212399874325656893498765234769979878679589995
8765989898998998768943987654697989998987896989878867942103456965412346789999979123458989988789678954
9654396789987987656792198967989978999896789898768756893214579876525457899889898938767996899898789543
8964245699876796545789999879878767998765458799657646789925689987436567988776767899878924456989899952
7893126988765987636799899998767656987654323678943434567899798998547679876565457789989313345678999891
6789019876854398747898789987654345698766434567932123456798987898798798765432345699993201234789998789
5698929865464239658975678998743234769876556778943254567987786799899899895601246899874312345899987678
4567897654321099899464568987645123456998667889765345678996695689932999954312356789965423656899876567
3456798976542987954353459876532012368909788998765456989985434599321098765424456899976534569998765456
2169989897953496543242365987432143499219899899878567899876745678934989989535767899876545698999764345
3298676789894569732101234599843254789398965674989698943987658789549876596546878999987676987899853234
4987545698789698654232345798754365678987894353499999432398767897698785459657989789998799876598767346
5988436789698789765343456799765496889876789212989894341239878998987654398768997696999899987439878457
9876521234589899876457667987989987898765678909978799210949999459896543239878978545789998996521989568
9983210123679956998568779995496798969876789998765678929898989598765432123989865734678987889432987689
9876421234598949987679899874345679654989892987654599898767678999989821012398754324569876767943698797
5987562455987898998789953965234989793498901296543456797656546789998733123579863213498765656794569896
4398778569876567899899899854345899989567892395432345987545435679987654534589965624569654345989678989
3219887698765475899999798765456789878978993987541256796432323778998768675799876787679643234578999678
7323998987654234789987659896567896767899989995432348989321012567899879786795987898798932123456789567
6567919999843044578996543919878975456999879876744559875458723478956989899894398999987321034569893478
7678909898762123459999432101989765349998767987865667996567654589547996987943239789996432123688954567
8789398789873434967898944912399854298767656799976878989998765695439875456894345678987853434567897678
9893299632986547898997899894498743197654745679987889877899877899523989567965456789998754545688999789
9999986521297656789986798789999654986543234568998998766899988978912399699979569899879867659789987993
8998765434398767898765987677898965697654345789999989454598999569923498789998678999769878767892396432
7899896565499878919874396565987896899865466789899876323457892398896569997999799987643989898901987541
6789987776987989434965989454576789919876778898798765212348901986789789346799989899952197999329987632
5899998889876799949879876323345892101987989987659854105459912965678991299899878788891016799498998543
6999879998785679899989965410236943232398999998639873212567899896789210989998767677789125678987899654
9899867987654563778999874322345894345989878996521964323456976789894329878987656565678934589776789765
8785659876543212567898765434456789657976569985439876434769765678965498769876543454567895997645678976
7654545987654103458929876546567898779865478976556989545678954567896987659876532143458797896534569987
6543434598763216567910997657698999989976678987767898767789323456789498546998321012345689994323459998
7652123569854347678921598798789998999987899499898929978894312348992349656989652123478998989012368999
8761012478965456989432459899899877889998989325979213989953201237893498979876543456567897879923479989
9872123569876578996563769953998765678999878934569101296543213456789987895987654567878996567895678978
9983237678989989987875878992199954345698967996678932987854324567899876843498765878989985459998789569
9994345789896593598989989989989893296987656789789949899765435878998765672349876989799876347899893458
8987656998795432699399899979878789989876545678999898789976546789329984421234987897653975456921932567
7699867897654321989498787656965698878965434567898785678987987899909793210123498998792986567890953678
6549878998975439878987655649879987667896545878987674569898998999897654321266569329989597678979894589
5432989899876598769876544434998865456789676789996543456789999898789987432345689499965498799567779694
4321399789987789653989432123987654325698989899985432359899988789678976543476789987857349895435567943
5930298698999899542398949239876543214567899969876545468999878656567897854589899976744234989323457892
9891976567899987653987898949989876523678998654998656567898764545456799765678999865433145678939598921
9769895498979898769856767898999965434789998743479787978999653232345678976789598764321034989998999210
7658789348968769898745654567899876545899989432369898989987632101457899987895439989452123899987895439
4545678957954356987632343456899987856789876543456999699876543234578995498999621296573234789876789598
3234569767893239898521012347898998768999987656767896545987654567899789239998990987684345679765678987
2123678979964198765432129469957989879459998797898987634598765678956689123987989998765456789654569876
1038789989979099898547298989239876989968929989989654323989876789244579019876567999876567894323498765
2149898798998987997656987892198765699899439878978993219876989892133459199965456789998778943212987854
3234989656487896598969876989019854569789598765767889109954599943012998989876345678919889654309876543
4545679545346789459896965678929543298678987654656778998765678932199897678987558789324998785412987642
5656795431257899398795454589998754596567896543244567899877899643989676567987678999939019876543965431
8767896910129998987684343456899765985456897654123456789989987659876543456798999998898923989659878932
9898959891235987897543212369989899874346798765034567896593498767987852349899212987677894598767989543
9949346789349876789654301298878987653245689872123458989432379979876543498954329876566789679878998654
8431245678998765678963212987767898732176789943234589878943467989987694987895698765445678989989998765
7510356799239854345894349896556789621018997654345678967894568997698989876799987654334569496799879876
6421237890198765676789498765445698542129998765456789656789879876569878985678998643210579345997764988
6532345893239976987896569874325987653457899877667897645678989985498765434567987654323478959876543499
7643456789345987898987689985434598764567899988798998798799999876349876546778998765434567899987201234`
