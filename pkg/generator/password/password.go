package password

import (
	"errors"
	"math/big"
	"math/rand/v2"
	"strings"

	crRand "crypto/rand"
)

// PasswordType represents the type of password to generate
type PasswordType string

var wordList = []string{
	"Sordid", "Optional", "Faraway", "Computer", "Programming", "Excellent", "Fantastic", "Incredible",
	"Amazing", "Beautiful", "Wonderful", "Ordinary", "Common", "Rare", "Unusual", "Strange", "Silly",
	"Serious", "Happy", "Sad", "Angry", "Calm", "Peaceful", "Turbulent", "Quiet", "Noisy", "Bright",
	"Dark", "Light", "Heavy", "Small", "Large", "Giant", "Tiny", "Fast", "Slow", "Quick", "Lazy",
	"Eager", "Reluctant", "Brave", "Cowardly", "Loyal", "Faithful", "True", "False", "Real", "Fake",
	"Solid", "Liquid", "Gas", "Sharp", "Blunt", "Hot", "Cold", "Warm", "Cool", "Clean", "Dirty",
	"Rich", "Poor", "Wealthy", "Needy", "Generous", "Selfish", "Kind", "Cruel", "Gentle", "Rough",
	"Polite", "Rude", "Wise", "Foolish", "Clever", "Stupid", "Strong", "Weak", "Healthy", "Sick",
	"Alive", "Dead", "Young", "Old", "New", "Ancient", "Modern", "Early", "Late", "First", "Last",
	"Best", "Worst", "Easy", "Hard", "Simple", "Complex", "Possible", "Impossible", "Visible", "Hidden",
	"Clear", "Obscure", "Valid", "Invalid", "Legal", "Illegal", "Safe", "Dangerous", "Public", "Private",
	"Open", "Closed", "Free", "Bound", "Active", "Passive", "Patient", "Impatient", "Honest", "Dishonest",
	"Silent", "Vocal", "Vivid", "Dull", "Perfect", "Flawed", "Unique", "Commonplace", "Original", "Copied",
	"Natural", "Artificial", "Organic", "Synthetic", "Local", "Foreign", "Global", "Domestic", "Wild", "Tame",
	"Creative", "Destructive", "Productive", "Useless", "Necessary", "Optional", "Similar", "Different",
	"United", "Divided", "Connected", "Isolated", "Included", "Excluded", "Forward", "Backward", "Upward", "Downward",
	"Inner", "Outer", "Central", "Peripheral", "Vertical", "Horizontal", "Diagonal", "Straight", "Curved",
	"Round", "Square", "Triangular", "Circular", "Rectangular", "Oval", "Spiral", "Linear", "Angular",
	"Parallel", "Perpendicular", "Symmetrical", "Asymmetrical", "Regular", "Irregular", "Ordered", "Chaotic",
	"Consistent", "Variable", "Constant", "Changing", "Stable", "Unstable", "Fixed", "Flexible", "Rigid",
	"Elastic", "Fluid", "Fragile", "Tough", "Resilient", "Vulnerable", "Delicate", "Robust", "Solid", "Hollow",
	"Dense", "Sparse", "Deep", "Shallow", "High", "Low", "Wide", "Narrow", "Long", "Short", "Tall", "Flat",
	"Thick", "Thin", "Full", "Empty", "Complete", "Incomplete", "Whole", "Part", "Many", "Few", "Several",
	"Numerous", "Multiple", "Single", "Double", "Triple", "Zero", "One", "Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine", "Ten", "Hundred", "Thousand", "Million", "Billion", "Trillion",
	"First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth", "Tenth",
	"Eleventh", "Twelfth", "Thirteenth", "Fourteenth", "Fifteenth", "Sixteenth", "Seventeenth", "Eighteenth", "Nineteenth", "Twentieth",
	"Twentieth", "Thirtieth", "Fortieth", "Fiftieth", "Sixtieth", "Seventieth", "Eightieth", "Ninetieth", "Hundredth",
	"Thousandth", "Millionth", "Billionth", "Trillionth", "Once", "Twice", "Thrice", "Again", "Often", "Seldom",
	"Always", "Never", "Sometimes", "Usually", "Rarely", "Frequently", "Occasionally", "Constantly", "Regularly", "Irregularly",
	"Generally", "Specifically", "Particularly", "Notably", "Especially", "Primarily", "Mainly", "Chiefly", "Largely", "Mostly",
	"Overall", "Entirely", "Completely", "Totally", "Fully", "Perfectly", "Absolutely", "Utterly", "Purely", "Merely",
	"Simply", "Just", "Only", "Exactly", "Precisely", "Approximately", "Roughly", "Nearly", "Closely", "Slightly",
	"Somewhat", "Rather", "Quite", "Very", "Extremely", "Highly", "Deeply", "Strongly", "Weakly", "Mildly", "Gently",
	"Softly", "Hardly", "Barely"}

const (
	WordPassword = iota
	PhrasePassword
	WordWithSpecial
	PhraseWithSpecial
	SecurePassword
	maxWordList = 100 // Can be expanded
)

func GeneratePassword(passwordLength int, passwordType int) (string, string, error) {
	if passwordLength <= 0 {
		return "", "", errors.New("password length must be greater than 0")
	}

	switch passwordType {
	case WordPassword:
		return generateWordPassword(passwordLength)
	case PhrasePassword:
		return generatePhrasePassword(passwordLength)
	case WordWithSpecial:
		return generateWordWithSpecial(passwordLength)
	case PhraseWithSpecial:
		return generatePhraseWithSpecial(passwordLength)
	case SecurePassword:
		return generateSecurePassword(passwordLength)
	default:
		return "", "", errors.New("invalid password type")
	}
}

func generateWordPassword(length int) (string, string, error) {
	var password strings.Builder
	var hint strings.Builder

	for i := 0; i < length; i++ {
		word := getRandomWord()
		password.WriteString(word)
		hint.WriteString(word)
		if i < length-1 {
			hint.WriteString(" ")
		}
	}
	return password.String(), hint.String(), nil
}

func generatePhrasePassword(length int) (string, string, error) {
	var password strings.Builder
	var hint strings.Builder

	for i := 0; i < length; i++ {
		word := getRandomWord()
		password.WriteString(word)
		hint.WriteString(word)
		if i < length-1 {
			password.WriteString("-")
			hint.WriteString(" ")
		}
	}
	return password.String(), hint.String(), nil
}

func generateWordWithSpecial(length int) (string, string, error) {
	password, hint, err := generateWordPassword(length)
	if err != nil {
		return "", "", err
	}
	return addSpecialCharacters(password), hint, nil
}

func generatePhraseWithSpecial(length int) (string, string, error) {
	password, hint, err := generatePhrasePassword(length)
	if err != nil {
		return "", "", err
	}
	return addSpecialCharacters(password), hint, nil
}

func generateSecurePassword(length int) (string, string, error) {
	const charset = LOWER_CASE + UPPER_CASE + DIGITS + SPECIAL_CHARS
	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomIndex, _ := crRand.Int(crRand.Reader, big.NewInt(int64(len(charset))))
		sb.WriteByte(charset[randomIndex.Int64()])
	}
	return sb.String(), "", nil // No hint for secure passwords
}

func getRandomWord() string {
	randomIndex, _ := crRand.Int(crRand.Reader, big.NewInt(int64(len(wordList))))
	return wordList[randomIndex.Int64()]
}

func addSpecialCharacters(s string) string {
	var result strings.Builder
	for _, char := range s {
		result.WriteRune(char)
		if rand.IntN(10) < 3 { // 30% chance to add a special character
			randomIndex, _ := crRand.Int(crRand.Reader, big.NewInt(int64(len(SPECIAL_CHARS))))
			result.WriteByte(SPECIAL_CHARS[randomIndex.Int64()])
		}
		if rand.IntN(10) < 3 && char >= '0' && char <= '9' {
			// 30% chance to replace number to special char
			randomIndex, _ := crRand.Int(crRand.Reader, big.NewInt(int64(len(SPECIAL_CHARS))))
			result.WriteByte(SPECIAL_CHARS[randomIndex.Int64()])
		}
		if rand.IntN(10) < 3 && char >= 'A' && char <= 'Z' {
			// 30% chance to replace number to special char
			randomIndex, _ := crRand.Int(crRand.Reader, big.NewInt(int64(len(SPECIAL_CHARS))))
			result.WriteByte(SPECIAL_CHARS[randomIndex.Int64()])
		}

	}
	return result.String()
}
