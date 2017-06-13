package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("withpayloads", func() {
	Title("With Payloads")
	Description("Failing test case that demonstrates the problems introduced by 2098c07b944cd4dd951de695b242ef582c3fc211")
	Version("1.0.0")

	Host("localhost:8010")
	Scheme("http")

	Consumes("application/json")
	Produces("application/json")
})

var Greeting = Type("Greeting", func() {
	Attribute("id", Integer, func() {
		Description("A required int field in the parent type.")
	})
	Attribute("message", String, func() {
		Description("A required string field in the parent type.")
	})
	Attribute("parent_optional", Boolean, func() {
		Description("An optional boolean field in the parent type.")
	})
	Required("id", "message")
})

var ExtendedGreeting = Type("ExtendedGreeting", func() {
	Reference(Greeting)

	Attribute("id")
	Attribute("message")
	Attribute("parent_optional")

	Attribute("active", Boolean, func() {
		Description("A required boolean field in the child type.")
	})
	Attribute("optional", Integer, func() {
		Description("An optional integer field in the child type.")
	})
	Attribute("created_at", DateTime, func() {
		Description("A required date field in the child type.")
	})

	Required("id", "message", "active", "created_at")
})

var GreetingMedia = MediaType("application/vnd.io.bluecanvas.helloworld.greeting.v1+json", func() {
	TypeName("GreetingMedia")
	Reference(Greeting)

	Attributes(func() {
		Attribute("id")
		Attribute("message")
		Attribute("parent_optional")
		Attribute("href", String, func() {
			Description("A required string field in the response media type.")
		})
		Required("id", "message", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("message")
		Attribute("parent_optional")
		Attribute("href")
	})
})

var ExtendedGreetingMedia = MediaType("application/vnd.io.bluecanvas.helloworld.extended-greeting.v1+json", func() {
	TypeName("ExtendedGreetingMedia")
	Reference(ExtendedGreeting)

	Attributes(func() {
		Attribute("id")
		Attribute("message")
		Attribute("parent_optional")
		Attribute("active")
		Attribute("optional")
		Attribute("created_at")
		Attribute("href", String, func() {
			Description("A required string field in the response media type.")
		})
		Required("id", "message", "active", "created_at", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("message")
		Attribute("parent_optional")
		Attribute("active")
		Attribute("optional")
		Attribute("created_at")
		Attribute("href")
	})
})

var _ = Resource("Greeting", func() {
	Description("The greeting service provides greetings so everybody has a pleasant day.")
	DefaultMedia(GreetingMedia)

	Action("show", func() {
		Routing(
			GET("/"))
		Response(OK, GreetingMedia)
		Response(BadRequest)
	})

	Action("create", func() {
		Routing(
			POST("/"))
		Payload(func() {
			Member("message")
			Member("parent_optional")
			Required("message")
		})
		Response(Created)
		Response(BadRequest)
	})
})

var _ = Resource("ExtendedGreeting", func() {
	Description("The greeting service provides greetings so everybody has a pleasant day.")
	DefaultMedia(ExtendedGreetingMedia)
	BasePath("/extended")

	Action("show", func() {
		Routing(
			GET("/"))
		Response(OK, ExtendedGreetingMedia)
		Response(BadRequest)
	})

	Action("create", func() {
		Routing(
			POST("/"))
		Payload(func() {
			Member("message")
			Member("parent_optional")
			Member("active")
			Member("optional")
			Required("message", "active")
		})
		Response(Created)
		Response(BadRequest)
	})

	Action("createOptional", func() {
		Routing(
			POST("/"))
		OptionalPayload(func() {
			Member("message")
			Member("parent_optional")
			Member("active")
			Member("optional")
			Required("message", "active")
		})
		Response(Created)
		Response(BadRequest)
	})

	Action("createDirect", func() {
		Routing(
			POST("/"))
		Payload(ExtendedGreeting)
		Response(Created)
		Response(BadRequest)
	})
})
