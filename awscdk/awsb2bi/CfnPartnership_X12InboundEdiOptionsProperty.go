package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12InboundEdiOptionsProperty := &X12InboundEdiOptionsProperty{
//   	AcknowledgmentOptions: &X12AcknowledgmentOptionsProperty{
//   		FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   		TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12inboundedioptions.html
//
type CfnPartnership_X12InboundEdiOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12inboundedioptions.html#cfn-b2bi-partnership-x12inboundedioptions-acknowledgmentoptions
	//
	AcknowledgmentOptions interface{} `field:"optional" json:"acknowledgmentOptions" yaml:"acknowledgmentOptions"`
}

