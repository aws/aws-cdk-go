package awsiot


// Parameters used when defining a mitigation action that enable AWS IoT Core logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enableIoTLoggingParamsProperty := &EnableIoTLoggingParamsProperty{
//   	LogLevel: jsii.String("logLevel"),
//   	RoleArnForLogging: jsii.String("roleArnForLogging"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-enableiotloggingparams.html
//
type CfnMitigationAction_EnableIoTLoggingParamsProperty struct {
	// Specifies the type of information to be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-enableiotloggingparams.html#cfn-iot-mitigationaction-enableiotloggingparams-loglevel
	//
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// The Amazon Resource Name (ARN) of the IAM role used for logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-enableiotloggingparams.html#cfn-iot-mitigationaction-enableiotloggingparams-rolearnforlogging
	//
	RoleArnForLogging *string `field:"required" json:"roleArnForLogging" yaml:"roleArnForLogging"`
}

