package awsec2


// Properties for defining a `CfnVPCEndpointConnectionNotification`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEndpointConnectionNotificationProps := &CfnVPCEndpointConnectionNotificationProps{
//   	ConnectionEvents: []*string{
//   		jsii.String("connectionEvents"),
//   	},
//   	ConnectionNotificationArn: jsii.String("connectionNotificationArn"),
//
//   	// the properties below are optional
//   	ServiceId: jsii.String("serviceId"),
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html
//
type CfnVPCEndpointConnectionNotificationProps struct {
	// The endpoint events for which to receive notifications.
	//
	// Valid values are `Accept` , `Connect` , `Delete` , and `Reject` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-connectionevents
	//
	ConnectionEvents *[]*string `field:"required" json:"connectionEvents" yaml:"connectionEvents"`
	// The ARN of the SNS topic for the notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-connectionnotificationarn
	//
	ConnectionNotificationArn *string `field:"required" json:"connectionNotificationArn" yaml:"connectionNotificationArn"`
	// The ID of the endpoint service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-serviceid
	//
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// The ID of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

