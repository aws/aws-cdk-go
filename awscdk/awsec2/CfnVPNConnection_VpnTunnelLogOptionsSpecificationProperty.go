package awsec2


// Options for logging VPN tunnel activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpnTunnelLogOptionsSpecificationProperty := &VpnTunnelLogOptionsSpecificationProperty{
//   	CloudwatchLogOptions: &CloudwatchLogOptionsSpecificationProperty{
//   		LogEnabled: jsii.Boolean(false),
//   		LogGroupArn: jsii.String("logGroupArn"),
//   		LogOutputFormat: jsii.String("logOutputFormat"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunnellogoptionsspecification.html
//
type CfnVPNConnection_VpnTunnelLogOptionsSpecificationProperty struct {
	// Options for sending VPN tunnel logs to CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunnellogoptionsspecification.html#cfn-ec2-vpnconnection-vpntunnellogoptionsspecification-cloudwatchlogoptions
	//
	CloudwatchLogOptions interface{} `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

