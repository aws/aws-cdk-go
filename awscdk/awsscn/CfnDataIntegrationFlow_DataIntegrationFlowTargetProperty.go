package awsscn


// The DataIntegrationFlow target parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataIntegrationFlowTargetProperty := &DataIntegrationFlowTargetProperty{
//   	TargetType: jsii.String("targetType"),
//
//   	// the properties below are optional
//   	DatasetTarget: &DatasetTargetConfigurationProperty{
//   		DatasetIdentifier: jsii.String("datasetIdentifier"),
//
//   		// the properties below are optional
//   		Options: &DatasetOptionsProperty{
//   			DedupeRecords: jsii.Boolean(false),
//   			DedupeStrategy: &DedupeStrategyProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   					Fields: []interface{}{
//   						&FieldPriorityDedupeFieldProperty{
//   							Name: jsii.String("name"),
//   							SortOrder: jsii.String("sortOrder"),
//   						},
//   					},
//   				},
//   			},
//   			LoadType: jsii.String("loadType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtarget.html
//
type CfnDataIntegrationFlow_DataIntegrationFlowTargetProperty struct {
	// The target type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtarget.html#cfn-scn-dataintegrationflow-dataintegrationflowtarget-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// The dataset target configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtarget.html#cfn-scn-dataintegrationflow-dataintegrationflowtarget-datasettarget
	//
	DatasetTarget interface{} `field:"optional" json:"datasetTarget" yaml:"datasetTarget"`
}

