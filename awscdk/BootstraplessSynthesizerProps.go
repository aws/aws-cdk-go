package awscdk


// Construction properties of `BootstraplessSynthesizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstraplessSynthesizerProps := &BootstraplessSynthesizerProps{
//   	CloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	DeployRoleArn: jsii.String("deployRoleArn"),
//   }
//
type BootstraplessSynthesizerProps struct {
	// The CFN execution Role ARN to use.
	// Default: - No CloudFormation role (use CLI credentials).
	//
	CloudFormationExecutionRoleArn *string `field:"optional" json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The deploy Role ARN to use.
	// Default: - No deploy role (use CLI credentials).
	//
	DeployRoleArn *string `field:"optional" json:"deployRoleArn" yaml:"deployRoleArn"`
}

