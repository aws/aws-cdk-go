package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customInstructionsProperty := &CustomInstructionsProperty{
//   	CustomInstructionsString: jsii.String("customInstructionsString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-custominstructions.html
//
type CfnTopic_CustomInstructionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-custominstructions.html#cfn-quicksight-topic-custominstructions-custominstructionsstring
	//
	CustomInstructionsString *string `field:"required" json:"customInstructionsString" yaml:"customInstructionsString"`
}

