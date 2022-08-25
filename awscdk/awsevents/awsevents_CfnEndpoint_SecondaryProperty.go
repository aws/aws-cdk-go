package awsevents


// The secondary Region that processes events when failover is triggered or replication is enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryProperty := &secondaryProperty{
//   	route: jsii.String("route"),
//   }
//
type CfnEndpoint_SecondaryProperty struct {
	// Defines the secondary Region.
	Route *string `field:"required" json:"route" yaml:"route"`
}

