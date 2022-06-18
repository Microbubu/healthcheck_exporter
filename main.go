package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	serviceStatusGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_status_gauge",
			Help: "Servic status monitor gauge",
		},
		[]string{"ip_address", "service_name"},
	)

	serviceCallerCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_caller_counter",
			Help: "call service times counter",
		},
		[]string{"url"},
	)
)

func main() {
	config, err := ReadYamlConfig()
	if err != nil {
		log.Fatal(err)
	}

	prometheus.MustRegister(serviceStatusGauge)
	prometheus.MustRegister(serviceCallerCounter)

	runTasks(config.Tasks)

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(config.Domain, nil))
}

func runTasks(tasks []CallerTask) {
	for _, task := range tasks {
		task := task
		go func() {
			for {
				caller, err := NewCaller(task.Type)
				if err != nil {
					log.Fatal(err)
				}

				caller.Call(&CallerRequest{
					Url:        task.Url,
					HttpMethod: task.HttpMethod,
				})

				serviceCallerCounter.With(prometheus.Labels{"url": task.Url}).Inc()
				serviceStatusGauge.WithLabelValues(task.Url, task.Name).Set(1.0)

				time.Sleep(time.Duration(task.Interval * int(time.Second)))
			}
		}()
	}
}
