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
type CfnResourceSpecificLoggingProps struct {
	// The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// The target name.
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// The target type.
	//
	// Valid Values: `DEFAULT | THING_GROUP`.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

