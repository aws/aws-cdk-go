package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedOptionsProperty := &AdvancedOptionsProperty{
//   	X12: &X12AdvancedOptionsProperty{
//   		SplitOptions: &X12SplitOptionsProperty{
//   			SplitBy: jsii.String("splitBy"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-advancedoptions.html
//
type CfnTransformer_AdvancedOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-advancedoptions.html#cfn-b2bi-transformer-advancedoptions-x12
	//
	X12 interface{} `field:"optional" json:"x12" yaml:"x12"`
}

