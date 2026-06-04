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
//   harnessToolConfigurationProperty := &HarnessToolConfigurationProperty{
//   	AgentCoreBrowser: &HarnessAgentCoreBrowserConfigProperty{
//   		BrowserArn: jsii.String("browserArn"),
//   	},
//   	AgentCoreCodeInterpreter: &HarnessAgentCoreCodeInterpreterConfigProperty{
//   		CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   	},
//   	AgentCoreGateway: &HarnessAgentCoreGatewayConfigProperty{
//   		GatewayArn: jsii.String("gatewayArn"),
//
//   		// the properties below are optional
//   		OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   			AwsIam: awsIam,
//   			None: none,
//   			Oauth: &OAuthCredentialProviderProperty{
//   				ProviderArn: jsii.String("providerArn"),
//   				Scopes: []*string{
//   					jsii.String("scopes"),
//   				},
//
//   				// the properties below are optional
//   				CustomParameters: map[string]*string{
//   					"customParametersKey": jsii.String("customParameters"),
//   				},
//   				DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   				GrantType: jsii.String("grantType"),
//   			},
//   		},
//   	},
//   	InlineFunction: &HarnessInlineFunctionConfigProperty{
//   		Description: jsii.String("description"),
//   		InputSchema: inputSchema,
//   	},
//   	RemoteMcp: &HarnessRemoteMcpConfigProperty{
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstoolconfiguration.html
//
type CfnHarness_HarnessToolConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstoolconfiguration.html#cfn-bedrockagentcore-harness-harnesstoolconfiguration-agentcorebrowser
	//
	AgentCoreBrowser interface{} `field:"optional" json:"agentCoreBrowser" yaml:"agentCoreBrowser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstoolconfiguration.html#cfn-bedrockagentcore-harness-harnesstoolconfiguration-agentcorecodeinterpreter
	//
	AgentCoreCodeInterpreter interface{} `field:"optional" json:"agentCoreCodeInterpreter" yaml:"agentCoreCodeInterpreter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstoolconfiguration.html#cfn-bedrockagentcore-harness-harnesstoolconfiguration-agentcoregateway
	//
	AgentCoreGateway interface{} `field:"optional" json:"agentCoreGateway" yaml:"agentCoreGateway"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstoolconfiguration.html#cfn-bedrockagentcore-harness-harnesstoolconfiguration-inlinefunction
	//
	InlineFunction interface{} `field:"optional" json:"inlineFunction" yaml:"inlineFunction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstoolconfiguration.html#cfn-bedrockagentcore-harness-harnesstoolconfiguration-remotemcp
	//
	RemoteMcp interface{} `field:"optional" json:"remoteMcp" yaml:"remoteMcp"`
}

