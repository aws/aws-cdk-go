package awsbedrockagentcore


// Authentication method for calling a remote MCP server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var none interface{}
//
//   harnessRemoteMcpOutboundAuthProperty := &HarnessRemoteMcpOutboundAuthProperty{
//   	AwsIam: &IamCredentialProviderProperty{
//   		Service: jsii.String("service"),
//
//   		// the properties below are optional
//   		Region: jsii.String("region"),
//   	},
//   	None: none,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpoutboundauth.html
//
type CfnHarness_HarnessRemoteMcpOutboundAuthProperty struct {
	// IAM credentials used to SigV4-sign requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpoutboundauth.html#cfn-bedrockagentcore-harness-harnessremotemcpoutboundauth-awsiam
	//
	AwsIam interface{} `field:"optional" json:"awsIam" yaml:"awsIam"`
	// No authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessremotemcpoutboundauth.html#cfn-bedrockagentcore-harness-harnessremotemcpoutboundauth-none
	//
	None interface{} `field:"optional" json:"none" yaml:"none"`
}

