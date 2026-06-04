package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var awsIam interface{}
//   var none interface{}
//
//   harnessAgentCoreGatewayConfigProperty := &HarnessAgentCoreGatewayConfigProperty{
//   	GatewayArn: jsii.String("gatewayArn"),
//
//   	// the properties below are optional
//   	OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   		AwsIam: awsIam,
//   		None: none,
//   		Oauth: &OAuthCredentialProviderProperty{
//   			ProviderArn: jsii.String("providerArn"),
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//
//   			// the properties below are optional
//   			CustomParameters: map[string]*string{
//   				"customParametersKey": jsii.String("customParameters"),
//   			},
//   			DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   			GrantType: jsii.String("grantType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoregatewayconfig.html
//
type CfnHarness_HarnessAgentCoreGatewayConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoregatewayconfig.html#cfn-bedrockagentcore-harness-harnessagentcoregatewayconfig-gatewayarn
	//
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoregatewayconfig.html#cfn-bedrockagentcore-harness-harnessagentcoregatewayconfig-outboundauth
	//
	OutboundAuth interface{} `field:"optional" json:"outboundAuth" yaml:"outboundAuth"`
}

