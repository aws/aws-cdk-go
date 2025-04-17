package awstimestream


// MixedMeasureMappings are mappings that can be used to ingest data into a mixture of narrow and multi measures in the derived table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mixedMeasureMappingProperty := &MixedMeasureMappingProperty{
//   	MeasureValueType: jsii.String("measureValueType"),
//
//   	// the properties below are optional
//   	MeasureName: jsii.String("measureName"),
//   	MultiMeasureAttributeMappings: []interface{}{
//   		&MultiMeasureAttributeMappingProperty{
//   			MeasureValueType: jsii.String("measureValueType"),
//   			SourceColumn: jsii.String("sourceColumn"),
//
//   			// the properties below are optional
//   			TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   		},
//   	},
//   	SourceColumn: jsii.String("sourceColumn"),
//   	TargetMeasureName: jsii.String("targetMeasureName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-mixedmeasuremapping.html
//
type CfnScheduledQuery_MixedMeasureMappingProperty struct {
	// Type of the value that is to be read from sourceColumn.
	//
	// If the mapping is for MULTI, use MeasureValueType.MULTI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-mixedmeasuremapping.html#cfn-timestream-scheduledquery-mixedmeasuremapping-measurevaluetype
	//
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Refers to the value of measure_name in a result row.
	//
	// This field is required if MeasureNameColumn is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-mixedmeasuremapping.html#cfn-timestream-scheduledquery-mixedmeasuremapping-measurename
	//
	MeasureName *string `field:"optional" json:"measureName" yaml:"measureName"`
	// Required when measureValueType is MULTI.
	//
	// Attribute mappings for MULTI value measures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-mixedmeasuremapping.html#cfn-timestream-scheduledquery-mixedmeasuremapping-multimeasureattributemappings
	//
	MultiMeasureAttributeMappings interface{} `field:"optional" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// This field refers to the source column from which measure-value is to be read for result materialization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-mixedmeasuremapping.html#cfn-timestream-scheduledquery-mixedmeasuremapping-sourcecolumn
	//
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Target measure name to be used.
	//
	// If not provided, the target measure name by default would be measure-name if provided, or sourceColumn otherwise.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-mixedmeasuremapping.html#cfn-timestream-scheduledquery-mixedmeasuremapping-targetmeasurename
	//
	TargetMeasureName *string `field:"optional" json:"targetMeasureName" yaml:"targetMeasureName"`
}

