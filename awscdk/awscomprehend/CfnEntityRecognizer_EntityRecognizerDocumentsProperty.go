package awscomprehend


// Describes the training documents submitted with an entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityRecognizerDocumentsProperty := &EntityRecognizerDocumentsProperty{
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	InputFormat: jsii.String("inputFormat"),
//   	TestS3Uri: jsii.String("testS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerdocuments.html
//
type CfnEntityRecognizer_EntityRecognizerDocumentsProperty struct {
	// Specifies the Amazon S3 location where the training documents are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerdocuments.html#cfn-comprehend-entityrecognizer-entityrecognizerdocuments-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Specifies how the text in an input file should be processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerdocuments.html#cfn-comprehend-entityrecognizer-entityrecognizerdocuments-inputformat
	//
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Specifies the Amazon S3 location where the test documents are located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerdocuments.html#cfn-comprehend-entityrecognizer-entityrecognizerdocuments-tests3uri
	//
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}

