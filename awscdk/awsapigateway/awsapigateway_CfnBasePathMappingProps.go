package awsapigateway


// Properties for defining a `CfnBasePathMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBasePathMappingProps := &cfnBasePathMappingProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	id: jsii.String("id"),
//   	restApiId: jsii.String("restApiId"),
//   	stage: jsii.String("stage"),
//   }
//
type CfnBasePathMappingProps struct {
	// The `DomainName` of an [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide in the URL after the domain name.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// `AWS::ApiGateway::BasePathMapping.Id`.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the API.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The name of the API's stage.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

