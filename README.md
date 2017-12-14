# gin-bindata-loadstatic
gin use go-bindata load static source
## useage
use [go-bindata](https://github.com/jteeuwen/go-bindata) generate assets.go

use [gin web framework](https://github.com/gin-gonic/gin) load assets resource

in main.go
```
import (
    static "github.com/youngbloood/gin-bindata/loadstatic"
    "github.com/gin-gonic/gin"
    "assets"
)

func init(){
    engin:=gin.Default()
    static.NewAssetsFS(engine, assets.AssetDir, assets.Asset, "").LoadStatic()
}
```