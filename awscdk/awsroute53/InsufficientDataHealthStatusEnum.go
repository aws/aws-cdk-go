package awsroute53


// The status of the health check when CloudWatch has insufficient data about the state of associated alarm.
type InsufficientDataHealthStatusEnum string

const (
	// Route 53 health check status will be healthy.
	InsufficientDataHealthStatusEnum_HEALTHY InsufficientDataHealthStatusEnum = "HEALTHY"
	// Route 53 health check status will be unhealthy.
	InsufficientDataHealthStatusEnum_UNHEALTHY InsufficientDataHealthStatusEnum = "UNHEALTHY"
	// Route 53 health check status will be the status of the health check before Route 53 had insufficient data.
	InsufficientDataHealthStatusEnum_LAST_KNOWN_STATUS InsufficientDataHealthStatusEnum = "LAST_KNOWN_STATUS"
)

