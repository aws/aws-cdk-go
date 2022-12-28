package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::DomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnDomainNameV2Props := &cfnDomainNameV2Props{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	domainNameConfigurations: []interface{}{
//   		&domainNameConfigurationProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			certificateName: jsii.String("certificateName"),
//   			endpointType: jsii.String("endpointType"),
//   		},
//   	},
//   	tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnDomainNameV2Props struct {
	// `AWS::ApiGatewayV2::DomainName.DomainName`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-domainname
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// `AWS::ApiGatewayV2::DomainName.DomainNameConfigurations`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-domainnameconfigurations
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DomainNameConfigurations interface{} `field:"optional" json:"domainNameConfigurations" yaml:"domainNameConfigurations"`
	// `AWS::ApiGatewayV2::DomainName.Tags`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-tags
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

