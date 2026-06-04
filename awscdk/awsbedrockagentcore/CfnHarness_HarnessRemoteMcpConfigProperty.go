package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessRemoteMcpConfigProperty := &HarnessRemoteMcpConfigProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
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
}

