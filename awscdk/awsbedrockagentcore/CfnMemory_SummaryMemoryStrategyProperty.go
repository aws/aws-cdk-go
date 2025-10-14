package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   summaryMemoryStrategyProperty := &SummaryMemoryStrategyProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   	Status: jsii.String("status"),
//   	StrategyId: jsii.String("strategyId"),
//   	Type: jsii.String("type"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html
//
type CfnMemory_SummaryMemoryStrategyProperty struct {
	// Name of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Description of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Status of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Unique identifier for the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// Type of memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Last update timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

