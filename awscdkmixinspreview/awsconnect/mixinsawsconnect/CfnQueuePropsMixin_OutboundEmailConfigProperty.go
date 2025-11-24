package mixinsawsconnect


// The outbound email address ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outboundEmailConfigProperty := &OutboundEmailConfigProperty{
//   	OutboundEmailAddressId: jsii.String("outboundEmailAddressId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundemailconfig.html
//
type CfnQueuePropsMixin_OutboundEmailConfigProperty struct {
	// The identifier of the email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-queue-outboundemailconfig.html#cfn-connect-queue-outboundemailconfig-outboundemailaddressid
	//
	OutboundEmailAddressId *string `field:"optional" json:"outboundEmailAddressId" yaml:"outboundEmailAddressId"`
}

