package awsec2


// A reference to a VPCEndpointConnectionNotification resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCEndpointConnectionNotificationReference := &VPCEndpointConnectionNotificationReference{
//   	VpcEndpointConnectionNotificationId: jsii.String("vpcEndpointConnectionNotificationId"),
//   }
//
type VPCEndpointConnectionNotificationReference struct {
	// The VPCEndpointConnectionNotificationId of the VPCEndpointConnectionNotification resource.
	VpcEndpointConnectionNotificationId *string `field:"required" json:"vpcEndpointConnectionNotificationId" yaml:"vpcEndpointConnectionNotificationId"`
}

