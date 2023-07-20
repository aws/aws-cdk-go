package awsquicksight


// A structure that represents a negative format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   negativeFormatProperty := &NegativeFormatProperty{
//   	Prefix: jsii.String("prefix"),
//   	Suffix: jsii.String("suffix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-negativeformat.html
//
type CfnTopic_NegativeFormatProperty struct {
	// The prefix for a negative format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-negativeformat.html#cfn-quicksight-topic-negativeformat-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The suffix for a negative format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-negativeformat.html#cfn-quicksight-topic-negativeformat-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

