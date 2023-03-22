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
//   clientConnectOptionsProperty := &clientConnectOptionsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   }
//
type CfnClientVpnEndpoint_ClientConnectOptionsProperty struct {
	// Indicates whether client connect options are enabled.
	//
	// The default is `false` (not enabled).
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the AWS Lambda function used for connection authorization.
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

