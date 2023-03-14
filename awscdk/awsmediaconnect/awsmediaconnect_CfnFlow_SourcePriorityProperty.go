package awsmediaconnect


// The priority you want to assign to a source.
//
// You can have a primary stream and a backup stream or two equally prioritized streams. This setting only applies when Failover Mode is set to FAILOVER.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourcePriorityProperty := &sourcePriorityProperty{
//   	primarySource: jsii.String("primarySource"),
//   }
//
type CfnFlow_SourcePriorityProperty struct {
	// The name of the source you choose as the primary source for this flow.
	PrimarySource *string `field:"required" json:"primarySource" yaml:"primarySource"`
}

