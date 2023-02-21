package awslookoutmetrics


// Contains information about how a source CSV data file should be analyzed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvFormatDescriptorProperty := &csvFormatDescriptorProperty{
//   	charset: jsii.String("charset"),
//   	containsHeader: jsii.Boolean(false),
//   	delimiter: jsii.String("delimiter"),
//   	fileCompression: jsii.String("fileCompression"),
//   	headerList: []*string{
//   		jsii.String("headerList"),
//   	},
//   	quoteSymbol: jsii.String("quoteSymbol"),
//   }
//
type CfnAnomalyDetector_CsvFormatDescriptorProperty struct {
	// The character set in which the source CSV file is written.
	Charset *string `field:"optional" json:"charset" yaml:"charset"`
	// Whether or not the source CSV file contains a header.
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// The character used to delimit the source CSV file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The level of compression of the source CSV file.
	FileCompression *string `field:"optional" json:"fileCompression" yaml:"fileCompression"`
	// A list of the source CSV file's headers, if any.
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
	// The character used as a quote character.
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
}

