package awsapigateway


// Properties for defining a `CfnBasePathMappingV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBasePathMappingV2Props := &CfnBasePathMappingV2Props{
//   	DomainNameArn: jsii.String("domainNameArn"),
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	BasePath: jsii.String("basePath"),
//   	Stage: jsii.String("stage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html
//
type CfnBasePathMappingV2Props struct {
	// The ARN of the domain name for the BasePathMappingV2 resource to be described.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-domainnamearn
	//
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// The private API's identifier.
	//
	// This identifier is unique across all of your APIs in API Gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The base path name that callers of the private API must provide as part of the URL after the domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// Represents a unique identifier for a version of a deployed private RestApi that is callable by users.
	//
	// The Stage must depend on the `RestApi` 's stage. To create a dependency, add a DependsOn attribute to the BasePathMappingV2 resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

