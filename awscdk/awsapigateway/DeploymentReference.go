package awsapigateway


// A reference to a Deployment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReference := &DeploymentReference{
//   	DeploymentId: jsii.String("deploymentId"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type DeploymentReference struct {
	// The DeploymentId of the Deployment resource.
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// The RestApiId of the Deployment resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

