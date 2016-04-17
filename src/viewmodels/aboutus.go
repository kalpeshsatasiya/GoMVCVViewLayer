package viewmodels

type AboutUs struct  {
	Title string
	Active string
}

func GetAboutUs() AboutUs  {
	result := AboutUs{
		Title: "This is an AboutUs page",
		Active:"aboutus",
	}

	return  result
}