package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVerifiedDestinationNumberPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnVerifiedDestinationNumberMixinProps := &CfnVerifiedDestinationNumberMixinProps{
//   	DestinationPhoneNumber: jsii.String("destinationPhoneNumber"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-verifieddestinationnumber.html
//
type CfnVerifiedDestinationNumberMixinProps struct {
	// The verified destination phone number, in E.164 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-verifieddestinationnumber.html#cfn-smsvoice-verifieddestinationnumber-destinationphonenumber
	//
	DestinationPhoneNumber *string `field:"optional" json:"destinationPhoneNumber" yaml:"destinationPhoneNumber"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-verifieddestinationnumber.html#cfn-smsvoice-verifieddestinationnumber-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

