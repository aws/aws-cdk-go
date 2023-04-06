package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   authProperty := &AuthProperty{
//   	ApiKeyRequired: jsii.Boolean(false),
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	Authorizer: jsii.String("authorizer"),
//   	ResourcePolicy: &AuthResourcePolicyProperty{
//   		AwsAccountBlacklist: []*string{
//   			jsii.String("awsAccountBlacklist"),
//   		},
//   		AwsAccountWhitelist: []*string{
//   			jsii.String("awsAccountWhitelist"),
//   		},
//   		CustomStatements: []interface{}{
//   			customStatements,
//   		},
//   		IntrinsicVpcBlacklist: []*string{
//   			jsii.String("intrinsicVpcBlacklist"),
//   		},
//   		IntrinsicVpceBlacklist: []*string{
//   			jsii.String("intrinsicVpceBlacklist"),
//   		},
//   		IntrinsicVpceWhitelist: []*string{
//   			jsii.String("intrinsicVpceWhitelist"),
//   		},
//   		IntrinsicVpcWhitelist: []*string{
//   			jsii.String("intrinsicVpcWhitelist"),
//   		},
//   		IpRangeBlacklist: []*string{
//   			jsii.String("ipRangeBlacklist"),
//   		},
//   		IpRangeWhitelist: []*string{
//   			jsii.String("ipRangeWhitelist"),
//   		},
//   		SourceVpcBlacklist: []*string{
//   			jsii.String("sourceVpcBlacklist"),
//   		},
//   		SourceVpcWhitelist: []*string{
//   			jsii.String("sourceVpcWhitelist"),
//   		},
//   	},
//   }
//
type CfnFunction_AuthProperty struct {
	// `CfnFunction.AuthProperty.ApiKeyRequired`.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// `CfnFunction.AuthProperty.AuthorizationScopes`.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// `CfnFunction.AuthProperty.Authorizer`.
	Authorizer *string `field:"optional" json:"authorizer" yaml:"authorizer"`
	// `CfnFunction.AuthProperty.ResourcePolicy`.
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
}

