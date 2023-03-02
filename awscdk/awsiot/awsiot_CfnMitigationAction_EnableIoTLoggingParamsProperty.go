package awsiot


// Parameters used when defining a mitigation action that enable AWS IoT Core logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enableIoTLoggingParamsProperty := &enableIoTLoggingParamsProperty{
//   	logLevel: jsii.String("logLevel"),
//   	roleArnForLogging: jsii.String("roleArnForLogging"),
//   }
//
type CfnMitigationAction_EnableIoTLoggingParamsProperty struct {
	// Specifies the type of information to be logged.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// The Amazon Resource Name (ARN) of the IAM role used for logging.
	RoleArnForLogging *string `field:"required" json:"roleArnForLogging" yaml:"roleArnForLogging"`
}

