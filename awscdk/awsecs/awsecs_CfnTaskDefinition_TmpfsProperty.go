package awsecs


// The `Tmpfs` property specifies the container path, mount options, and size of the tmpfs mount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfsProperty := &tmpfsProperty{
//   	size: jsii.Number(123),
//
//   	// the properties below are optional
//   	containerPath: jsii.String("containerPath"),
//   	mountOptions: []*string{
//   		jsii.String("mountOptions"),
//   	},
//   }
//
type CfnTaskDefinition_TmpfsProperty struct {
	// The maximum size (in MiB) of the tmpfs volume.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The list of tmpfs volume mount options.
	//
	// Valid values: `"defaults" | "ro" | "rw" | "suid" | "nosuid" | "dev" | "nodev" | "exec" | "noexec" | "sync" | "async" | "dirsync" | "remount" | "mand" | "nomand" | "atime" | "noatime" | "diratime" | "nodiratime" | "bind" | "rbind" | "unbindable" | "runbindable" | "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave" | "relatime" | "norelatime" | "strictatime" | "nostrictatime" | "mode" | "uid" | "gid" | "nr_inodes" | "nr_blocks" | "mpol"`.
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

