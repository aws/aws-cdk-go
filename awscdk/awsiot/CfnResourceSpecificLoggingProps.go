package awsiot


// Properties for defining a `CfnResourceSpecificLogging`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceSpecificLoggingProps := &CfnResourceSpecificLoggingProps{
//   	LogLevel: jsii.String("logLevel"),
//   	TargetName: jsii.String("targetName"),
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html
//
type CfnResourceSpecificLoggingProps struct {
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html#cfn-iot-resourcespecificlogging-loglevel
	//
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// The target name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html#cfn-iot-resourcespecificlogging-targetname
	//
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// The target type.
	//
	// Valid Values: `DEFAULT | THING_GROUP`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html#cfn-iot-resourcespecificlogging-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

