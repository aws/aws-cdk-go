package awsec2


// Describes a load balancer target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupProperty := &targetGroupProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnSpotFleet_TargetGroupProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

