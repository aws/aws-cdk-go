package awscomprehend


// Describes the entity list submitted with an entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   entityRecognizerEntityListProperty := &EntityRecognizerEntityListProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerentitylist.html
//
type CfnEntityRecognizerPropsMixin_EntityRecognizerEntityListProperty struct {
	// Specifies the Amazon S3 location where the entity list is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerentitylist.html#cfn-comprehend-entityrecognizer-entityrecognizerentitylist-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

