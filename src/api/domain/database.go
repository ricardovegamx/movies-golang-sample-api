package domain

var Database = []Movie{
	{
		Id:          1,
		Title:       "Birds of Prey (and the Fantabulous Emancipation of One Harley Quinn)",
		ReleaseDate: "02/07/2020",
		Categories:  []string{"Action", "Crime", "Comedy"},
		Runtime:     109,
		UserScore:   .72,
		Overview:    "Harley Quinn joins forces with a singer, an assassin and a police detective to help a young girl who had a hit placed on her after she stole a rare diamond from a crime lord.",
		Director:    "Cathy Yan",
	},
	{
		Id:          2,
		Title:       "Beauty and the Beast",
		ReleaseDate: "03/17/2017",
		Categories:  []string{"Family", "Fantasy", "Romance"},
		Runtime:     129,
		UserScore:   .7,
		Overview:    "A live-action adaptation of Disney's version of the classic tale of a cursed prince and a beautiful young woman who helps him break the spell.",
		Director:    "Bill Condon",
	},
	{
		Id:          3,
		Title:       "Batman v Superman: Dawn of Justice",
		ReleaseDate: "03/25/2016",
		Categories:  []string{"Action", "Adventure", "Fantasy"},
		Runtime:     151,
		UserScore:   .59,
		Overview:    "Fearing the actions of a god-like Super Hero left unchecked, Gotham City’s own formidable, forceful vigilante takes on Metropolis’s most revered, modern-day savior, while the world wrestles with what sort of hero it really needs. And with Batman and Superman at war with one another, a new threat quickly arises, putting mankind in greater danger than it’s ever known before.",
		Director:    "Zack Snyder",
	},
}
