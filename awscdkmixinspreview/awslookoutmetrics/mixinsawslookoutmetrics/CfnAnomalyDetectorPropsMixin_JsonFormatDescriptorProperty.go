package mixinsawslookoutmetrics


// Contains information about how a source JSON data file should be analyzed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jsonFormatDescriptorProperty := &JsonFormatDescriptorProperty{
//   	Charset: jsii.String("charset"),
//   	FileCompression: jsii.String("fileCompression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-jsonformatdescriptor.html
//
type CfnAnomalyDetectorPropsMixin_JsonFormatDescriptorProperty struct {
	// The character set in which the source JSON file is written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-jsonformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-jsonformatdescriptor-charset
	//
	Charset *string `field:"optional" json:"charset" yaml:"charset"`
	// The level of compression of the source CSV file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-jsonformatdescriptor.html#cfn-lookoutmetrics-anomalydetector-jsonformatdescriptor-filecompression
	//
	FileCompression *string `field:"optional" json:"fileCompression" yaml:"fileCompression"`
}

