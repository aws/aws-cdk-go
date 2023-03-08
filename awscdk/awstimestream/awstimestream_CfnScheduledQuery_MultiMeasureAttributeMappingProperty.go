package awstimestream


// Attribute mapping for MULTI value measures.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiMeasureAttributeMappingProperty := &multiMeasureAttributeMappingProperty{
//   	measureValueType: jsii.String("measureValueType"),
//   	sourceColumn: jsii.String("sourceColumn"),
//
//   	// the properties below are optional
//   	targetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   }
//
type CfnScheduledQuery_MultiMeasureAttributeMappingProperty struct {
	// Type of the attribute to be read from the source column.
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Source column from where the attribute value is to be read.
	SourceColumn *string `field:"required" json:"sourceColumn" yaml:"sourceColumn"`
	// Custom name to be used for attribute name in derived table.
	//
	// If not provided, source column name would be used.
	TargetMultiMeasureAttributeName *string `field:"optional" json:"targetMultiMeasureAttributeName" yaml:"targetMultiMeasureAttributeName"`
}

