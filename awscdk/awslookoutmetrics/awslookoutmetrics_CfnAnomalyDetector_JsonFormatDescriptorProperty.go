package awslookoutmetrics


// Contains information about how a source JSON data file should be analyzed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFormatDescriptorProperty := &jsonFormatDescriptorProperty{
//   	charset: jsii.String("charset"),
//   	fileCompression: jsii.String("fileCompression"),
//   }
//
type CfnAnomalyDetector_JsonFormatDescriptorProperty struct {
	// The character set in which the source JSON file is written.
	Charset *string `field:"optional" json:"charset" yaml:"charset"`
	// The level of compression of the source CSV file.
	FileCompression *string `field:"optional" json:"fileCompression" yaml:"fileCompression"`
}

