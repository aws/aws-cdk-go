package awslookoutmetrics


// Contains information about the column used to track time in a source data file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestampColumnProperty := &timestampColumnProperty{
//   	columnFormat: jsii.String("columnFormat"),
//   	columnName: jsii.String("columnName"),
//   }
//
type CfnAnomalyDetector_TimestampColumnProperty struct {
	// The format of the timestamp column.
	ColumnFormat *string `field:"optional" json:"columnFormat" yaml:"columnFormat"`
	// The name of the timestamp column.
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
}

