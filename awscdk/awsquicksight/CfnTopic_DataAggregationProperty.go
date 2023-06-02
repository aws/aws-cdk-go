package awsquicksight


// The definition of a data aggregation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAggregationProperty := &DataAggregationProperty{
//   	DatasetRowDateGranularity: jsii.String("datasetRowDateGranularity"),
//   	DefaultDateColumnName: jsii.String("defaultDateColumnName"),
//   }
//
type CfnTopic_DataAggregationProperty struct {
	// The level of time precision that is used to aggregate `DateTime` values.
	DatasetRowDateGranularity *string `field:"optional" json:"datasetRowDateGranularity" yaml:"datasetRowDateGranularity"`
	// The column name for the default date.
	DefaultDateColumnName *string `field:"optional" json:"defaultDateColumnName" yaml:"defaultDateColumnName"`
}

