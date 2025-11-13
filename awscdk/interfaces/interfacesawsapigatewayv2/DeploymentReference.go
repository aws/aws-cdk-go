package interfacesawsapigatewayv2


// A reference to a Deployment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReference := &DeploymentReference{
//   	ApiId: jsii.String("apiId"),
//   	DeploymentId: jsii.String("deploymentId"),
//   }
//
type DeploymentReference struct {
	// The ApiId of the Deployment resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The DeploymentId of the Deployment resource.
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
}

