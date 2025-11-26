package previewawstimestreammixins


// Configuration used for writing the output of a query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   		MeasureNameColumn: jsii.String("measureNameColumn"),
//   		MixedMeasureMappings: []interface{}{
//   			&MixedMeasureMappingProperty{
//   				MeasureName: jsii.String("measureName"),
//   				MeasureValueType: jsii.String("measureValueType"),
//   				MultiMeasureAttributeMappings: []interface{}{
//   					&MultiMeasureAttributeMappingProperty{
//   						MeasureValueType: jsii.String("measureValueType"),
//   						SourceColumn: jsii.String("sourceColumn"),
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
//   					TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   				},
//   			},
//   			TargetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   		},
//   		TableName: jsii.String("tableName"),
//   		TimeColumn: jsii.String("timeColumn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-targetconfiguration.html
//
type CfnScheduledQueryPropsMixin_TargetConfigurationProperty struct {
	// Configuration needed to write data into the Timestream database and table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-targetconfiguration.html#cfn-timestream-scheduledquery-targetconfiguration-timestreamconfiguration
	//
	TimestreamConfiguration interface{} `field:"optional" json:"timestreamConfiguration" yaml:"timestreamConfiguration"`
}

