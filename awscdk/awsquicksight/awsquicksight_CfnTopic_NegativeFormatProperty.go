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
type CfnTopic_NegativeFormatProperty struct {
	// The prefix for a negative format.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The suffix for a negative format.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

