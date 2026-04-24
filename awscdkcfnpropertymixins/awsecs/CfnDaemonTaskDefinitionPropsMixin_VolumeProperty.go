package awsecs


// The data volume configuration for tasks launched using this task definition.
//
// Specifying a volume configuration in a task definition is optional. The volume configuration may contain multiple volumes but only one volume configured at launch is supported. Each volume defined in the volume configuration may only specify a ``name`` and one of either ``configuredAtLaunch``, ``dockerVolumeConfiguration``, ``efsVolumeConfiguration``, ``fsxWindowsFileServerVolumeConfiguration``, or ``host``. If an empty volume configuration is specified, by default Amazon ECS uses a host volume. For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnDaemonTaskDefinitionPropsMixin_VolumeProperty struct {
	// Details on a container instance bind mount host volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-volume.html#cfn-ecs-daemontaskdefinition-volume-host
	//
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.
	//  When using a volume configured at launch, the ``name`` is required and must also be specified as the volume name in the ``ServiceVolumeConfiguration`` or ``TaskVolumeConfiguration`` parameter when creating your service or standalone task.
	//  For all other types of volumes, this name is referenced in the ``sourceVolume`` parameter of the ``mountPoints`` object in the container definition.
	//  When a volume is using the ``efsVolumeConfiguration``, the name is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-volume.html#cfn-ecs-daemontaskdefinition-volume-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

