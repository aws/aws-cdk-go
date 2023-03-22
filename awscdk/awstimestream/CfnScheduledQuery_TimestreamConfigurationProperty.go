package awstimestream


// Configuration to write data into Timestream database and table.
//
// This configuration allows the user to map the query result select columns into the destination table columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestreamConfigurationProperty := &TimestreamConfigurationProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	DimensionMappings: []interface{}{
//   		&DimensionMappingProperty{
//   			DimensionValueType: jsii.String("dimensionValueType"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TableName: jsii.String("tableName"),
//   	TimeColumn: jsii.String("timeColumn"),
//
//   	// the properties below are optional
//   	MeasureNameColumn: jsii.String("measureNameColumn"),
//   	MixedMeasureMappings: []interface{}{
//   		&MixedMeasureMappingProperty{
//   			MeasureValueType: jsii.String("measureValueType"),
//
//   			// the properties below are optional
//   			MeasureName: jsii.String("measureName"),
//   			MultiMeasureAttributeMappings: []interface{}{
//   				&MultiMeasureAttributeMappingProperty{
//   					MeasureValueType: jsii.String("measureValueType"),
//   					SourceColumn: jsii.String("sourceColumn"),
//
//   					// the properties below are optional
//   					TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   				},
//   			},
//   			SourceColumn: jsii.String("sourceColumn"),
//   			TargetMeasureName: jsii.String("targetMeasureName"),
//   		},
//   	},
//   	MultiMeasureMappings: &MultiMeasureMappingsProperty{
//   		MultiMeasureAttributeMappings: []interface{}{
//   			&MultiMeasureAttributeMappingProperty{
//   				MeasureValueType: jsii.String("measureValueType"),
//   				SourceColumn: jsii.String("sourceColumn"),
//
//   				// the properties below are optional
//   				TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		TargetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   	},
//   }
//
type CfnScheduledQuery_TimestreamConfigurationProperty struct {
	// Name of Timestream database to which the query result will be written.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// This is to allow mapping column(s) from the query result to the dimension in the destination table.
	DimensionMappings interface{} `field:"required" json:"dimensionMappings" yaml:"dimensionMappings"`
	// Name of Timestream table that the query result will be written to.
	//
	// The table should be within the same database that is provided in Timestream configuration.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Column from query result that should be used as the time column in destination table.
	//
	// Column type for this should be TIMESTAMP.
	TimeColumn *string `field:"required" json:"timeColumn" yaml:"timeColumn"`
	// Name of the measure column.
	MeasureNameColumn *string `field:"optional" json:"measureNameColumn" yaml:"measureNameColumn"`
	// Specifies how to map measures to multi-measure records.
	MixedMeasureMappings interface{} `field:"optional" json:"mixedMeasureMappings" yaml:"mixedMeasureMappings"`
	// Multi-measure mappings.
	MultiMeasureMappings interface{} `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
}

