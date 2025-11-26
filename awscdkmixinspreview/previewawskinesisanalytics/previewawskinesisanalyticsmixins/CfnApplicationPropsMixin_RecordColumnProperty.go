package previewawskinesisanalyticsmixins


// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordColumnProperty := &RecordColumnProperty{
//   	Mapping: jsii.String("mapping"),
//   	Name: jsii.String("name"),
//   	SqlType: jsii.String("sqlType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html
//
type CfnApplicationPropsMixin_RecordColumnProperty struct {
	// Reference to the data element in the streaming input or the reference data source.
	//
	// This element is required if the [RecordFormatType](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_RecordFormat.html#analytics-Type-RecordFormat-RecordFormatTypel) is `JSON` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html#cfn-kinesisanalytics-application-recordcolumn-mapping
	//
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
	// Name of the column created in the in-application input stream or reference table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html#cfn-kinesisanalytics-application-recordcolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type of column created in the in-application input stream or reference table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordcolumn.html#cfn-kinesisanalytics-application-recordcolumn-sqltype
	//
	SqlType *string `field:"optional" json:"sqlType" yaml:"sqlType"`
}

