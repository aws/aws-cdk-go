package awsiot


// Properties for defining a `CfnLogging`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggingProps := &CfnLoggingProps{
//   	AccountId: jsii.String("accountId"),
//   	DefaultLogLevel: jsii.String("defaultLogLevel"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html
//
type CfnLoggingProps struct {
	// The account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The default log level.
	//
	// Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-defaultloglevel
	//
	DefaultLogLevel *string `field:"required" json:"defaultLogLevel" yaml:"defaultLogLevel"`
	// The role ARN used for the log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-rolearn
	//
	RoleArn interface{} `field:"required" json:"roleArn" yaml:"roleArn"`
}

