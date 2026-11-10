package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fallbackTriggersProperty := &FallbackTriggersProperty{
//   	Errors: []*string{
//   		jsii.String("errors"),
//   	},
//   	StatusCodes: []*string{
//   		jsii.String("statusCodes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktriggers.html
//
type CfnGatewayRule_FallbackTriggersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktriggers.html#cfn-bedrockagentcore-gatewayrule-fallbacktriggers-errors
	//
	Errors *[]*string `field:"optional" json:"errors" yaml:"errors"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallbacktriggers.html#cfn-bedrockagentcore-gatewayrule-fallbacktriggers-statuscodes
	//
	StatusCodes *[]*string `field:"optional" json:"statusCodes" yaml:"statusCodes"`
}

