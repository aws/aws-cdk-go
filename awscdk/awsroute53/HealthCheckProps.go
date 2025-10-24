package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a new health check.
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
//   route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
//   	HealthCheck: HealthCheck,
//   	Weight: jsii.Number(100),
//   })
//   route53.NewARecord(this, jsii.String("ARecord2"), &ARecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("5.6.7.8")),
//   	Weight: jsii.Number(0),
//   })
//
type HealthCheckProps struct {
	// The type of health check to be associated with the record.
	Type HealthCheckType `field:"required" json:"type" yaml:"type"`
	// CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.
	// Default: - if the type is CLOUDWATCH_METRIC, this property is required. Otherwise, it is not configured.
	//
	AlarmIdentifier *AlarmIdentifier `field:"optional" json:"alarmIdentifier" yaml:"alarmIdentifier"`
	// A list of health checks to monitor for this 'CALCULATED' health check.
	// Default: - if the type is CALCULATED, this property is required. Otherwise, it is not configured.
	//
	ChildHealthChecks *[]IHealthCheck `field:"optional" json:"childHealthChecks" yaml:"childHealthChecks"`
	// Specify whether you want Amazon Route 53 to send the value of FullyQualifiedDomainName to the endpoint in the client_hello message during TLS negotiation.
	//
	// This allows the endpoint to respond to HTTPS health check requests with the applicable SSL/TLS certificate.
	// Default: - if the type is HTTPS or HTTPS_STR_MATCH, this property default value is true. Otherwise, it is not configured.
	//
	EnableSNI *bool `field:"optional" json:"enableSNI" yaml:"enableSNI"`
	// The number of consecutive health checks that an endpoint must pass or fail for Amazon Route 53 to change the current status of the endpoint from unhealthy to healthy or vice versa.
	// Default: - if the type is CALCULATED it's not configured
	// - if the type is CLOUDWATCH_METRIC it's not configured
	// - otherwise, the default value is 3.
	//
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Fully qualified domain name of the endpoint to be checked.
	//
	// Amazon Route 53 behavior depends on whether you specify a value for IPAddress.
	//
	// If you specify a value for IPAddress:
	//
	// Amazon Route 53 sends health check requests to the specified IPv4 or IPv6 address and passes the value of FullyQualifiedDomainName in the Host header for all health checks except TCP health checks. This is typically the fully qualified DNS name of the endpoint on which you want Route 53 to perform health checks.
	// Note: If you specify a value for Port property other than 80 or 443, Route 53 will construct the value for Host header as FullyQualifiedDomainName:Port.
	//
	// If you don't specify a value for IPAddress:
	//
	// Route 53 sends a DNS request to the domain that you specify for FullyQualifiedDomainName at the interval that you specify for RequestInterval. Using an IPv4 address that DNS returns, Route 53 then checks the health of the endpoint.
	//
	// Additionally, if the type of the health check is HTTP, HTTPS, HTTP_STR_MATCH, or HTTPS_STR_MATCH, Route 53 passes the value of FullyQualifiedDomainName in the Host header, as it does when you specify value for IPAddress. If the type is TCP, Route 53 doesn't pass a Host header.
	// Default: - not configured.
	//
	Fqdn *string `field:"optional" json:"fqdn" yaml:"fqdn"`
	// The number of child health checks that are associated with a CALCULATED health that Amazon Route 53 must consider healthy for the CALCULATED health check to be considered healthy.
	// Default: - if the type is CALCULATED, the default value is number of child health checks. Otherwise, it is not configured.
	//
	HealthThreshold *float64 `field:"optional" json:"healthThreshold" yaml:"healthThreshold"`
	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm.
	// Default: - if the type is CLOUDWATCH_METRIC, the default value is InsufficientDataHealthStatus.LAST_KNOWN_STATUS. Otherwise, it is not configured.
	//
	InsufficientDataHealthStatus InsufficientDataHealthStatusEnum `field:"optional" json:"insufficientDataHealthStatus" yaml:"insufficientDataHealthStatus"`
	// Specify whether you want Amazon Route 53 to invert the status of a health check, so a health check that would normally be considered unhealthy is considered healthy, and vice versa.
	// Default: false.
	//
	Inverted *bool `field:"optional" json:"inverted" yaml:"inverted"`
	// The IPv4 or IPv6 IP address for the endpoint that you want Amazon Route 53 to perform health checks on.
	//
	// If you don't specify a value for IPAddress, Route 53 sends a DNS request to resolve the domain name that you specify in FullyQualifiedDomainName at the interval that you specify in RequestInterval. Using an IPv4 address that DNS returns, Route 53 then checks the health of the endpoint.
	// Default: - not configured.
	//
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Specify whether you want Amazon Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint, and to display CloudWatch latency graphs on the Health Checks page in the Route 53 console.
	// Default: - if the type is CALCULATED it's not configured
	// - if the type is CLOUDWATCH_METRIC it's not configured
	// - otherwise, the default value is false.
	//
	MeasureLatency *bool `field:"optional" json:"measureLatency" yaml:"measureLatency"`
	// The port on the endpoint that you want Amazon Route 53 to perform health checks on.
	// Default: - if the type is HTTP or HTTP_STR_MATCH, the default value is 80.
	// - if the type is HTTPS or HTTPS_STR_MATCH, the default value is 443.
	// - otherwise, it is not configured.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An array of region identifiers that you want Amazon Route 53 health checkers to check the health of the endpoint from.
	//
	// Please refer to the CloudFormation documentation for the most up-to-date list of regions.
	// See: https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html
	//
	// Default: - if the type is CALCULATED, CLOUDWATCH_METRIC, or RECOVERY_CONTROL, this property is not configured.
	// - otherwise, the default value will be set by CloudFormation itself and will include all valid regions. Please refer to the CloudFormation documentation for the most up-to-date list of regions.
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// The duration between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health check request.
	//
	// Each Route 53 health checker makes requests at this interval.
	// Default: - if the type is CALCULATED it's not configured
	// - if the type is CLOUDWATCH_METRIC it's not configured
	// - otherwise, the default value is 30 seconds.
	//
	RequestInterval awscdk.Duration `field:"optional" json:"requestInterval" yaml:"requestInterval"`
	// The path that you want Amazon Route 53 to request when performing health checks.
	//
	// The path can be any value for which your endpoint will return an HTTP status code of 2xx or 3xx when the endpoint is healthy, for example the file /docs/route53-health-check.html. Route 53 automatically adds the DNS name for the service and a leading forward slash (/) character.
	// Default: - if the type is HTTP, HTTPS, HTTP_STR_MATCH, or HTTPS_STR_MATCH, the default value is empty string.
	// - otherwise, it is not configured.
	//
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// The Amazon Resource Name (ARN) of the Route 53 Application Recovery Controller routing control that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.
	// Default: - if the type is RECOVERY_CONTROL, this property is required. Otherwise, it is not configured.
	//
	RoutingControl *string `field:"optional" json:"routingControl" yaml:"routingControl"`
	// The string that you want Amazon Route 53 to search for in the response body from the specified resource.
	//
	// If the string appears in the response body, Route 53 considers the resource healthy.
	//
	// Route 53 considers case when searching for SearchString in the response body.
	// Default: - if the type is HTTP_STR_MATCH or HTTPS_STR_MATCH, this property is required. Otherwise, it is not configured.
	//
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
}

