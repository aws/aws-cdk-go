package awsbedrock


// Indicates where a cache checkpoint is located.
//
// All information before this checkpoint is cached to be accessed on subsequent requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachePointBlockProperty := &CachePointBlockProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-cachepointblock.html
//
type CfnPrompt_CachePointBlockProperty struct {
	// Indicates that the CachePointBlock is of the default type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-cachepointblock.html#cfn-bedrock-prompt-cachepointblock-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

