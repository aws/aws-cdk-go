package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to create a Route 53 private hosted zone.
//
// Example:
//   var vpc vpc
//
//
//   zone := route53.NewPrivateHostedZone(this, jsii.String("HostedZone"), &PrivateHostedZoneProps{
//   	ZoneName: jsii.String("fully.qualified.domain.com"),
//   	Vpc: Vpc,
//   })
//
type PrivateHostedZoneProps struct {
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
	// Private hosted zones must be associated with at least one VPC. You can
	// associated additional VPCs using `addVpc(vpc)`.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

