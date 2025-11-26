package previewawstimestreammixins


// Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided.
//
// MultiMeasureMappings can be used to ingest data as multi measures in the derived table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiMeasureMappingsProperty := &MultiMeasureMappingsProperty{
//   	MultiMeasureAttributeMappings: []interface{}{
//   		&MultiMeasureAttributeMappingProperty{
//   			MeasureValueType: jsii.String("measureValueType"),
//   			SourceColumn: jsii.String("sourceColumn"),
//   			TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   		},
//   	},
//   	TargetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasuremappings.html
//
type CfnScheduledQueryPropsMixin_MultiMeasureMappingsProperty struct {
	// Required.
	//
	// Attribute mappings to be used for mapping query results to ingest data for multi-measure attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasuremappings.html#cfn-timestream-scheduledquery-multimeasuremappings-multimeasureattributemappings
	//
	MultiMeasureAttributeMappings interface{} `field:"optional" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// The name of the target multi-measure name in the derived table.
	//
	// This input is required when measureNameColumn is not provided. If MeasureNameColumn is provided, then value from that column will be used as multi-measure name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasuremappings.html#cfn-timestream-scheduledquery-multimeasuremappings-targetmultimeasurename
	//
	TargetMultiMeasureName *string `field:"optional" json:"targetMultiMeasureName" yaml:"targetMultiMeasureName"`
}

