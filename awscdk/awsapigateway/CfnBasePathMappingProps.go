package awsapigateway


// Properties for defining a `CfnBasePathMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBasePathMappingProps := &CfnBasePathMappingProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	BasePath: jsii.String("basePath"),
//   	Id: jsii.String("id"),
//   	RestApiId: jsii.String("restApiId"),
//   	Stage: jsii.String("stage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html
//
type CfnBasePathMappingProps struct {
	// The domain name of the BasePathMapping resource to be described.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide as part of the URL after the domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-restapiid
	//
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The name of the associated stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

