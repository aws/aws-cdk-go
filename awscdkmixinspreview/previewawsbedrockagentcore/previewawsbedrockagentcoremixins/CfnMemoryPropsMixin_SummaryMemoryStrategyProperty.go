package previewawsbedrockagentcoremixins


// The memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   summaryMemoryStrategyProperty := &SummaryMemoryStrategyProperty{
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
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
type CfnMemoryPropsMixin_SummaryMemoryStrategyProperty struct {
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The memory strategy description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The memory strategy name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The summary memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// The memory strategy status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The memory strategy ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// The memory strategy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The memory strategy update date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html#cfn-bedrockagentcore-memory-summarymemorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

