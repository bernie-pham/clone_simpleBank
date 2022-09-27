package ultilities

import (
	"math/rand"
	"strings"
	"time"
)

var (
	alphaChars = "qwertyuiopasdfghjklzxcvbnm"
	numChars   = "1234567890"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var rb strings.Builder
	k := len(alphaChars)

	for i := 0; i < n; i++ {
		c := alphaChars[rand.Intn(k)]
		rb.WriteByte(c)
	}
	return rb.String()
}

func RandomCurrency() string {
	currencies := []string{"VND", "USD", "EUR", "REN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomMoney() int64 {
	return RandomInt(100, 10000)
}

func RandomFullName() string {
	firstName := []string{"David", "Cyrus", "Vladimir", "Bernie", "Flo", "Daniel", "Lucas", "Helsinki"}
	middleName := []string{"Naun", "Mountain", "Vir"}
	lastName := []string{"Falk", "Rob", "Mil", "Wolf", "Hillan"}

	var fullname strings.Builder
	fullname.WriteString(firstName[rand.Intn(len(firstName))])
	fullname.WriteString(middleName[rand.Intn(len(middleName))])
	fullname.WriteString(lastName[rand.Intn(len(lastName))])

	return fullname.String()
}
