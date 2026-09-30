package awsscn


// The dataset options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   datasetOptionsProperty := &DatasetOptionsProperty{
//   	DedupeRecords: jsii.Boolean(false),
//   	DedupeStrategy: &DedupeStrategyProperty{
//   		FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   			Fields: []interface{}{
//   				&FieldPriorityDedupeFieldProperty{
//   					Name: jsii.String("name"),
//   					SortOrder: jsii.String("sortOrder"),
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	LoadType: jsii.String("loadType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetoptions.html
//
type CfnDataIntegrationFlowPropsMixin_DatasetOptionsProperty struct {
	// The option to perform deduplication on data records sharing same primary key values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetoptions.html#cfn-scn-dataintegrationflow-datasetoptions-deduperecords
	//
	DedupeRecords interface{} `field:"optional" json:"dedupeRecords" yaml:"dedupeRecords"`
	// The deduplication strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetoptions.html#cfn-scn-dataintegrationflow-datasetoptions-dedupestrategy
	//
	DedupeStrategy interface{} `field:"optional" json:"dedupeStrategy" yaml:"dedupeStrategy"`
	// The load type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetoptions.html#cfn-scn-dataintegrationflow-datasetoptions-loadtype
	//
	LoadType *string `field:"optional" json:"loadType" yaml:"loadType"`
}

