package awsecs


// The Linux-specific options that are applied to the container, such as Linux [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   linuxParametersProperty := &LinuxParametersProperty{
//   	Capabilities: &KernelCapabilitiesProperty{
//   		Add: []*string{
//   			jsii.String("add"),
//   		},
//   		Drop: []*string{
//   			jsii.String("drop"),
//   		},
//   	},
//   	Devices: []interface{}{
//   		&DeviceProperty{
//   			ContainerPath: jsii.String("containerPath"),
//   			HostPath: jsii.String("hostPath"),
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   		},
//   	},
//   	InitProcessEnabled: jsii.Boolean(false),
//   	Tmpfs: []interface{}{
//   		&TmpfsProperty{
//   			ContainerPath: jsii.String("containerPath"),
//   			MountOptions: []*string{
//   				jsii.String("mountOptions"),
//   			},
//   			Size: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html
//
type CfnDaemonTaskDefinitionPropsMixin_LinuxParametersProperty struct {
	// The Linux capabilities to add or remove from the default Docker configuration for a container defined in the task definition.
	//
	// For more detailed information about these Linux capabilities, see the [capabilities(7)](https://docs.aws.amazon.com/http://man7.org/linux/man-pages/man7/capabilities.7.html) Linux manual page.
	//  The following describes how Docker processes the Linux capabilities specified in the ``add`` and ``drop`` request parameters. For information about the latest behavior, see [Docker Compose: order of cap_drop and cap_add](https://docs.aws.amazon.com/https://forums.docker.com/t/docker-compose-order-of-cap-drop-and-cap-add/97136/1) in the Docker Community Forum.
	//   +  When the container is a privleged container, the container capabilities are all of the default Docker capabilities. The capabilities specified in the ``add`` request parameter, and the ``drop`` request parameter are ignored.
	//   +  When the ``add`` request parameter is set to ALL, the container capabilities are all of the default Docker capabilities, excluding those specified in the ``drop`` request parameter.
	//   +  When the ``drop`` request parameter is set to ALL, the container capabilities are the capabilities specified in the ``add`` request parameter.
	//   +  When the ``add`` request parameter and the ``drop`` request parameter are both empty, the capabilities the container capabilities are all of the default Docker capabilities.
	//   +  The default is to first drop the capabilities specified in the ``drop`` request parameter, and then add the capabilities specified in the ``add`` request parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-capabilities
	//
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Any host devices to expose to the container.
	//
	// This parameter maps to ``Devices`` in the docker container create command and the ``--device`` option to docker run.
	//   If you're using tasks that use the Fargate launch type, the ``devices`` parameter isn't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-devices
	//
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Run an ``init`` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the ``--init`` option to docker run. This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ``sudo docker version --format '{{.Server.APIVersion}}'``
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-initprocessenabled
	//
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the ``--tmpfs`` option to docker run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-tmpfs
	//
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

