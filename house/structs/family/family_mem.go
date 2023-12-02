package family

type Member struct{
	Name string
	Title string
	Age int16
	Occupation string
}

type Family struct{
	Family []Member
}

func addFamily Family{
	var family []Member

	family = append(family, Member{Name: "Cago", Title: "Brother", Age: 20, Occupation: "Plumber"})
	family = append(family, Member{Name: "Kawasaki", Title: "Brother", Age: 19, Occupation: "Racer"})
	family = append(family, Member{Name: "Krico", Title: "Brother", Age: 21, Occupation: "Smoker"})

	return Family{Family: family}
}
