package awsscn


// The dataset target configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetTargetConfigurationProperty := &DatasetTargetConfigurationProperty{
//   	DatasetIdentifier: jsii.String("datasetIdentifier"),
//
//   	// the properties below are optional
//   	Options: &DatasetOptionsProperty{
//   		DedupeRecords: jsii.Boolean(false),
//   		DedupeStrategy: &DedupeStrategyProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   				Fields: []interface{}{
//   					&FieldPriorityDedupeFieldProperty{
//   						Name: jsii.String("name"),
//   						SortOrder: jsii.String("sortOrder"),
//   					},
//   				},
//   			},
//   		},
//   		LoadType: jsii.String("loadType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasettargetconfiguration.html
//
type CfnDataIntegrationFlow_DatasetTargetConfigurationProperty struct {
	// The dataset ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasettargetconfiguration.html#cfn-scn-dataintegrationflow-datasettargetconfiguration-datasetidentifier
	//
	DatasetIdentifier *string `field:"required" json:"datasetIdentifier" yaml:"datasetIdentifier"`
	// The dataset options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasettargetconfiguration.html#cfn-scn-dataintegrationflow-datasettargetconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

