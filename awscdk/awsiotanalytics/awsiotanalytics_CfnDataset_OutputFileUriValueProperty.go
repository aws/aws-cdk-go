package awsiotanalytics


// The value of the variable as a structure that specifies an output file URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFileUriValueProperty := &outputFileUriValueProperty{
//   	fileName: jsii.String("fileName"),
//   }
//
type CfnDataset_OutputFileUriValueProperty struct {
	// The URI of the location where dataset contents are stored, usually the URI of a file in an S3 bucket.
	FileName *string `field:"required" json:"fileName" yaml:"fileName"`
}

