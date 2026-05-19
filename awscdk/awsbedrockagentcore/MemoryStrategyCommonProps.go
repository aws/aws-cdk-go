package awsbedrockagentcore


// Configuration parameters common for any memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryStrategyCommonProps := &MemoryStrategyCommonProps{
//   	StrategyName: jsii.String("strategyName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type MemoryStrategyCommonProps struct {
	// The name for the strategy.
	StrategyName *string `field:"required" json:"strategyName" yaml:"strategyName"`
	// The description of the strategy.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

