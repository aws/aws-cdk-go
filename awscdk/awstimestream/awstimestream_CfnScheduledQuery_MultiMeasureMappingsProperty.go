package awstimestream


// Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided.
//
// MultiMeasureMappings can be used to ingest data as multi measures in the derived table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiMeasureMappingsProperty := &multiMeasureMappingsProperty{
//   	multiMeasureAttributeMappings: []interface{}{
//   		&multiMeasureAttributeMappingProperty{
//   			measureValueType: jsii.String("measureValueType"),
//   			sourceColumn: jsii.String("sourceColumn"),
//
//   			// the properties below are optional
//   			targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	targetMultiMeasureName: jsii.String("targetMultiMeasureName"),
//   }
//
type CfnScheduledQuery_MultiMeasureMappingsProperty struct {
	// Required.
	//
	// Attribute mappings to be used for mapping query results to ingest data for multi-measure attributes.
	MultiMeasureAttributeMappings interface{} `field:"required" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// The name of the target multi-measure name in the derived table.
	//
	// This input is required when measureNameColumn is not provided. If MeasureNameColumn is provided, then value from that column will be used as multi-measure name.
	TargetMultiMeasureName *string `field:"optional" json:"targetMultiMeasureName" yaml:"targetMultiMeasureName"`
}

