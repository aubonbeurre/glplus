package glplus

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"path"

	"github.com/aubonbeurre/go-obj/obj"
)

// Obj ...
type Obj struct {
	Bounds      Bounds
	ObjVertices []float32
	Name        string
	TexName     string
	TexImg      *image.RGBA
}

// LoadObj ...
func LoadObj(filename string, texname string) (objs []*Obj, err error) {

	var f *os.File
	f, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	var o *obj.Object
	o, err = obj.NewReader(f).Read()
	if err != nil {
		return nil, err
	}

	// convert our object into cube vertices for opengl
	var startFaceIndex int
	for _, sub := range o.Subobjects {
		var objVertices []float32
		var builder BoundBuilder
		builder.reset()
		HasUVs := false

		for faceIndex := startFaceIndex; faceIndex < sub.FaceStartIndex; faceIndex++ {
			f := &o.Faces[faceIndex]

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
			if v1.Texture != nil {
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
			}

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
		}

		var rgba *image.RGBA

		if HasUVs && texname != "" {
			imgPath := path.Join(path.Dir(filename), texname)

			var imgFile *os.File
			if imgFile, err = os.Open(imgPath); err != nil {
				return nil, err
			}
			defer imgFile.Close()

			var img image.Image
			if img, _, err = image.Decode(imgFile); err != nil {
				return nil, err
			}

			rgba = image.NewRGBA(img.Bounds())
			if rgba.Stride != rgba.Rect.Size().X*4 {
				return nil, fmt.Errorf("unsupported stride")
			}
			draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
		}

		newobj := &Obj{
			ObjVertices: objVertices,
			Name:        sub.Name,
			Bounds:      builder.build(),
			TexName:     texname,
			TexImg:      rgba,
		}

		objs = append(objs, newobj)
		startFaceIndex = sub.FaceStartIndex
		fmt.Printf("%s, bounds=%v, center=%v, uvs=%v\n", newobj.Name, newobj.Bounds, newobj.Bounds.Center(), HasUVs)
	}
	return objs, nil
}
