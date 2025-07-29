package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12AcknowledgmentOptionsProperty := &X12AcknowledgmentOptionsProperty{
//   	FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   	TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12acknowledgmentoptions.html
//
type CfnPartnership_X12AcknowledgmentOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12acknowledgmentoptions.html#cfn-b2bi-partnership-x12acknowledgmentoptions-functionalacknowledgment
	//
	FunctionalAcknowledgment *string `field:"required" json:"functionalAcknowledgment" yaml:"functionalAcknowledgment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12acknowledgmentoptions.html#cfn-b2bi-partnership-x12acknowledgmentoptions-technicalacknowledgment
	//
	TechnicalAcknowledgment *string `field:"required" json:"technicalAcknowledgment" yaml:"technicalAcknowledgment"`
}

