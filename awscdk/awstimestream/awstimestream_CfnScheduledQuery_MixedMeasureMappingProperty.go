package awstimestream


// MixedMeasureMappings are mappings that can be used to ingest data into a mixture of narrow and multi measures in the derived table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mixedMeasureMappingProperty := &mixedMeasureMappingProperty{
//   	measureValueType: jsii.String("measureValueType"),
//
//   	// the properties below are optional
//   	measureName: jsii.String("measureName"),
//   	multiMeasureAttributeMappings: []interface{}{
//   		&multiMeasureAttributeMappingProperty{
//   			measureValueType: jsii.String("measureValueType"),
//   			sourceColumn: jsii.String("sourceColumn"),
//
//   			// the properties below are optional
//   			targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   		},
//   	},
//   	sourceColumn: jsii.String("sourceColumn"),
//   	targetMeasureName: jsii.String("targetMeasureName"),
//   }
//
type CfnScheduledQuery_MixedMeasureMappingProperty struct {
	// Type of the value that is to be read from sourceColumn.
	//
	// If the mapping is for MULTI, use MeasureValueType.MULTI.
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Refers to the value of measure_name in a result row.
	//
	// This field is required if MeasureNameColumn is provided.
	MeasureName *string `field:"optional" json:"measureName" yaml:"measureName"`
	// Required when measureValueType is MULTI.
	//
	// Attribute mappings for MULTI value measures.
	MultiMeasureAttributeMappings interface{} `field:"optional" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// This field refers to the source column from which measure-value is to be read for result materialization.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Target measure name to be used.
	//
	// If not provided, the target measure name by default would be measure-name if provided, or sourceColumn otherwise.
	TargetMeasureName *string `field:"optional" json:"targetMeasureName" yaml:"targetMeasureName"`
}

