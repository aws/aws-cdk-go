package awsec2


// The status of the transit gateway peering attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   peeringAttachmentStatusProperty := &PeeringAttachmentStatusProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypeeringattachment-peeringattachmentstatus.html
//
type CfnTransitGatewayPeeringAttachment_PeeringAttachmentStatusProperty struct {
	// The status code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypeeringattachment-peeringattachmentstatus.html#cfn-ec2-transitgatewaypeeringattachment-peeringattachmentstatus-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The status message, if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaypeeringattachment-peeringattachmentstatus.html#cfn-ec2-transitgatewaypeeringattachment-peeringattachmentstatus-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

