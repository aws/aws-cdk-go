package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wrapOptionsProperty := &WrapOptionsProperty{
//   	LineLength: jsii.Number(123),
//   	LineTerminator: jsii.String("lineTerminator"),
//   	WrapBy: jsii.String("wrapBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html
//
type CfnPartnership_WrapOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html#cfn-b2bi-partnership-wrapoptions-linelength
	//
	LineLength *float64 `field:"optional" json:"lineLength" yaml:"lineLength"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html#cfn-b2bi-partnership-wrapoptions-lineterminator
	//
	LineTerminator *string `field:"optional" json:"lineTerminator" yaml:"lineTerminator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html#cfn-b2bi-partnership-wrapoptions-wrapby
	//
	WrapBy *string `field:"optional" json:"wrapBy" yaml:"wrapBy"`
}

