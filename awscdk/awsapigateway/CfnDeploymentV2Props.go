package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Deployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentV2Props := &CfnDeploymentV2Props{
//   	ApiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	StageName: jsii.String("stageName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnDeploymentV2Props struct {
	// `AWS::ApiGatewayV2::Deployment.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::Deployment.Description`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-description
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::ApiGatewayV2::Deployment.StageName`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html#cfn-apigatewayv2-deployment-stagename
	//
	// Deprecated: moved to package aws-apigatewayv2.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

