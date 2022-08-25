package awsapigatewayv2


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &cfnDeploymentProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnDeploymentProps struct {
	// The API identifier.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The description for the deployment resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of an existing stage to associate with the deployment.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

