package awsscn


// The dataset source configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetSourceConfigurationProperty := &DatasetSourceConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetsourceconfiguration.html
//
type CfnDataIntegrationFlow_DatasetSourceConfigurationProperty struct {
	// The ARN of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetsourceconfiguration.html#cfn-scn-dataintegrationflow-datasetsourceconfiguration-datasetidentifier
	//
	DatasetIdentifier *string `field:"required" json:"datasetIdentifier" yaml:"datasetIdentifier"`
	// The dataset options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-datasetsourceconfiguration.html#cfn-scn-dataintegrationflow-datasetsourceconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

