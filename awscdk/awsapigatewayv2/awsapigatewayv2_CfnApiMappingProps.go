package awsapigatewayv2


// Properties for defining a `CfnApiMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiMappingProps := &cfnApiMappingProps{
//   	apiId: jsii.String("apiId"),
//   	domainName: jsii.String("domainName"),
//   	stage: jsii.String("stage"),
//
//   	// the properties below are optional
//   	apiMappingKey: jsii.String("apiMappingKey"),
//   }
//
type CfnApiMappingProps struct {
	// The identifier of the API.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The API stage.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The API mapping key.
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
}

