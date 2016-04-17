package viewmodels

type Pricing struct {
	Title             string
	Active            string
	HostingCategories []HostingCategory
}

type HostingCategory struct {
	Name        string
	Description string
	Price       float64
}

func GetPricing() Pricing {
	result := Pricing{
		Title: "This is a Pricing page",
		Active: "pricing",
	}

	sharedHosting := HostingCategory{
		Name: "SHARED HOSTING",
		Description:"Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. Nunc at viverra risus.",
		Price:199,
	}

	vpsHosting := HostingCategory{
		Name: "VPS HOSTING",
		Description:"Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. Nunc at viverra risus.",
		Price:299,
	}

	dedicatedHosting := HostingCategory{
		Name: "DEDICATED HOSTING",
		Description:"Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. Nunc at viverra risus.",
		Price:599,
	}

	windowsHosting := HostingCategory{
		Name: "WINDOWS HOSTING",
		Description:"Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. In euismod quam ac dictum varius. Nunc at viverra risus. Nunc at viverra risus.",
		Price:99,
	}

	result.HostingCategories = []HostingCategory{
		sharedHosting,
		vpsHosting,
		dedicatedHosting,
		windowsHosting,
	}

	return result
}