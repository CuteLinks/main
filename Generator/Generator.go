package generator

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz_0123456789"

type Generator struct {
	randomizer randomizer
}

func New (randomizer randomizer) Generator {
	return Generator{randomizer: randomizer}
}

//var seededRand *rand.Rand = rand.New(
//	rand.NewSource(time.Now().UnixNano()))

func (g *Generator) StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[g.randomizer.GetRandomInt(len(charset))]
	}
	return string(b)
}

func (g *Generator) String(length int) string {
	return g.StringWithCharset(length, charset)
}

