package awskinesisanalytics


// Describes the data format when records are written to the destination in a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationSchemaProperty := &destinationSchemaProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//   }
//
type CfnApplicationOutputV2_DestinationSchemaProperty struct {
	// Specifies the format of the records on the output stream.
	RecordFormatType *string `field:"optional" json:"recordFormatType" yaml:"recordFormatType"`
}

