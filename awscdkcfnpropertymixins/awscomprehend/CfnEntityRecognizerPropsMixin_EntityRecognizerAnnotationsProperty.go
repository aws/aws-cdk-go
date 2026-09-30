package awscomprehend


// Describes the annotations associated with an entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   entityRecognizerAnnotationsProperty := &EntityRecognizerAnnotationsProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   	TestS3Uri: jsii.String("testS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerannotations.html
//
type CfnEntityRecognizerPropsMixin_EntityRecognizerAnnotationsProperty struct {
	// Specifies the Amazon S3 location where the annotations are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerannotations.html#cfn-comprehend-entityrecognizer-entityrecognizerannotations-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// Specifies the Amazon S3 location where the test annotations are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerannotations.html#cfn-comprehend-entityrecognizer-entityrecognizerannotations-tests3uri
	//
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

