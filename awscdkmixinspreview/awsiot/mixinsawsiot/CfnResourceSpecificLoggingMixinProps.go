package mixinsawsiot


// Properties for CfnResourceSpecificLoggingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceSpecificLoggingMixinProps := &CfnResourceSpecificLoggingMixinProps{
//   	LogLevel: jsii.String("logLevel"),
//   	TargetName: jsii.String("targetName"),
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html
//
type CfnResourceSpecificLoggingMixinProps struct {
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html#cfn-iot-resourcespecificlogging-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The target name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html#cfn-iot-resourcespecificlogging-targetname
	//
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
	// The target type.
	//
	// Valid Values: `DEFAULT | THING_GROUP`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-resourcespecificlogging.html#cfn-iot-resourcespecificlogging-targettype
	//
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
}

