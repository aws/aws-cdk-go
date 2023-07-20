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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-auth.html
//
type CfnFunction_AuthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-auth.html#cfn-serverless-function-auth-apikeyrequired
	//
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-auth.html#cfn-serverless-function-auth-authorizationscopes
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-auth.html#cfn-serverless-function-auth-authorizer
	//
	Authorizer *string `field:"optional" json:"authorizer" yaml:"authorizer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-auth.html#cfn-serverless-function-auth-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
}

