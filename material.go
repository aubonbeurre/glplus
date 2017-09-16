package glplus

// Material ...
type Material struct {
	Diffuse   []float32 `yaml:",flow"`
	Shininess float32   `yaml:"shininess"` // rename
	Specular  []float32 `yaml:",flow"`
	Ambient   []float32 `yaml:",flow"`
}

// MakeDefaultMaterial ...
func MakeDefaultMaterial() *Material {
	return &Material{
		Diffuse:   []float32{1, 0, 0, 1},
		Shininess: 100,
		Specular:  []float32{1, 1, 1, 1},
		Ambient:   []float32{0.2, 0.2, 0.2, 0.1},
	}
}
