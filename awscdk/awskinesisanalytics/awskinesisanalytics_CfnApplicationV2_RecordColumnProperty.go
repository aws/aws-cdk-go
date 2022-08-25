package awskinesisanalytics


// For a SQL-based Kinesis Data Analytics application, describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordColumnProperty := &recordColumnProperty{
//   	name: jsii.String("name"),
//   	sqlType: jsii.String("sqlType"),
//
//   	// the properties below are optional
//   	mapping: jsii.String("mapping"),
//   }
//
type CfnApplicationV2_RecordColumnProperty struct {
	// The name of the column that is created in the in-application input stream or reference table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of column created in the in-application input stream or reference table.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// A reference to the data element in the streaming input or the reference data source.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

