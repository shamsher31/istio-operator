package versions

import "k8s.io/apimachinery/pkg/util/sets"

const (
	DiscoveryChart       = "istio-discovery"
	GatewayEgressChart   = "istio-egress"
	GatewayIngressChart  = "istio-ingress"
	GrafanaChart         = "grafana"
	KialiChart           = "kiali"
	MeshConfigChart      = "mesh-config"
	MixerPolicyChart     = "istio-policy"
	MixerTelemetryChart  = "mixer-telemetry"
	PrometheusChart      = "prometheus"
	TelemetryCommonChart = "telemetry-common"
	ThreeScaleChart      = "maistra-threescale"
	TracingChart         = "tracing"
	WASMExtensionsChart  = "wasm-extensions"
	RLSChart             = "rls"

	// Event reasons
	eventReasonRendering = "Rendering"
)

type chartRenderingDetails struct {
	path         string
	enabledField string
}

var specialCharts = sets.NewString(GatewayIngressChart, GatewayEgressChart, ThreeScaleChart)
