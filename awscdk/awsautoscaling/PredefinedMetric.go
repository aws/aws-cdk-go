package awsautoscaling


// One of the predefined autoscaling metrics.
type PredefinedMetric string

const (
	// Average CPU utilization of the Auto Scaling group.
	PredefinedMetric_ASG_AVERAGE_CPU_UTILIZATION PredefinedMetric = "ASG_AVERAGE_CPU_UTILIZATION"
	// Average number of bytes received on all network interfaces by the Auto Scaling group.
	PredefinedMetric_ASG_AVERAGE_NETWORK_IN PredefinedMetric = "ASG_AVERAGE_NETWORK_IN"
	// Average number of bytes sent out on all network interfaces by the Auto Scaling group.
	PredefinedMetric_ASG_AVERAGE_NETWORK_OUT PredefinedMetric = "ASG_AVERAGE_NETWORK_OUT"
	// Number of requests completed per target in an Application Load Balancer target group.
	//
	// Specify the ALB to look at in the `resourceLabel` field.
	PredefinedMetric_ALB_REQUEST_COUNT_PER_TARGET PredefinedMetric = "ALB_REQUEST_COUNT_PER_TARGET"
)

