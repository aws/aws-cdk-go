package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties of a new hosted zone.
//
// Example:
//   hostedZone := route53.NewHostedZone(this, jsii.String("MyHostedZone"), &HostedZoneProps{
//   	ZoneName: jsii.String("example.org"),
//   })
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("AWS/Route53"),
//   	MetricName: jsii.String("DNSQueries"),
//   	DimensionsMap: map[string]*string{
//   		"HostedZoneId": hostedZone.hostedZoneId,
//   	},
//   })
//
type HostedZoneProps struct {
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
	// Whether to add a trailing dot to the zone name.
	// Default: true.
	//
	AddTrailingDot *bool `field:"optional" json:"addTrailingDot" yaml:"addTrailingDot"`
	// Any comments that you want to include about the hosted zone.
	// Default: none.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	// Default: disabled.
	//
	QueryLogsLogGroupArn *string `field:"optional" json:"queryLogsLogGroupArn" yaml:"queryLogsLogGroupArn"`
	// A VPC that you want to associate with this hosted zone.
	//
	// When you specify
	// this property, a private hosted zone will be created.
	//
	// You can associate additional VPCs to this private zone using `addVpc(vpc)`.
	// Default: public (no VPCs associated).
	//
	Vpcs *[]awsec2.IVpc `field:"optional" json:"vpcs" yaml:"vpcs"`
}

