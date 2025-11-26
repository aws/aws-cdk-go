package previewawsbedrockagentcoremixins


// The memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   semanticMemoryStrategyProperty := &SemanticMemoryStrategyProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html
//
type CfnMemoryPropsMixin_SemanticMemoryStrategyProperty struct {
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The memory strategy description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The memory strategy name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The memory strategy namespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Status of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The memory strategy ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// The memory strategy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Last update timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html#cfn-bedrockagentcore-memory-semanticmemorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

