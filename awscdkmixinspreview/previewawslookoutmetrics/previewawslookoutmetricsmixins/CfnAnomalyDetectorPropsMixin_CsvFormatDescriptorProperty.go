package previewawslookoutmetricsmixins


// Contains information about how a source CSV data file should be analyzed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   csvFormatDescriptorProperty := &CsvFormatDescriptorProperty{
//   	Charset: jsii.String("charset"),
//   	ContainsHeader: jsii.Boolean(false),
//   	Delimiter: jsii.String("delimiter"),
//   	FileCompression: jsii.String("fileCompression"),
//   	HeaderList: []*string{
//   		jsii.String("headerList"),
//   	},
//   	QuoteSymbol: jsii.String("quoteSymbol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html
//
type CfnAnomalyDetectorPropsMixin_CsvFormatDescriptorProperty struct {
	// The character set in which the source CSV file is written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-csvformatdescriptor-charset
	//
	Charset *string `field:"optional" json:"charset" yaml:"charset"`
	// Whether or not the source CSV file contains a header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-csvformatdescriptor-containsheader
	//
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// The character used to delimit the source CSV file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-csvformatdescriptor-delimiter
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The level of compression of the source CSV file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-csvformatdescriptor-filecompression
	//
	FileCompression *string `field:"optional" json:"fileCompression" yaml:"fileCompression"`
	// A list of the source CSV file's headers, if any.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-csvformatdescriptor-headerlist
	//
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
	// The character used as a quote character.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-csvformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-csvformatdescriptor-quotesymbol
	//
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
}

