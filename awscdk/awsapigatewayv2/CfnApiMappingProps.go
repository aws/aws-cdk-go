package awsapigatewayv2


// Properties for defining a `CfnApiMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiMappingProps := &CfnApiMappingProps{
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
type CfnApiMappingProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apiid
	//
	ApiId interface{} `field:"required" json:"apiId" yaml:"apiId"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-domainname
	//
	DomainName interface{} `field:"required" json:"domainName" yaml:"domainName"`
	// The API stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-stage
	//
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The API mapping key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apimappingkey
	//
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
}

