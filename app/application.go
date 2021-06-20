package app

import (
	"microservices-go/log/option_a"
	"microservices-go/log/option_b"
)

func StartApp() {
	option_a.Log.Info("option a : about to map the urls")
	option_b.Log.Info("option b : about to map the urls")
	option_a.Info("option a : restart aplication", "step:1", "result:ok")
	option_b.Info("option b : restart aplication",
		option_b.Field("step", 1),
		option_b.Field("result", "ok"),
	)
}
