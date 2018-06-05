package cars

import (
	"context"
	"log"

	carssvc "goa.design/goa/examples/cars/gen/cars"
)

// cars service example implementation.
// The example methods log the requests and return zero values.
type carsSvc struct {
	logger *log.Logger
}

// NewCars returns the cars service implementation.
func NewCars(logger *log.Logger) carssvc.Service {
	return &carsSvc{logger}
}

// Lists car models by body style.
func (s *carsSvc) List(ctx context.Context, p *carssvc.ListPayload, l carssvc.ListServerStream) (err error) {
	for _, c := range modelsByStyle[p.Style] {
		if err := l.Send(c); err != nil {
			return err
		}
	}
	return l.Close()
}

var modelsByStyle = map[string][]*carssvc.Car{
	"sedan": []*carssvc.Car{
		&carssvc.Car{Make: "Acura", Model: "TLX", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Audi", Model: "A4", BodyStyle: "sedan"},
		&carssvc.Car{Make: "BMW", Model: "M3", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Chevrolet", Model: "Cruze", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Ford", Model: "Focus", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Honda", Model: "Accord", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Hyundai", Model: "Accent", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Infiniti", Model: "Q50", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Kia", Model: "Rio", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Lexus", Model: "ES", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Mazda", Model: "6", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Mercedes", Model: "C-Class", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Nissan", Model: "Altima", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Porsche", Model: "Panamera", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Subaru", Model: "Impreza", BodyStyle: "sedan"},
		&carssvc.Car{Make: "Volkswagen", Model: "Passat", BodyStyle: "sedan"},
	},
	"hatchback": []*carssvc.Car{
		&carssvc.Car{Make: "Acura", Model: "MDX", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Audi", Model: "Q3", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "BMW", Model: "X3", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Chevrolet", Model: "Equinox", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Ford", Model: "Escape", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Honda", Model: "CRV", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Hyundai", Model: "Santa Fe", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Infiniti", Model: "QX30", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Kia", Model: "Sorento", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Lexus", Model: "NX", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Mazda", Model: "CX5", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Mercedes", Model: "GLA-Class", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Nissan", Model: "Rogue", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Porsche", Model: "Cayenne", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Subaru", Model: "Outback", BodyStyle: "hatchback"},
		&carssvc.Car{Make: "Volkswagen", Model: "Golf", BodyStyle: "hatchback"},
	},
}
