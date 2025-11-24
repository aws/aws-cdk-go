package mixinsawsapigatewayv2


// Properties for CfnApiMappingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApiMappingMixinProps := &CfnApiMappingMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	ApiMappingKey: jsii.String("apiMappingKey"),
//   	DomainName: jsii.String("domainName"),
//   	Stage: jsii.String("stage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html
//
type CfnApiMappingMixinProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The API mapping key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apimappingkey
	//
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
	// The domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The API stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

