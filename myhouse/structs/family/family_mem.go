package family

import (
	"fmt"
	"strings"
)

type Member struct {
	Name      string
	Title     string
	Age       int16
	Occupation string
}

type Family struct {
	Family []Member
}

func AddFamily() Family {
	var family []Member

	family = append(family, Member{Name: "Cago", Title: "Brother", Age: 20, Occupation: "Plumber"})
	family = append(family, Member{Name: "Kawasaki", Title: "Brother", Age: 19, Occupation: "Racer"})
	family = append(family, Member{Name: "Krico", Title: "Brother", Age: 21, Occupation: "Smoker"})

	return Family{Family: family}
}

func (f Family) String() string {
	var familyStrings []string

	for _, member := range f.Family {
		familyStrings = append(familyStrings, member.String())
	}

	return "\n" + strings.Join(familyStrings, "\n")
}

func (m Member) String() string {
	return fmt.Sprintf("Name: %s, Title: %s, Age: %d, Occupation: %s",
		m.Name, m.Title, m.Age, m.Occupation)
}