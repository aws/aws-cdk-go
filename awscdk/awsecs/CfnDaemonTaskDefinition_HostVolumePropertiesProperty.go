package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostVolumePropertiesProperty := &HostVolumePropertiesProperty{
//   	SourcePath: jsii.String("sourcePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-hostvolumeproperties.html
//
type CfnDaemonTaskDefinition_HostVolumePropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-hostvolumeproperties.html#cfn-ecs-daemontaskdefinition-hostvolumeproperties-sourcepath
	//
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

