package mixinsawsiotanalytics


// The value of the variable as a structure that specifies an output file URI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputFileUriValueProperty := &OutputFileUriValueProperty{
//   	FileName: jsii.String("fileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-outputfileurivalue.html
//
type CfnDatasetPropsMixin_OutputFileUriValueProperty struct {
	// The URI of the location where dataset contents are stored, usually the URI of a file in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-outputfileurivalue.html#cfn-iotanalytics-dataset-outputfileurivalue-filename
	//
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
}

