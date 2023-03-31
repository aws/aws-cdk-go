package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customStatements interface{}
//
//   authProperty := &authProperty{
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizer: jsii.String("authorizer"),
//   	resourcePolicy: &authResourcePolicyProperty{
//   		awsAccountBlacklist: []*string{
//   			jsii.String("awsAccountBlacklist"),
//   		},
//   		awsAccountWhitelist: []*string{
//   			jsii.String("awsAccountWhitelist"),
//   		},
//   		customStatements: []interface{}{
//   			customStatements,
//   		},
//   		intrinsicVpcBlacklist: []*string{
//   			jsii.String("intrinsicVpcBlacklist"),
//   		},
//   		intrinsicVpceBlacklist: []*string{
//   			jsii.String("intrinsicVpceBlacklist"),
//   		},
//   		intrinsicVpceWhitelist: []*string{
//   			jsii.String("intrinsicVpceWhitelist"),
//   		},
//   		intrinsicVpcWhitelist: []*string{
//   			jsii.String("intrinsicVpcWhitelist"),
//   		},
//   		ipRangeBlacklist: []*string{
//   			jsii.String("ipRangeBlacklist"),
//   		},
//   		ipRangeWhitelist: []*string{
//   			jsii.String("ipRangeWhitelist"),
//   		},
//   		sourceVpcBlacklist: []*string{
//   			jsii.String("sourceVpcBlacklist"),
//   		},
//   		sourceVpcWhitelist: []*string{
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

