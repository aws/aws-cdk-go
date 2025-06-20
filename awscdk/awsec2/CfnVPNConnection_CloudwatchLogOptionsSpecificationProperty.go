package awsec2


// Options for sending VPN tunnel logs to CloudWatch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchLogOptionsSpecificationProperty := &CloudwatchLogOptionsSpecificationProperty{
//   	LogEnabled: jsii.Boolean(false),
//   	LogGroupArn: jsii.String("logGroupArn"),
//   	LogOutputFormat: jsii.String("logOutputFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-cloudwatchlogoptionsspecification.html
//
type CfnVPNConnection_CloudwatchLogOptionsSpecificationProperty struct {
	// Enable or disable VPN tunnel logging feature. Default value is `False` .
	//
	// Valid values: `True` | `False`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-cloudwatchlogoptionsspecification.html#cfn-ec2-vpnconnection-cloudwatchlogoptionsspecification-logenabled
	//
	LogEnabled interface{} `field:"optional" json:"logEnabled" yaml:"logEnabled"`
	// The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-cloudwatchlogoptionsspecification.html#cfn-ec2-vpnconnection-cloudwatchlogoptionsspecification-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// Set log format. Default format is `json` .
	//
	// Valid values: `json` | `text`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-cloudwatchlogoptionsspecification.html#cfn-ec2-vpnconnection-cloudwatchlogoptionsspecification-logoutputformat
	//
	LogOutputFormat *string `field:"optional" json:"logOutputFormat" yaml:"logOutputFormat"`
}

