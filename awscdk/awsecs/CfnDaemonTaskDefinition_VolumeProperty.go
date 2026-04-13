package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeProperty := &VolumeProperty{
//   	Host: &HostVolumePropertiesProperty{
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-volume.html
//
type CfnDaemonTaskDefinition_VolumeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-volume.html#cfn-ecs-daemontaskdefinition-volume-host
	//
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-volume.html#cfn-ecs-daemontaskdefinition-volume-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

