package awskinesisfirehose


// Options for DecompressionProcessor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var decompressionProcessorCompressionFormat DecompressionProcessorCompressionFormat
//
//   decompressionProcessorOptions := &DecompressionProcessorOptions{
//   	CompressionFormat: decompressionProcessorCompressionFormat,
//   }
//
type DecompressionProcessorOptions struct {
	// The input compression format.
	// Default: DecompressionProcessorCompressionFormat.GZIP
	//
	CompressionFormat DecompressionProcessorCompressionFormat `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
}

