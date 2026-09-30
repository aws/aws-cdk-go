package awsscn


// The DataIntegrationFlow target parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataIntegrationFlowTargetProperty := &DataIntegrationFlowTargetProperty{
//   	DatasetTarget: &DatasetTargetConfigurationProperty{
//   		DatasetIdentifier: jsii.String("datasetIdentifier"),
//   		Options: &DatasetOptionsProperty{
//   			DedupeRecords: jsii.Boolean(false),
//   			DedupeStrategy: &DedupeStrategyProperty{
//   				FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   					Fields: []interface{}{
//   						&FieldPriorityDedupeFieldProperty{
//   							Name: jsii.String("name"),
//   							SortOrder: jsii.String("sortOrder"),
//   						},
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			LoadType: jsii.String("loadType"),
//   		},
//   	},
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtarget.html
//
type CfnDataIntegrationFlowPropsMixin_DataIntegrationFlowTargetProperty struct {
	// The dataset target configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtarget.html#cfn-scn-dataintegrationflow-dataintegrationflowtarget-datasettarget
	//
	DatasetTarget interface{} `field:"optional" json:"datasetTarget" yaml:"datasetTarget"`
	// The target type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtarget.html#cfn-scn-dataintegrationflow-dataintegrationflowtarget-targettype
	//
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
}

