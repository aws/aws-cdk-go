package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12AdvancedOptionsProperty := &X12AdvancedOptionsProperty{
//   	SplitOptions: &X12SplitOptionsProperty{
//   		SplitBy: jsii.String("splitBy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12advancedoptions.html
//
type CfnTransformer_X12AdvancedOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12advancedoptions.html#cfn-b2bi-transformer-x12advancedoptions-splitoptions
	//
	SplitOptions interface{} `field:"optional" json:"splitOptions" yaml:"splitOptions"`
}

