package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var awsIam interface{}
//   var none interface{}
//
//   harnessAgentCoreGatewayConfigProperty := &HarnessAgentCoreGatewayConfigProperty{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   		AwsIam: awsIam,
//   		None: none,
//   		Oauth: &OAuthCredentialProviderProperty{
//   			CustomParameters: map[string]*string{
//   				"customParametersKey": jsii.String("customParameters"),
//   			},
//   			DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   			GrantType: jsii.String("grantType"),
//   			ProviderArn: jsii.String("providerArn"),
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoregatewayconfig.html
//
type CfnHarnessPropsMixin_HarnessAgentCoreGatewayConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoregatewayconfig.html#cfn-bedrockagentcore-harness-harnessagentcoregatewayconfig-gatewayarn
	//
	GatewayArn *string `field:"optional" json:"gatewayArn" yaml:"gatewayArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoregatewayconfig.html#cfn-bedrockagentcore-harness-harnessagentcoregatewayconfig-outboundauth
	//
	OutboundAuth interface{} `field:"optional" json:"outboundAuth" yaml:"outboundAuth"`
}

