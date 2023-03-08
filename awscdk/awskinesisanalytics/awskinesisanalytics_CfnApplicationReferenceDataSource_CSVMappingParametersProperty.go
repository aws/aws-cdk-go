package awskinesisanalytics


// Provides additional mapping information when the record format uses delimiters, such as CSV.
//
// For example, the following sample records use CSV format, where the records use the *'\n'* as the row delimiter and a comma (",") as the column delimiter:
//
// `"name1", "address1"`
//
// `"name2", "address2"`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cSVMappingParametersProperty := &cSVMappingParametersProperty{
//   	recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   	recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   }
//
type CfnApplicationReferenceDataSource_CSVMappingParametersProperty struct {
	// Column delimiter.
	//
	// For example, in a CSV format, a comma (",") is the typical column delimiter.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Row delimiter.
	//
	// For example, in a CSV format, *'\n'* is the typical row delimiter.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

