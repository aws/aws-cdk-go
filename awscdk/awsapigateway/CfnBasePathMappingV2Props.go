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
	// The Arn of an AWS::ApiGateway::DomainNameV2 resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-domainnamearn
	//
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// The ID of the API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The base path name that callers of the API must provide in the URL after the domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The name of the API's stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmappingv2.html#cfn-apigateway-basepathmappingv2-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

