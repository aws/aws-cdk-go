package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPartnership`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPartnershipProps := &CfnPartnershipProps{
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Email: jsii.String("email"),
//   	Name: jsii.String("name"),
//   	ProfileId: jsii.String("profileId"),
//
//   	// the properties below are optional
//   	CapabilityOptions: &CapabilityOptionsProperty{
//   		OutboundEdi: &OutboundEdiOptionsProperty{
//   			X12: &X12EnvelopeProperty{
//   				Common: &X12OutboundEdiHeadersProperty{
//   					Delimiters: &X12DelimitersProperty{
//   						ComponentSeparator: jsii.String("componentSeparator"),
//   						DataElementSeparator: jsii.String("dataElementSeparator"),
//   						SegmentTerminator: jsii.String("segmentTerminator"),
//   					},
//   					FunctionalGroupHeaders: &X12FunctionalGroupHeadersProperty{
//   						ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   						ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   						ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   					},
//   					InterchangeControlHeaders: &X12InterchangeControlHeadersProperty{
//   						AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   						ReceiverId: jsii.String("receiverId"),
//   						ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   						RepetitionSeparator: jsii.String("repetitionSeparator"),
//   						SenderId: jsii.String("senderId"),
//   						SenderIdQualifier: jsii.String("senderIdQualifier"),
//   						UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   					},
//   					ValidateEdi: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   	},
//   	Phone: jsii.String("phone"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html
//
type CfnPartnershipProps struct {
	// Returns one or more capabilities associated with this partnership.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-capabilities
	//
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-email
	//
	Email *string `field:"required" json:"email" yaml:"email"`
	// Returns the name of the partnership.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Returns the unique, system-generated identifier for the profile connected to this partnership.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Contains the details for an Outbound EDI capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-capabilityoptions
	//
	CapabilityOptions interface{} `field:"optional" json:"capabilityOptions" yaml:"capabilityOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-phone
	//
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
	// A key-value pair for a specific partnership.
	//
	// Tags are metadata that you can use to search for and group capabilities for various purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

