// info:{"Path":"github.com/dave/images","Hash":11461617162488775357}
package images

import (
	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
	"reflect"
)

// Automatically created basic rule for photo
type PhotoRule struct {
	*system.Object
	*system.Rule
}
type Photo struct {
	*system.Object
	Url *system.String `json:"url"`
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("github.com/dave/images", 11461617162488775357)
	pkg.InitType("photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
