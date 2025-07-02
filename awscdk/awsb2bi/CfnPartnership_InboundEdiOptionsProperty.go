package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inboundEdiOptionsProperty := &InboundEdiOptionsProperty{
//   	X12: &X12InboundEdiOptionsProperty{
//   		AcknowledgmentOptions: &X12AcknowledgmentOptionsProperty{
//   			FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   			TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-inboundedioptions.html
//
type CfnPartnership_InboundEdiOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-inboundedioptions.html#cfn-b2bi-partnership-inboundedioptions-x12
	//
	X12 interface{} `field:"optional" json:"x12" yaml:"x12"`
}

