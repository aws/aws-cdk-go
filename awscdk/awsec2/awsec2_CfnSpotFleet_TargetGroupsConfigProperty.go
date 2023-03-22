package awsec2


// Describes the target groups to attach to a Spot Fleet.
//
// Spot Fleet registers the running Spot Instances with these target groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupsConfigProperty := &targetGroupsConfigProperty{
//   	targetGroups: []interface{}{
//   		&targetGroupProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   }
//
type CfnSpotFleet_TargetGroupsConfigProperty struct {
	// One or more target groups.
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

