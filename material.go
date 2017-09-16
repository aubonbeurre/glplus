package glplus

// Material ...
type Material struct {
	Diffuse   []float32 `yaml:",flow"`
	Shininess float32   `yaml:"shininess"` // rename
	Specular  []float32 `yaml:",flow"`
	Ambient   []float32 `yaml:",flow"`
}
