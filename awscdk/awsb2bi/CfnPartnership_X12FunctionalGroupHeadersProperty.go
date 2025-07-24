package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12FunctionalGroupHeadersProperty := &X12FunctionalGroupHeadersProperty{
//   	ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   	ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   	ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html
//
type CfnPartnership_X12FunctionalGroupHeadersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html#cfn-b2bi-partnership-x12functionalgroupheaders-applicationreceivercode
	//
	ApplicationReceiverCode *string `field:"optional" json:"applicationReceiverCode" yaml:"applicationReceiverCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html#cfn-b2bi-partnership-x12functionalgroupheaders-applicationsendercode
	//
	ApplicationSenderCode *string `field:"optional" json:"applicationSenderCode" yaml:"applicationSenderCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12functionalgroupheaders.html#cfn-b2bi-partnership-x12functionalgroupheaders-responsibleagencycode
	//
	ResponsibleAgencyCode *string `field:"optional" json:"responsibleAgencyCode" yaml:"responsibleAgencyCode"`
}

