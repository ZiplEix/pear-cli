package initfiles

type replacementApi struct {
	ImportApiLib string
}

type replacementOrm struct {
	ImportOrmLib string
}

var FiberReplacement = replacementApi{
	ImportApiLib: "github.com/gofiber/fiber/v2",
}

var GinReplacement = replacementApi{
	ImportApiLib: "github.com/gin-gonic/gin",
}

var GormReplacement = replacementOrm{
	ImportOrmLib: "gorm.io/gorm",
}

var XormReplacement = replacementOrm{
	ImportOrmLib: "xorm.io/xorm",
}
