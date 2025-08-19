package pokeapi

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []struct {
		BaseStat int
		Effort   int
		Stat     struct {
			Name string
			Url  string
		}
	}
	Types []struct {
		Slot int
		Type struct {
			Name string
			Url  string
		}
	}
}
