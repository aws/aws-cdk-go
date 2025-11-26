package previewawsecsmixins


// The Linux capabilities to add or remove from the default Docker configuration for a container defined in the task definition.
//
// For more detailed information about these Linux capabilities, see the [capabilities(7)](https://docs.aws.amazon.com/http://man7.org/linux/man-pages/man7/capabilities.7.html) Linux manual page.
//
// The following describes how Docker processes the Linux capabilities specified in the `add` and `drop` request parameters. For information about the latest behavior, see [Docker Compose: order of cap_drop and cap_add](https://docs.aws.amazon.com/https://forums.docker.com/t/docker-compose-order-of-cap-drop-and-cap-add/97136/1) in the Docker Community Forum.
//
// - When the container is a privleged container, the container capabilities are all of the default Docker capabilities. The capabilities specified in the `add` request parameter, and the `drop` request parameter are ignored.
// - When the `add` request parameter is set to ALL, the container capabilities are all of the default Docker capabilities, excluding those specified in the `drop` request parameter.
// - When the `drop` request parameter is set to ALL, the container capabilities are the capabilities specified in the `add` request parameter.
// - When the `add` request parameter and the `drop` request parameter are both empty, the capabilities the container capabilities are all of the default Docker capabilities.
// - The default is to first drop the capabilities specified in the `drop` request parameter, and then add the capabilities specified in the `add` request parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kernelCapabilitiesProperty := &KernelCapabilitiesProperty{
//   	Add: []*string{
//   		jsii.String("add"),
//   	},
//   	Drop: []*string{
//   		jsii.String("drop"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-kernelcapabilities.html
//
type CfnTaskDefinitionPropsMixin_KernelCapabilitiesProperty struct {
	// The Linux capabilities for the container that have been added to the default configuration provided by Docker.
	//
	// This parameter maps to `CapAdd` in the docker container create command and the `--cap-add` option to docker run.
	//
	// > Tasks launched on AWS Fargate only support adding the `SYS_PTRACE` kernel capability.
	//
	// Valid values: `"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-kernelcapabilities.html#cfn-ecs-taskdefinition-kernelcapabilities-add
	//
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// The Linux capabilities for the container that have been removed from the default configuration provided by Docker.
	//
	// This parameter maps to `CapDrop` in the docker container create command and the `--cap-drop` option to docker run.
	//
	// Valid values: `"ALL" | "AUDIT_CONTROL" | "AUDIT_WRITE" | "BLOCK_SUSPEND" | "CHOWN" | "DAC_OVERRIDE" | "DAC_READ_SEARCH" | "FOWNER" | "FSETID" | "IPC_LOCK" | "IPC_OWNER" | "KILL" | "LEASE" | "LINUX_IMMUTABLE" | "MAC_ADMIN" | "MAC_OVERRIDE" | "MKNOD" | "NET_ADMIN" | "NET_BIND_SERVICE" | "NET_BROADCAST" | "NET_RAW" | "SETFCAP" | "SETGID" | "SETPCAP" | "SETUID" | "SYS_ADMIN" | "SYS_BOOT" | "SYS_CHROOT" | "SYS_MODULE" | "SYS_NICE" | "SYS_PACCT" | "SYS_PTRACE" | "SYS_RAWIO" | "SYS_RESOURCE" | "SYS_TIME" | "SYS_TTY_CONFIG" | "SYSLOG" | "WAKE_ALARM"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-kernelcapabilities.html#cfn-ecs-taskdefinition-kernelcapabilities-drop
	//
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

