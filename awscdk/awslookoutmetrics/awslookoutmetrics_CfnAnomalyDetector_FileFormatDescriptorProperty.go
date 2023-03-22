package awslookoutmetrics


// Contains information about a source file's formatting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileFormatDescriptorProperty := &fileFormatDescriptorProperty{
//   	csvFormatDescriptor: &csvFormatDescriptorProperty{
//   		charset: jsii.String("charset"),
//   		containsHeader: jsii.Boolean(false),
//   		delimiter: jsii.String("delimiter"),
//   		fileCompression: jsii.String("fileCompression"),
//   		headerList: []*string{
//   			jsii.String("headerList"),
//   		},
//   		quoteSymbol: jsii.String("quoteSymbol"),
//   	},
//   	jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   		charset: jsii.String("charset"),
//   		fileCompression: jsii.String("fileCompression"),
//   	},
//   }
//
type CfnAnomalyDetector_FileFormatDescriptorProperty struct {
	// Contains information about how a source CSV data file should be analyzed.
	CsvFormatDescriptor interface{} `field:"optional" json:"csvFormatDescriptor" yaml:"csvFormatDescriptor"`
	// Contains information about how a source JSON data file should be analyzed.
	JsonFormatDescriptor interface{} `field:"optional" json:"jsonFormatDescriptor" yaml:"jsonFormatDescriptor"`
}

