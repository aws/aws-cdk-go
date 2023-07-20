package awsecs


// The container path, mount options, and size of the tmpfs mount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfsProperty := &TmpfsProperty{
//   	Size: jsii.Number(123),
//
//   	// the properties below are optional
//   	ContainerPath: jsii.String("containerPath"),
//   	MountOptions: []*string{
//   		jsii.String("mountOptions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html
//
type CfnTaskDefinition_TmpfsProperty struct {
	// The maximum size (in MiB) of the tmpfs volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html#cfn-ecs-taskdefinition-tmpfs-size
	//
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The absolute file path where the tmpfs volume is to be mounted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html#cfn-ecs-taskdefinition-tmpfs-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The list of tmpfs volume mount options.
	//
	// Valid values: `"defaults" | "ro" | "rw" | "suid" | "nosuid" | "dev" | "nodev" | "exec" | "noexec" | "sync" | "async" | "dirsync" | "remount" | "mand" | "nomand" | "atime" | "noatime" | "diratime" | "nodiratime" | "bind" | "rbind" | "unbindable" | "runbindable" | "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave" | "relatime" | "norelatime" | "strictatime" | "nostrictatime" | "mode" | "uid" | "gid" | "nr_inodes" | "nr_blocks" | "mpol"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-tmpfs.html#cfn-ecs-taskdefinition-tmpfs-mountoptions
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

