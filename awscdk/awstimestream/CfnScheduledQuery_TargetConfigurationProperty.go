package awstimestream


// Configuration used for writing the output of a query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetConfigurationProperty := &TargetConfigurationProperty{
//   	TimestreamConfiguration: &TimestreamConfigurationProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		DimensionMappings: []interface{}{
//   			&DimensionMappingProperty{
//   				DimensionValueType: jsii.String("dimensionValueType"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TableName: jsii.String("tableName"),
//   		TimeColumn: jsii.String("timeColumn"),
//
//   		// the properties below are optional
//   		MeasureNameColumn: jsii.String("measureNameColumn"),
//   		MixedMeasureMappings: []interface{}{
//   			&MixedMeasureMappingProperty{
//   				MeasureValueType: jsii.String("measureValueType"),
//
//   				// the properties below are optional
//   				MeasureName: jsii.String("measureName"),
//   				MultiMeasureAttributeMappings: []interface{}{
//   					&MultiMeasureAttributeMappingProperty{
//   						MeasureValueType: jsii.String("measureValueType"),
//   						SourceColumn: jsii.String("sourceColumn"),
//
//   						// the properties below are optional
//   						TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   					},
//   				},
//   				SourceColumn: jsii.String("sourceColumn"),
//   				TargetMeasureName: jsii.String("targetMeasureName"),
//   			},
//   		},
//   		MultiMeasureMappings: &MultiMeasureMappingsProperty{
//   			MultiMeasureAttributeMappings: []interface{}{
//   				&MultiMeasureAttributeMappingProperty{
//   					MeasureValueType: jsii.String("measureValueType"),
//   					SourceColumn: jsii.String("sourceColumn"),
//
//   					// the properties below are optional
//   					TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			TargetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-targetconfiguration.html
//
type CfnScheduledQuery_TargetConfigurationProperty struct {
	// Configuration needed to write data into the Timestream database and table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-targetconfiguration.html#cfn-timestream-scheduledquery-targetconfiguration-timestreamconfiguration
	//
	TimestreamConfiguration interface{} `field:"required" json:"timestreamConfiguration" yaml:"timestreamConfiguration"`
}

