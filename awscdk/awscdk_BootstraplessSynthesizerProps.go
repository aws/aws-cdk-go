// An experiment to bundle the entire CDK into a single module
package awscdk


// Construction properties of {@link BootstraplessSynthesizer}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstraplessSynthesizerProps := &bootstraplessSynthesizerProps{
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	deployRoleArn: jsii.String("deployRoleArn"),
//   }
//
// Experimental.
type BootstraplessSynthesizerProps struct {
	// The CFN execution Role ARN to use.
	// Experimental.
	CloudFormationExecutionRoleArn *string `field:"optional" json:"cloudFormationExecutionRoleArn" yaml:"cloudFormationExecutionRoleArn"`
	// The deploy Role ARN to use.
	// Experimental.
	DeployRoleArn *string `field:"optional" json:"deployRoleArn" yaml:"deployRoleArn"`
}

