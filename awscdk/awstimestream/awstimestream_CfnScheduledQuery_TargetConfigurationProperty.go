package awstimestream


// Configuration used for writing the output of a query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetConfigurationProperty := &targetConfigurationProperty{
//   	timestreamConfiguration: &timestreamConfigurationProperty{
//   		databaseName: jsii.String("databaseName"),
//   		dimensionMappings: []interface{}{
//   			&dimensionMappingProperty{
//   				dimensionValueType: jsii.String("dimensionValueType"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		tableName: jsii.String("tableName"),
//   		timeColumn: jsii.String("timeColumn"),
//
//   		// the properties below are optional
//   		measureNameColumn: jsii.String("measureNameColumn"),
//   		mixedMeasureMappings: []interface{}{
//   			&mixedMeasureMappingProperty{
//   				measureValueType: jsii.String("measureValueType"),
//
//   				// the properties below are optional
//   				measureName: jsii.String("measureName"),
//   				multiMeasureAttributeMappings: []interface{}{
//   					&multiMeasureAttributeMappingProperty{
//   						measureValueType: jsii.String("measureValueType"),
//   						sourceColumn: jsii.String("sourceColumn"),
//
//   						// the properties below are optional
//   						targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   					},
//   				},
//   				sourceColumn: jsii.String("sourceColumn"),
//   				targetMeasureName: jsii.String("targetMeasureName"),
//   			},
//   		},
//   		multiMeasureMappings: &multiMeasureMappingsProperty{
//   			multiMeasureAttributeMappings: []interface{}{
//   				&multiMeasureAttributeMappingProperty{
//   					measureValueType: jsii.String("measureValueType"),
//   					sourceColumn: jsii.String("sourceColumn"),
//
//   					// the properties below are optional
//   					targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			targetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   		},
//   	},
//   }
//
type CfnScheduledQuery_TargetConfigurationProperty struct {
	// Configuration needed to write data into the Timestream database and table.
	TimestreamConfiguration interface{} `field:"required" json:"timestreamConfiguration" yaml:"timestreamConfiguration"`
}

