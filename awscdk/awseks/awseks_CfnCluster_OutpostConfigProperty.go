package awseks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outpostConfigProperty := &outpostConfigProperty{
//   	controlPlaneInstanceType: jsii.String("controlPlaneInstanceType"),
//   	outpostArns: []*string{
//   		jsii.String("outpostArns"),
//   	},
//
//   	// the properties below are optional
//   	controlPlanePlacement: &controlPlanePlacementProperty{
//   		groupName: jsii.String("groupName"),
//   	},
//   }
//
type CfnCluster_OutpostConfigProperty struct {
	// `CfnCluster.OutpostConfigProperty.ControlPlaneInstanceType`.
	ControlPlaneInstanceType *string `field:"required" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// `CfnCluster.OutpostConfigProperty.OutpostArns`.
	OutpostArns *[]*string `field:"required" json:"outpostArns" yaml:"outpostArns"`
	// `CfnCluster.OutpostConfigProperty.ControlPlanePlacement`.
	ControlPlanePlacement interface{} `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
}

