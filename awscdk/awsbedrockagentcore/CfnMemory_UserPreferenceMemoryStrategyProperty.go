package awsbedrockagentcore


// The memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPreferenceMemoryStrategyProperty := &UserPreferenceMemoryStrategyProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html
//
type CfnMemory_UserPreferenceMemoryStrategyProperty struct {
	// The memory strategy name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The memory strategy description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The memory namespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// The memory strategy status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The memory strategy ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// The memory strategy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The memory strategy update date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html#cfn-bedrockagentcore-memory-userpreferencememorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

