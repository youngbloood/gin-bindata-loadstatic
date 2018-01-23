# gin-bindata-loadstatic
gin use go-bindata load static source

## prepare
use [go-bindata](https://github.com/jteeuwen/go-bindata) to generate assets.go
## purpose


in [gin web framework](https://github.com/gin-gonic/gin)   project then we can load assets resource from assets.go
## useage

> go get github.com/youngbloood/gin-bindata/loadstatic

in main.go
```
import (
    static "github.com/youngbloood/gin-bindata/loadstatic"
    "github.com/gin-gonic/gin"
    "assets"
)

func init(){
    engine := gin.Default()
    static.NewAssetsFS(engine, assets.AssetDir, assets.Asset, "").LoadStatic()
}
```