package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomainNameV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var managementPolicy interface{}
//   var policy interface{}
//
//   cfnDomainNameV2Props := &CfnDomainNameV2Props{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DomainName: jsii.String("domainName"),
//   	EndpointConfiguration: &EndpointConfigurationProperty{
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   	ManagementPolicy: managementPolicy,
//   	Policy: policy,
//   	SecurityPolicy: jsii.String("securityPolicy"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html
//
type CfnDomainNameV2Props struct {
	// The reference to an AWS -managed certificate that will be used by the private endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Represents a custom domain name as a user-friendly host name of an API (RestApi).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The endpoint configuration to indicate the types of endpoints an API (RestApi) or its custom domain name (DomainName) has.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-endpointconfiguration
	//
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-managementpolicy
	//
	ManagementPolicy interface{} `field:"optional" json:"managementPolicy" yaml:"managementPolicy"`
	// A stringified JSON policy document that applies to the `execute-api` service for this DomainName regardless of the caller and Method configuration.
	//
	// You can use `Fn::ToJsonString` to enter your `policy` . For more information, see [Fn::ToJsonString](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ToJsonString.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	//
	// Only `TLS_1_2` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-securitypolicy
	//
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html#cfn-apigateway-domainnamev2-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

