package viewmodels

type  Index struct {
	Title  string
	Active string
}

func GetIndex() Index  {
	result := Index{
		Title:"This is a Demo Application",
		Active: "index",
	}

	return  result
}