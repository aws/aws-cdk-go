package interfacesawslaunchwizard


// A reference to a Deployment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentReference := &DeploymentReference{
//   	DeploymentArn: jsii.String("deploymentArn"),
//   }
//
type DeploymentReference struct {
	// The Arn of the Deployment resource.
	DeploymentArn *string `field:"required" json:"deploymentArn" yaml:"deploymentArn"`
}

