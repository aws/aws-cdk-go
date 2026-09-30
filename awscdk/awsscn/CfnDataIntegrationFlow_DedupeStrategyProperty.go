package awsscn


// The deduplication strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dedupeStrategyProperty := &DedupeStrategyProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   		Fields: []interface{}{
//   			&FieldPriorityDedupeFieldProperty{
//   				Name: jsii.String("name"),
//   				SortOrder: jsii.String("sortOrder"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dedupestrategy.html
//
type CfnDataIntegrationFlow_DedupeStrategyProperty struct {
	// The deduplication strategy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dedupestrategy.html#cfn-scn-dataintegrationflow-dedupestrategy-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The field priority deduplication strategy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dedupestrategy.html#cfn-scn-dataintegrationflow-dedupestrategy-fieldpriority
	//
	FieldPriority interface{} `field:"optional" json:"fieldPriority" yaml:"fieldPriority"`
}

