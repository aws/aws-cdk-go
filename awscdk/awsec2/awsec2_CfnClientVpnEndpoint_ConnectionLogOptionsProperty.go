package awsec2


// Describes the client connection logging options for the Client VPN endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionLogOptionsProperty := &connectionLogOptionsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	cloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   	cloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   }
//
type CfnClientVpnEndpoint_ConnectionLogOptionsProperty struct {
	// Indicates whether connection logging is enabled.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the CloudWatch Logs log group.
	//
	// Required if connection logging is enabled.
	CloudwatchLogGroup *string `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// The name of the CloudWatch Logs log stream to which the connection data is published.
	CloudwatchLogStream *string `field:"optional" json:"cloudwatchLogStream" yaml:"cloudwatchLogStream"`
}

