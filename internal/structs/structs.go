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
	Border            bool      `yaml:"border"`
	Center            []float32 `yaml:"center"`
	Corridor          bool      `yaml:"corridor"`
	Fringe            bool      `yaml:"fringe"`
	Hub               bool      `yaml:"hub"`
	International     bool      `yaml:"international"`
	Luminosity        float64   `yaml:"luminosity"`
	Max               []float32 `yaml:"max"`
	Min               []float32 `yaml:"min"`
	Planets           struct{}  `yaml:"planets"`
	Radius            float64   `yaml:"radius"`
	Regional          bool      `yaml:"regional"`
	Security          float64   `yaml:"security"`
	SecurityClass     string    `yaml:"securityClass"`
	SolarSystemID     uint32    `yaml:"solarSystemID"`
	SolarSystemNameID uint32    `yaml:"solarSystemNameID"`
	NameID            uint32    `yaml:"nameID"`
	Star              struct{}  `yaml:"star"`
	Stargates         map[uint32]struct {
		Destination uint32    `yaml:"destination"`
		Position    []float32 `yaml:"position"`
		TypeID      uint32    `yaml:"typeID"`
	} `yaml:"stargates"`
	SunTypeID       uint32 `yaml:"sunTypeID"`
	WormholeClassID uint32 `yaml:"wormholeClassID"`
}
