package glplus

import (
	"fmt"
	"image"
	"io"
	"strings"

	"github.com/aubonbeurre/go-obj/obj"
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
	Marker       Bounds
}

// ObjOptions ...
type ObjOptions struct {
	TexImg *image.RGBA
	Colors map[string]float32
	Single bool
}

// LoadObj ...
// 'colors' relate to usemtl
func LoadObj(f io.Reader, opts *ObjOptions) (objs []*Obj, err error) {

	var o *obj.Object
	o, err = obj.NewReader(f).Read()
	if err != nil {
		return nil, err
	}

	var findColor = func(faceIndex int) (float32, error) {
		for _, material := range o.SubMaterials {
			if faceIndex <= material.FaceEndIndex {
				if material.Name == "None" {
					return 0, nil
				}
				if col, ok := opts.Colors[material.Name]; ok {
					return col, nil
				}
				return 0, fmt.Errorf("Unknown material %s", material.Name)
			}
		}
		return 0, fmt.Errorf("Unknown error")
	}

	// convert our object into cube vertices for opengl
	var startFaceIndex int
	for _, sub := range o.Subobjects {
		var objVertices []float32
		var builder BoundBuilder
		builder.reset()
		HasUVs := false

		for faceIndex := startFaceIndex; faceIndex < sub.FaceEndIndex; faceIndex++ {
			f := &o.Faces[faceIndex]
			var faceColorPacked float32
			if opts.Colors != nil {
				if faceColorPacked, err = findColor(faceIndex); err != nil {
					return nil, err
				}
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

			if v1.Texture != nil && opts.TexImg != nil {
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

		if HasUVs && opts.TexImg != nil {
			rgba = opts.TexImg
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

	if opts.Single && len(objs) > 1 {
		var newobjs []*Obj

		for ind, o := range objs {
			if ind == 0 {
				newobjs = append(newobjs, o)
			} else if strings.HasPrefix(o.Name, "marker") {
				newobjs[0].Marker = o.Bounds
			} else {
				newobjs[0].Bounds = newobjs[0].Bounds.Union(o.Bounds)
				newobjs[0].ObjVertices = append(newobjs[0].ObjVertices, o.ObjVertices...)
			}
		}

		objs = newobjs
	}

	return objs, nil
}
