package awsevents


// The secondary Region that processes events when failover is triggered or replication is enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryProperty := &SecondaryProperty{
//   	Route: jsii.String("route"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-secondary.html
//
type CfnEndpoint_SecondaryProperty struct {
	// Defines the secondary Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-secondary.html#cfn-events-endpoint-secondary-route
	//
	Route *string `field:"required" json:"route" yaml:"route"`
}

