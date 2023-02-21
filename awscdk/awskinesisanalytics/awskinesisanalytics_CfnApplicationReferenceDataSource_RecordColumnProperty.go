package awskinesisanalytics


// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
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
type CfnApplicationReferenceDataSource_RecordColumnProperty struct {
	// Name of the column created in the in-application input stream or reference table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of column created in the in-application input stream or reference table.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Reference to the data element in the streaming input or the reference data source.
	//
	// This element is required if the [RecordFormatType](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_RecordFormat.html#analytics-Type-RecordFormat-RecordFormatTypel) is `JSON` .
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

