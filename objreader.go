package glplus

import (
	"fmt"
	"image"
	"io"

	"github.com/aubonbeurre/go-obj/obj"
	"github.com/go-gl/mathgl/mgl32"
)

// SubObject ...
type SubObject struct {
	Name         string
	FaceEndIndex int
}

// SubMaterial ...
type SubMaterial struct {
	Name         string
	FaceEndIndex int
}

// Obj ...
type Obj struct {
	Bounds       Bounds
	ObjVertices  []float32
	Name         string
	TexImg       *image.RGBA
	SubObjects   []SubObject
	SubMaterials []SubMaterial
}

// LoadObj ...
// 'colors' relate to usemtl
func LoadObj(f io.Reader, png *image.RGBA, colors map[string]mgl32.Vec3) (objs []*Obj, err error) {

	var o *obj.Object
	o, err = obj.NewReader(f).Read()
	if err != nil {
		return nil, err
	}

	var findColor = func(faceIndex int) mgl32.Vec3 {
		for _, material := range o.SubMaterials {
			if faceIndex <= material.FaceEndIndex {
				if col, ok := colors[material.Name]; ok {
					return col
				}
				panic(fmt.Errorf("Unknown material %s", material.Name))
			}
		}
		panic(fmt.Errorf("Unknown error"))
	}

	// convert our object into cube vertices for opengl
	var startFaceIndex int
	for _, sub := range o.Subobjects {
		var objVertices []float32
		var builder BoundBuilder
		builder.reset()
		HasUVs := false

		/*var unpackColor = func(f float64) mgl32.Vec3 {
			var color mgl32.Vec3
			color[2] = float32(math.Floor(f / 256.0 / 256.0))
			color[1] = float32(math.Floor((f - float64(color[2])*256.0*256.0) / 256.0))
			color[0] = float32(math.Floor(f - float64(color[2])*256.0*256.0 - float64(color[1])*256.0))
			// now we have a vec3 with the 3 components in range [0..255]. Let's normalize it!
			return color.Mul(1 / 255.0)
		}*/

		for faceIndex := startFaceIndex; faceIndex < sub.FaceEndIndex; faceIndex++ {
			f := &o.Faces[faceIndex]
			var faceColorPacked float32
			if colors != nil {
				faceColor := findColor(faceIndex)
				faceColorPacked = faceColor[0]*255.0 + faceColor[1]*255.0*256.0 + faceColor[2]*255.0*256.0*256.0
				//fmt.Printf("%f %v\n", faceColorPacked, unpackColor(float64(faceColorPacked)))
			}

			has4 := len(f.Points) == 4
			var v1, v2, v3, v4 *obj.Point
			v1 = f.Points[0]
			v2 = f.Points[1]
			v3 = f.Points[2]
			builder.include64(v1.Vertex.X, v1.Vertex.Y, v1.Vertex.Z)
			builder.include64(v2.Vertex.X, v2.Vertex.Y, v2.Vertex.Z)
			builder.include64(v3.Vertex.X, v3.Vertex.Y, v3.Vertex.Z)
			if has4 {
				v4 = f.Points[3]
				builder.include64(v4.Vertex.X, v4.Vertex.Y, v4.Vertex.Z)
			}

			var u, v [4]float64

			if v1.Texture != nil && png != nil {
				HasUVs = true
				u[0] = v1.Texture.U
				v[0] = v1.Texture.V
				u[1] = v2.Texture.U
				v[1] = v2.Texture.V
				u[2] = v3.Texture.U
				v[2] = v3.Texture.V
				if has4 {
					u[3] = v4.Texture.U
					v[3] = v4.Texture.V
				}
			} else if HasUVs {
				return nil, fmt.Errorf("Inconsistent UVs")
			}

			if HasUVs {
				objVertices = append(objVertices,
					[]float32{
						float32(v1.Vertex.X), float32(v1.Vertex.Y), float32(v1.Vertex.Z), float32(u[0]), float32(v[0]),
						float32(v1.Normal.X), float32(v1.Normal.Y), float32(v1.Normal.Z),
						float32(v2.Vertex.X), float32(v2.Vertex.Y), float32(v2.Vertex.Z), float32(u[1]), float32(v[1]),
						float32(v2.Normal.X), float32(v2.Normal.Y), float32(v2.Normal.Z),
						float32(v3.Vertex.X), float32(v3.Vertex.Y), float32(v3.Vertex.Z), float32(u[2]), float32(v[2]),
						float32(v3.Normal.X), float32(v3.Normal.Y), float32(v3.Normal.Z),
					}...)

				if has4 {
					objVertices = append(objVertices,
						[]float32{
							float32(v3.Vertex.X), float32(v3.Vertex.Y), float32(v3.Vertex.Z), float32(u[2]), float32(v[2]),
							float32(v3.Normal.X), float32(v3.Normal.Y), float32(v3.Normal.Z),
							float32(v4.Vertex.X), float32(v4.Vertex.Y), float32(v4.Vertex.Z), float32(u[3]), float32(v[3]),
							float32(v4.Normal.X), float32(v4.Normal.Y), float32(v4.Normal.Z),
							float32(v1.Vertex.X), float32(v1.Vertex.Y), float32(v1.Vertex.Z), float32(u[0]), float32(v[0]),
							float32(v1.Normal.X), float32(v1.Normal.Y), float32(v1.Normal.Z),
						}...)
				}
			} else {
				objVertices = append(objVertices,
					[]float32{
						float32(v1.Vertex.X), float32(v1.Vertex.Y), float32(v1.Vertex.Z), faceColorPacked,
						float32(v1.Normal.X), float32(v1.Normal.Y), float32(v1.Normal.Z),
						float32(v2.Vertex.X), float32(v2.Vertex.Y), float32(v2.Vertex.Z), faceColorPacked,
						float32(v2.Normal.X), float32(v2.Normal.Y), float32(v2.Normal.Z),
						float32(v3.Vertex.X), float32(v3.Vertex.Y), float32(v3.Vertex.Z), faceColorPacked,
						float32(v3.Normal.X), float32(v3.Normal.Y), float32(v3.Normal.Z),
					}...)

				if has4 {
					objVertices = append(objVertices,
						[]float32{
							float32(v3.Vertex.X), float32(v3.Vertex.Y), float32(v3.Vertex.Z), faceColorPacked,
							float32(v3.Normal.X), float32(v3.Normal.Y), float32(v3.Normal.Z),
							float32(v4.Vertex.X), float32(v4.Vertex.Y), float32(v4.Vertex.Z), faceColorPacked,
							float32(v4.Normal.X), float32(v4.Normal.Y), float32(v4.Normal.Z),
							float32(v1.Vertex.X), float32(v1.Vertex.Y), float32(v1.Vertex.Z), faceColorPacked,
							float32(v1.Normal.X), float32(v1.Normal.Y), float32(v1.Normal.Z),
						}...)
				}
			}
		}

		var rgba *image.RGBA

		if HasUVs && png != nil {
			rgba = png
		}

		subobjects := make([]SubObject, 0)
		for _, subo := range o.Subobjects {
			subobjects = append(subobjects, SubObject{Name: subo.Name, FaceEndIndex: subo.FaceEndIndex})
		}

		materials := make([]SubMaterial, 0)
		for _, subm := range o.SubMaterials {
			materials = append(materials, SubMaterial{Name: subm.Name, FaceEndIndex: subm.FaceEndIndex})
		}

		newobj := &Obj{
			ObjVertices:  objVertices,
			Name:         sub.Name,
			Bounds:       builder.build(),
			TexImg:       rgba,
			SubObjects:   subobjects,
			SubMaterials: materials,
		}

		objs = append(objs, newobj)
		startFaceIndex = sub.FaceEndIndex
		//fmt.Printf("%s, bounds=%v, center=%v, uvs=%v\n", newobj.Name, newobj.Bounds, newobj.Bounds.Center(), HasUVs)
	}
	return objs, nil
}
