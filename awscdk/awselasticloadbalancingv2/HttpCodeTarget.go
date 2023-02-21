package awselasticloadbalancingv2


// Count of HTTP status originating from the targets.
type HttpCodeTarget string

const (
	// The number of 2xx response codes from targets.
	HttpCodeTarget_TARGET_2XX_COUNT HttpCodeTarget = "TARGET_2XX_COUNT"
	// The number of 3xx response codes from targets.
	HttpCodeTarget_TARGET_3XX_COUNT HttpCodeTarget = "TARGET_3XX_COUNT"
	// The number of 4xx response codes from targets.
	HttpCodeTarget_TARGET_4XX_COUNT HttpCodeTarget = "TARGET_4XX_COUNT"
	// The number of 5xx response codes from targets.
	HttpCodeTarget_TARGET_5XX_COUNT HttpCodeTarget = "TARGET_5XX_COUNT"
)

