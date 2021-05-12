package structs

type Region struct {
	Center          []float32 `yaml:"center"`
	DescriptionID   uint32    `yaml:"descriptionID"`
	Max             []float32 `yaml:"max"`
	Min             []float32 `yaml:"min"`
	NameID          uint32    `yaml:"nameID"`
	Nebula          uint32    `yaml:"nebula"`
	RegionID        uint32    `yaml:"regionID"`
	WormholeClassID uint32    `yaml:"wormholeClassID"`
}

type Constellation struct {
	Center          []float32 `yaml:"center"`
	ConstellationID uint32    `yaml:"constellationID"`
	Max             []float32 `yaml:"max"`
	Min             []float32 `yaml:"min"`
	NameID          uint32    `yaml:"nameID"`
	Radius          float64   `yaml:"radius"`
}

type SolarSystem struct {
	Border        bool      `yaml:"border"`
	Center        []float32 `yaml:"center"`
	Corridor      bool      `yaml:"corridor"`
	Fringe        bool      `yaml:"fringe"`
	Hub           bool      `yaml:"hub"`
	International bool      `yaml:"international"`
	Luminosity    float64   `yaml:"luminosity"`
	Max           []float32 `yaml:"max"`
	Min           []float32 `yaml:"min"`
	NameID        uint32    `yaml:"nameID"`
	Radius        float64   `yaml:"radius"`
}
