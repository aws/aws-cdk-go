package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::ApiMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiMappingV2Props := &CfnApiMappingV2Props{
//   	ApiId: jsii.String("apiId"),
//   	DomainName: jsii.String("domainName"),
//   	Stage: jsii.String("stage"),
//
//   	// the properties below are optional
//   	ApiMappingKey: jsii.String("apiMappingKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnApiMappingV2Props struct {
	// `AWS::ApiGatewayV2::ApiMapping.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::ApiMapping.DomainName`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-domainname
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// `AWS::ApiGatewayV2::ApiMapping.Stage`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-stage
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// `AWS::ApiGatewayV2::ApiMapping.ApiMappingKey`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apimappingkey
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
}

