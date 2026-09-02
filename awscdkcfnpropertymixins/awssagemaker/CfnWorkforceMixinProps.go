package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkforcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkforceMixinProps := &CfnWorkforceMixinProps{
//   	CognitoConfig: &CognitoConfigProperty{
//   		ClientId: jsii.String("clientId"),
//   		UserPool: jsii.String("userPool"),
//   	},
//   	IpAddressType: jsii.String("ipAddressType"),
//   	OidcConfig: &OidcConfigProperty{
//   		AuthenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		Issuer: jsii.String("issuer"),
//   		JwksUri: jsii.String("jwksUri"),
//   		LogoutEndpoint: jsii.String("logoutEndpoint"),
//   		Scope: jsii.String("scope"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   	},
//   	SourceIpConfig: &SourceIpConfigProperty{
//   		Cidrs: []*string{
//   			jsii.String("cidrs"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkforceName: jsii.String("workforceName"),
//   	WorkforceVpcConfig: &WorkforceVpcConfigRequestProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html
//
type CfnWorkforceMixinProps struct {
	// The configuration of an Amazon Cognito workforce.
	//
	// A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-cognitoconfig
	//
	CognitoConfig interface{} `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// The IP address type for the workforce.
	//
	// IPv4 only or dualstack (IPv4 and IPv6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The configuration of an OIDC Identity Provider (IdP) private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-oidcconfig
	//
	OidcConfig interface{} `field:"optional" json:"oidcConfig" yaml:"oidcConfig"`
	// A list of IP address ranges used to access your training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-sourceipconfig
	//
	SourceIpConfig interface{} `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// An array of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-workforcename
	//
	WorkforceName *string `field:"optional" json:"workforceName" yaml:"workforceName"`
	// The VPC configuration for the workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html#cfn-sagemaker-workforce-workforcevpcconfig
	//
	WorkforceVpcConfig interface{} `field:"optional" json:"workforceVpcConfig" yaml:"workforceVpcConfig"`
}

