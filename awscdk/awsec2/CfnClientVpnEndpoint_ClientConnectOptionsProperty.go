package awsec2


// Indicates whether client connect options are enabled.
//
// The default is `false` (not enabled).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientConnectOptionsProperty := &ClientConnectOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientconnectoptions.html
//
type CfnClientVpnEndpoint_ClientConnectOptionsProperty struct {
	// Indicates whether client connect options are enabled.
	//
	// The default is `false` (not enabled).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientconnectoptions.html#cfn-ec2-clientvpnendpoint-clientconnectoptions-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the AWS Lambda function used for connection authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientconnectoptions.html#cfn-ec2-clientvpnendpoint-clientconnectoptions-lambdafunctionarn
	//
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

