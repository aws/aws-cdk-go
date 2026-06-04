package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var awsIam interface{}
//   var none interface{}
//
//   harnessGatewayOutboundAuthProperty := &HarnessGatewayOutboundAuthProperty{
//   	AwsIam: awsIam,
//   	None: none,
//   	Oauth: &OAuthCredentialProviderProperty{
//   		ProviderArn: jsii.String("providerArn"),
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//
//   		// the properties below are optional
//   		CustomParameters: map[string]*string{
//   			"customParametersKey": jsii.String("customParameters"),
//   		},
//   		DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   		GrantType: jsii.String("grantType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgatewayoutboundauth.html
//
type CfnHarness_HarnessGatewayOutboundAuthProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgatewayoutboundauth.html#cfn-bedrockagentcore-harness-harnessgatewayoutboundauth-awsiam
	//
	AwsIam interface{} `field:"optional" json:"awsIam" yaml:"awsIam"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgatewayoutboundauth.html#cfn-bedrockagentcore-harness-harnessgatewayoutboundauth-none
	//
	None interface{} `field:"optional" json:"none" yaml:"none"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgatewayoutboundauth.html#cfn-bedrockagentcore-harness-harnessgatewayoutboundauth-oauth
	//
	Oauth interface{} `field:"optional" json:"oauth" yaml:"oauth"`
}

