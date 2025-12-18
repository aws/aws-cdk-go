package awsroute53


// The type of health check to be associated with the record.
//
// Example:
//   var myZone HostedZone
//
//
//   healthCheck := route53.NewHealthCheck(this, jsii.String("HealthCheck"), &HealthCheckProps{
//   	Type: route53.HealthCheckType_HTTP,
//   	Fqdn: jsii.String("example.com"),
//   	Port: jsii.Number(80),
//   	ResourcePath: jsii.String("/health"),
//   	FailureThreshold: jsii.Number(3),
//   	RequestInterval: awscdk.Duration_Seconds(jsii.Number(30)),
//   })
//
//   route53.NewARecord(this, jsii.String("ARecordFailoverPrimary"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
//   	Failover: route53.Failover_PRIMARY,
//   	HealthCheck: HealthCheck,
//   	SetIdentifier: jsii.String("failover-primary"),
//   })
//
//   route53.NewARecord(this, jsii.String("ARecordFailoverSecondary"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("5.6.7.8")),
//   	Failover: route53.Failover_SECONDARY,
//   	SetIdentifier: jsii.String("failover-secondary"),
//   })
//
type HealthCheckType string

const (
	// HTTP health check.
	//
	// Route 53 tries to establish a TCP connection. If successful, Route 53 submits an HTTP request and waits for an HTTP status code of 200 or greater and less than 400.
	HealthCheckType_HTTP HealthCheckType = "HTTP"
	// HTTPS health check.
	//
	// Route 53 tries to establish a TCP connection. If successful, Route 53 submits an HTTPS request and waits for an HTTP status code of 200 or greater and less than 400.
	HealthCheckType_HTTPS HealthCheckType = "HTTPS"
	// HTTP health check with string matching.
	//
	// Route 53 tries to establish a TCP connection. If successful, Route 53 submits an HTTP request and searches the first 5,120 bytes of the response body for the string that you specify in SearchString.
	HealthCheckType_HTTP_STR_MATCH HealthCheckType = "HTTP_STR_MATCH"
	// HTTPS health check with string matching.
	//
	// Route 53 tries to establish a TCP connection. If successful, Route 53 submits an HTTPS request and searches the first 5,120 bytes of the response body for the string that you specify in SearchString.
	HealthCheckType_HTTPS_STR_MATCH HealthCheckType = "HTTPS_STR_MATCH"
	// TCP health check.
	//
	// Route 53 tries to establish a TCP connection.
	HealthCheckType_TCP HealthCheckType = "TCP"
	// CloudWatch metric health check.
	//
	// The health check is associated with a CloudWatch alarm. If the state of the alarm is OK, the health check is considered healthy. If the state is ALARM, the health check is considered unhealthy. If CloudWatch doesn't have sufficient data to determine whether the state is OK or ALARM, the health check status depends on the setting for InsufficientDataHealthStatus: Healthy, Unhealthy, or LastKnownStatus.
	HealthCheckType_CLOUDWATCH_METRIC HealthCheckType = "CLOUDWATCH_METRIC"
	// Calculated health check.
	//
	// For health checks that monitor the status of other health checks, Route 53 adds up the number of health checks that Route 53 health checkers consider to be healthy and compares that number with the value of HealthThreshold.
	HealthCheckType_CALCULATED HealthCheckType = "CALCULATED"
	// Recovery control health check.
	//
	// The health check is associated with a Route53 Application Recovery Controller routing control. If the routing control state is ON, the health check is considered healthy. If the state is OFF, the health check is considered unhealthy.
	HealthCheckType_RECOVERY_CONTROL HealthCheckType = "RECOVERY_CONTROL"
)

