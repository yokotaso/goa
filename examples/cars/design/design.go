package design

import . "goa.design/goa/http/design"
import . "goa.design/goa/http/dsl"

var _ = API("cars", func() {
	Title("Cars Service")
	Description("HTTP service to lookup car models by body style.")
	Server("http://localhost:8080")
})

var Car = ResultType("application/vnd.goa.car", func() {
	TypeName("car")
	Attributes(func() {
		Attribute("make", String, "The make of the car")
		Attribute("model", String, "The car model")
		Attribute("body_style", String, "The car body style")
		Required("make", "model", "body_style")
	})
})

var _ = Service("cars", func() {
	HTTP(func() {
		Path("/cars")
	})

	Description("The cars service lists car models by body style.")
	Method("list", func() {
		Description("Lists car models by body style.")
		Payload(func() {
			Attribute("style", String, "The car body style.", func() {
				Enum("sedan", "hatchback")
			})
			Required("style")
		})
		//Result(CollectionOf(Car))
		StreamingResult(Car)
		HTTP(func() {
			GET("")
			Param("style")
		})
	})
})
