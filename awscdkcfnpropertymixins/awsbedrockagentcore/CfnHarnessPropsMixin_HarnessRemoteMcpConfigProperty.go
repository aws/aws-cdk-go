package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var none interface{}
//
//   harnessRemoteMcpConfigProperty := &HarnessRemoteMcpConfigProperty{
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	OutboundAuth: &HarnessRemoteMcpOutboundAuthProperty{
//   		AwsIam: &IamCredentialProviderProperty{
//   			Region: jsii.String("region"),
//   			Service: jsii.String("service"),
//   		},
//   		None: none,
//   	},
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html
//
type CfnHarnessPropsMixin_HarnessRemoteMcpConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html#cfn-bedrockagentcore-harness-harnessremotemcpconfig-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Authentication method for calling a remote MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html#cfn-bedrockagentcore-harness-harnessremotemcpconfig-outboundauth
	//
	OutboundAuth interface{} `field:"optional" json:"outboundAuth" yaml:"outboundAuth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html#cfn-bedrockagentcore-harness-harnessremotemcpconfig-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

