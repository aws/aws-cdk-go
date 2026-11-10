package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDriverVolumeConfigurationProperty := &LogDriverVolumeConfigurationProperty{
//   	Driver: jsii.String("driver"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-logdrivervolumeconfiguration.html
//
type CfnDaemonTaskDefinition_LogDriverVolumeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-logdrivervolumeconfiguration.html#cfn-ecs-daemontaskdefinition-logdrivervolumeconfiguration-driver
	//
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

