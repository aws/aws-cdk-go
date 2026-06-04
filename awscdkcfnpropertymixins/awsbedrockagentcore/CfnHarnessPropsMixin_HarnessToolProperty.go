package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var awsIam interface{}
//   var inputSchema interface{}
//   var none interface{}
//
//   harnessToolProperty := &HarnessToolProperty{
//   	Config: &HarnessToolConfigurationProperty{
//   		AgentCoreBrowser: &HarnessAgentCoreBrowserConfigProperty{
//   			BrowserArn: jsii.String("browserArn"),
//   		},
//   		AgentCoreCodeInterpreter: &HarnessAgentCoreCodeInterpreterConfigProperty{
//   			CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   		},
//   		AgentCoreGateway: &HarnessAgentCoreGatewayConfigProperty{
//   			GatewayArn: jsii.String("gatewayArn"),
//   			OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   				AwsIam: awsIam,
//   				None: none,
//   				Oauth: &OAuthCredentialProviderProperty{
//   					CustomParameters: map[string]*string{
//   						"customParametersKey": jsii.String("customParameters"),
//   					},
//   					DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   					GrantType: jsii.String("grantType"),
//   					ProviderArn: jsii.String("providerArn"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   		},
//   		InlineFunction: &HarnessInlineFunctionConfigProperty{
//   			Description: jsii.String("description"),
//   			InputSchema: inputSchema,
//   		},
//   		RemoteMcp: &HarnessRemoteMcpConfigProperty{
//   			Headers: map[string]*string{
//   				"headersKey": jsii.String("headers"),
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html
//
type CfnHarnessPropsMixin_HarnessToolProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html#cfn-bedrockagentcore-harness-harnesstool-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html#cfn-bedrockagentcore-harness-harnesstool-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstool.html#cfn-bedrockagentcore-harness-harnesstool-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

