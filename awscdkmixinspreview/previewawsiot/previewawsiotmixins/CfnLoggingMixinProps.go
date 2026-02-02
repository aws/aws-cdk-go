package previewawsiotmixins


// Properties for CfnLoggingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoggingMixinProps := &CfnLoggingMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	DefaultLogLevel: jsii.String("defaultLogLevel"),
//   	EventConfigurations: []interface{}{
//   		&EventConfigurationProperty{
//   			EventType: jsii.String("eventType"),
//   			LogDestination: jsii.String("logDestination"),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html
//
type CfnLoggingMixinProps struct {
	// The account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The default log level.
	//
	// Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-defaultloglevel
	//
	DefaultLogLevel *string `field:"optional" json:"defaultLogLevel" yaml:"defaultLogLevel"`
	// Configurations for event-based logging that specifies which event types to log and their logging settings.
	//
	// Overrides account-level logging for the specified event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-eventconfigurations
	//
	EventConfigurations interface{} `field:"optional" json:"eventConfigurations" yaml:"eventConfigurations"`
	// The role ARN used for the log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html#cfn-iot-logging-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

