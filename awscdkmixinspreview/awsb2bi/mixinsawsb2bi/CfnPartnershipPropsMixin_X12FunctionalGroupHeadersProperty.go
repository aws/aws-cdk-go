package mixinsawsb2bi


// Part of the X12 message structure.
//
// These are the functional group headers for the X12 EDI object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12FunctionalGroupHeadersProperty := &X12FunctionalGroupHeadersProperty{
//   	ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   	ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   	ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html
//
type CfnPartnershipPropsMixin_X12FunctionalGroupHeadersProperty struct {
	// A value representing the code used to identify the party receiving a message, at position GS-03.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html#cfn-b2bi-partnership-x12functionalgroupheaders-applicationreceivercode
	//
	ApplicationReceiverCode *string `field:"optional" json:"applicationReceiverCode" yaml:"applicationReceiverCode"`
	// A value representing the code used to identify the party transmitting a message, at position GS-02.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html#cfn-b2bi-partnership-x12functionalgroupheaders-applicationsendercode
	//
	ApplicationSenderCode *string `field:"optional" json:"applicationSenderCode" yaml:"applicationSenderCode"`
	// A code that identifies the issuer of the standard, at position GS-07.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html#cfn-b2bi-partnership-x12functionalgroupheaders-responsibleagencycode
	//
	ResponsibleAgencyCode *string `field:"optional" json:"responsibleAgencyCode" yaml:"responsibleAgencyCode"`
}

