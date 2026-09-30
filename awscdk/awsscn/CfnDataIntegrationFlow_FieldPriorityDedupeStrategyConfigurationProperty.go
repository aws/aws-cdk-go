package awsscn


// The field priority deduplication strategy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldPriorityDedupeStrategyConfigurationProperty := &FieldPriorityDedupeStrategyConfigurationProperty{
//   	Fields: []interface{}{
//   		&FieldPriorityDedupeFieldProperty{
//   			Name: jsii.String("name"),
//   			SortOrder: jsii.String("sortOrder"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-fieldprioritydedupestrategyconfiguration.html
//
type CfnDataIntegrationFlow_FieldPriorityDedupeStrategyConfigurationProperty struct {
	// The list of field names and their sort order for deduplication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-fieldprioritydedupestrategyconfiguration.html#cfn-scn-dataintegrationflow-fieldprioritydedupestrategyconfiguration-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
}

