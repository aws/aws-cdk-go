package awsroute53


// The failover policy.
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
// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-failover.html
//
type Failover string

const (
	// The primary resource record set determines how Route 53 responds to DNS queries when the primary resource is healthy.
	Failover_PRIMARY Failover = "PRIMARY"
	// The secondary resource record set determines how Route 53 responds to DNS queries when the primary resource is unhealthy.
	Failover_SECONDARY Failover = "SECONDARY"
)

