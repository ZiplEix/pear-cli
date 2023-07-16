package initfiles

type replacementApi struct {
	ImportApiLib                 string
	ImportApiLibMiddlewareCors   string
	ImportApiLibMiddlewareLogger string
	ImportApiLibRecover          string
	AppCreation                  string
	ApiLogger                    string
	ApiCors                      string
	ApiListen                    string
}

type replacementOrm struct {
	ImportOrmLib string
}

var FiberReplacement = replacementApi{
	ImportApiLib:                 "github.com/gofiber/fiber/v2",
	ImportApiLibMiddlewareCors:   "github.com/gofiber/fiber/v2/middleware/cors",
	ImportApiLibMiddlewareLogger: "github.com/gofiber/fiber/v2/middleware/logger",
	ImportApiLibRecover:          "github.com/gofiber/fiber/v2/middleware/recover",
	AppCreation:                  "fiber.New()",
	ApiLogger:                    "logger.New()",
	ApiCors:                      "cors.New()",
	ApiListen:                    "Listen",
}

var GinReplacement = replacementApi{
	ImportApiLib:                 "github.com/gin-gonic/gin",
	ImportApiLibMiddlewareCors:   "github.com/gin-gonic/gin",
	ImportApiLibMiddlewareLogger: "github.com/gin-gonic/gin",
	ImportApiLibRecover:          "github.com/gin-gonic/gin",
	AppCreation:                  "gin.New()",
	ApiLogger:                    "gin.Logger()",
	ApiCors:                      "gin.Cors()",
	ApiListen:                    "Run",
}

var GormReplacement = replacementOrm{
	ImportOrmLib: "gorm.io/gorm",
}

var XormReplacement = replacementOrm{
	ImportOrmLib: "xorm.io/xorm",
}
