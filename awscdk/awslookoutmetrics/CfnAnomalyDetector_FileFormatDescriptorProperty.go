package awslookoutmetrics


// Contains information about a source file's formatting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileFormatDescriptorProperty := &FileFormatDescriptorProperty{
//   	CsvFormatDescriptor: &CsvFormatDescriptorProperty{
//   		Charset: jsii.String("charset"),
//   		ContainsHeader: jsii.Boolean(false),
//   		Delimiter: jsii.String("delimiter"),
//   		FileCompression: jsii.String("fileCompression"),
//   		HeaderList: []*string{
//   			jsii.String("headerList"),
//   		},
//   		QuoteSymbol: jsii.String("quoteSymbol"),
//   	},
//   	JsonFormatDescriptor: &JsonFormatDescriptorProperty{
//   		Charset: jsii.String("charset"),
//   		FileCompression: jsii.String("fileCompression"),
//   	},
//   }
//
type CfnAnomalyDetector_FileFormatDescriptorProperty struct {
	// Contains information about how a source CSV data file should be analyzed.
	CsvFormatDescriptor interface{} `field:"optional" json:"csvFormatDescriptor" yaml:"csvFormatDescriptor"`
	// Contains information about how a source JSON data file should be analyzed.
	JsonFormatDescriptor interface{} `field:"optional" json:"jsonFormatDescriptor" yaml:"jsonFormatDescriptor"`
}

