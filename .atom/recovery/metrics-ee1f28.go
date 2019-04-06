package listener

import "github.com/prometheus/client_golang/prometheus"

func init() {
	registerPrometheusMetrics()
}

func registerPrometheusMetrics() {
	prometheus.MustRegister(nodeIsSafe)
	prometheus.MustRegister(nodeAction)
}

var (
	nodeIsSafe = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "k8s_spot_termination_handler_node_is_safe",
			Help: "If the termination notice for node is not found, emit 1. Else 0",
		},
		[]string{"instanceId", "instanceType", "availabilityZone", "localHostname"},
	)

	nodeAction = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "k8s_spot_termination_handler_node_action",
			Help: "If an action is scheduled for a node (stop, hibernate, terminate), emit 1. Else 0",
		},
		[]string{"nodeAction", "actionTime"})
)
