package awscodedeploy


// Information about the instances that belong to the replacement environment in a blue/green deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greenFleetProvisioningOptionProperty := &GreenFleetProvisioningOptionProperty{
//   	Action: jsii.String("action"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-greenfleetprovisioningoption.html
//
type CfnDeploymentGroup_GreenFleetProvisioningOptionProperty struct {
	// The method used to add instances to a replacement environment.
	//
	// - `DISCOVER_EXISTING` : Use instances that already exist or will be created manually.
	// - `COPY_AUTO_SCALING_GROUP` : Use settings from a specified Auto Scaling group to define and create instances in a new Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-greenfleetprovisioningoption.html#cfn-codedeploy-deploymentgroup-greenfleetprovisioningoption-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
}

