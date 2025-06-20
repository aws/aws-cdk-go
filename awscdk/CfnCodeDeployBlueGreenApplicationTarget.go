package awscdk


// Type of the `CfnCodeDeployBlueGreenApplication.target` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenApplicationTarget := &CfnCodeDeployBlueGreenApplicationTarget{
//   	LogicalId: jsii.String("logicalId"),
//   	Type: jsii.String("type"),
//   }
//
type CfnCodeDeployBlueGreenApplicationTarget struct {
	// The logical id of the target resource.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The resource type of the target being deployed.
	//
	// Right now, the only allowed value is 'AWS::ECS::Service'.
	Type *string `field:"required" json:"type" yaml:"type"`
}

