package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12SplitOptionsProperty := &X12SplitOptionsProperty{
//   	SplitBy: jsii.String("splitBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12splitoptions.html
//
type CfnTransformer_X12SplitOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12splitoptions.html#cfn-b2bi-transformer-x12splitoptions-splitby
	//
	SplitBy *string `field:"optional" json:"splitBy" yaml:"splitBy"`
}

