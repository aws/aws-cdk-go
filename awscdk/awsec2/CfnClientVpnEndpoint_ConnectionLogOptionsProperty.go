package awsec2


// Describes the client connection logging options for the Client VPN endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionLogOptionsProperty := &ConnectionLogOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	CloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   	CloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html
//
type CfnClientVpnEndpoint_ConnectionLogOptionsProperty struct {
	// Indicates whether connection logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html#cfn-ec2-clientvpnendpoint-connectionlogoptions-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the CloudWatch Logs log group.
	//
	// Required if connection logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html#cfn-ec2-clientvpnendpoint-connectionlogoptions-cloudwatchloggroup
	//
	CloudwatchLogGroup *string `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// The name of the CloudWatch Logs log stream to which the connection data is published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html#cfn-ec2-clientvpnendpoint-connectionlogoptions-cloudwatchlogstream
	//
	CloudwatchLogStream *string `field:"optional" json:"cloudwatchLogStream" yaml:"cloudwatchLogStream"`
}

