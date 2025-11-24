package mixinsawsbedrockagentcore


// The token based trigger input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tokenBasedTriggerInputProperty := &TokenBasedTriggerInputProperty{
//   	TokenCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-tokenbasedtriggerinput.html
//
type CfnMemoryPropsMixin_TokenBasedTriggerInputProperty struct {
	// The token based trigger token count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-tokenbasedtriggerinput.html#cfn-bedrockagentcore-memory-tokenbasedtriggerinput-tokencount
	//
	TokenCount *float64 `field:"optional" json:"tokenCount" yaml:"tokenCount"`
}

