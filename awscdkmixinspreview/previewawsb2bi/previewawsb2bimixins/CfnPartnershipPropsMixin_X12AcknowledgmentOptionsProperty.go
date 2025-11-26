package previewawsb2bimixins


// Contains options for configuring X12 acknowledgments.
//
// These options control how functional and technical acknowledgments are handled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12AcknowledgmentOptionsProperty := &X12AcknowledgmentOptionsProperty{
//   	FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   	TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12acknowledgmentoptions.html
//
type CfnPartnershipPropsMixin_X12AcknowledgmentOptionsProperty struct {
	// Specifies whether functional acknowledgments (997/999) should be generated for incoming X12 transactions.
	//
	// Valid values are `DO_NOT_GENERATE` , `GENERATE_ALL_SEGMENTS` and `GENERATE_WITHOUT_TRANSACTION_SET_RESPONSE_LOOP` .
	//
	// If you choose `GENERATE_WITHOUT_TRANSACTION_SET_RESPONSE_LOOP` , AWS B2B Data Interchange skips the AK2_Loop when generating an acknowledgment document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12acknowledgmentoptions.html#cfn-b2bi-partnership-x12acknowledgmentoptions-functionalacknowledgment
	//
	FunctionalAcknowledgment *string `field:"optional" json:"functionalAcknowledgment" yaml:"functionalAcknowledgment"`
	// Specifies whether technical acknowledgments (TA1) should be generated for incoming X12 interchanges.
	//
	// Valid values are `DO_NOT_GENERATE` and `GENERATE_ALL_SEGMENTS` and.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12acknowledgmentoptions.html#cfn-b2bi-partnership-x12acknowledgmentoptions-technicalacknowledgment
	//
	TechnicalAcknowledgment *string `field:"optional" json:"technicalAcknowledgment" yaml:"technicalAcknowledgment"`
}

