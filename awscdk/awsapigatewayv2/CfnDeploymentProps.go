package awsapigatewayv2


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	ApiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	StageName: jsii.String("stageName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html
//
type CfnDeploymentProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The description for the deployment resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of an existing stage to associate with the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-stagename
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

