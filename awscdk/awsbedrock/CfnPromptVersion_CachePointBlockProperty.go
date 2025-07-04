package awsbedrock


// Defines a section of content to be cached for reuse in subsequent API calls.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-cachepointblock.html
//
type CfnPromptVersion_CachePointBlockProperty struct {
	// Specifies the type of cache point within the CachePointBlock.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-cachepointblock.html#cfn-bedrock-promptversion-cachepointblock-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

