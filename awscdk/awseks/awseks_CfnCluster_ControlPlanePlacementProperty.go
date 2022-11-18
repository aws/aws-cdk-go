package awseks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlPlanePlacementProperty := &controlPlanePlacementProperty{
//   	groupName: jsii.String("groupName"),
//   }
//
type CfnCluster_ControlPlanePlacementProperty struct {
	// `CfnCluster.ControlPlanePlacementProperty.GroupName`.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
}

