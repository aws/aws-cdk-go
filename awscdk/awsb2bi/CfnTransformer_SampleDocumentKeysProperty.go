package awsb2bi


// An array of the Amazon S3 keys used to identify the location for your sample documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleDocumentKeysProperty := &SampleDocumentKeysProperty{
//   	Input: jsii.String("input"),
//   	Output: jsii.String("output"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocumentkeys.html
//
type CfnTransformer_SampleDocumentKeysProperty struct {
	// An array of keys for your input sample documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocumentkeys.html#cfn-b2bi-transformer-sampledocumentkeys-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// An array of keys for your output sample documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-sampledocumentkeys.html#cfn-b2bi-transformer-sampledocumentkeys-output
	//
	Output *string `field:"optional" json:"output" yaml:"output"`
}

