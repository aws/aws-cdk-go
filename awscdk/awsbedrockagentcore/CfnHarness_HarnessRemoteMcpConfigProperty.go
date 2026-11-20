package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var none interface{}
//
//   harnessRemoteMcpConfigProperty := &HarnessRemoteMcpConfigProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	OutboundAuth: &HarnessRemoteMcpOutboundAuthProperty{
//   		AwsIam: &IamCredentialProviderProperty{
//   			Service: jsii.String("service"),
//
//   			// the properties below are optional
//   			Region: jsii.String("region"),
//   		},
//   		None: none,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html
//
type CfnHarness_HarnessRemoteMcpConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html#cfn-bedrockagentcore-harness-harnessremotemcpconfig-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html#cfn-bedrockagentcore-harness-harnessremotemcpconfig-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Authentication method for calling a remote MCP server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpconfig.html#cfn-bedrockagentcore-harness-harnessremotemcpconfig-outboundauth
	//
	OutboundAuth interface{} `field:"optional" json:"outboundAuth" yaml:"outboundAuth"`
}

