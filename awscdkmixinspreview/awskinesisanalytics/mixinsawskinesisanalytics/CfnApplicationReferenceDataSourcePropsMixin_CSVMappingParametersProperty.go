package mixinsawskinesisanalytics


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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cSVMappingParametersProperty := &CSVMappingParametersProperty{
//   	RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   	RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-csvmappingparameters.html
//
type CfnApplicationReferenceDataSourcePropsMixin_CSVMappingParametersProperty struct {
	// Column delimiter.
	//
	// For example, in a CSV format, a comma (",") is the typical column delimiter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-csvmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-csvmappingparameters-recordcolumndelimiter
	//
	RecordColumnDelimiter *string `field:"optional" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Row delimiter.
	//
	// For example, in a CSV format, *'\n'* is the typical row delimiter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-csvmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-csvmappingparameters-recordrowdelimiter
	//
	RecordRowDelimiter *string `field:"optional" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

