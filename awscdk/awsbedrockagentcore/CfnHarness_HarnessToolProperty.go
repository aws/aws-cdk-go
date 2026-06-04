package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var awsIam interface{}
//   var inputSchema interface{}
//   var none interface{}
//
//   harnessToolProperty := &HarnessToolProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Config: &HarnessToolConfigurationProperty{
//   		AgentCoreBrowser: &HarnessAgentCoreBrowserConfigProperty{
//   			BrowserArn: jsii.String("browserArn"),
//   		},
//   		AgentCoreCodeInterpreter: &HarnessAgentCoreCodeInterpreterConfigProperty{
//   			CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   		},
//   		AgentCoreGateway: &HarnessAgentCoreGatewayConfigProperty{
//   			GatewayArn: jsii.String("gatewayArn"),
//
//   			// the properties below are optional
//   			OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   				AwsIam: awsIam,
//   				None: none,
//   				Oauth: &OAuthCredentialProviderProperty{
//   					ProviderArn: jsii.String("providerArn"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//
//   					// the properties below are optional
//   					CustomParameters: map[string]*string{
//   						"customParametersKey": jsii.String("customParameters"),
//   					},
//   					DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   					GrantType: jsii.String("grantType"),
//   				},
//   			},
//   		},
//   		InlineFunction: &HarnessInlineFunctionConfigProperty{
//   			Description: jsii.String("description"),
//   			InputSchema: inputSchema,
//   		},
//   		RemoteMcp: &HarnessRemoteMcpConfigProperty{
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			Headers: map[string]*string{
//   				"headersKey": jsii.String("headers"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html
//
type CfnHarness_HarnessToolProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html#cfn-bedrockagentcore-harness-harnesstool-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html#cfn-bedrockagentcore-harness-harnesstool-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html#cfn-bedrockagentcore-harness-harnesstool-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

