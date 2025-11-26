package previewawsbatchmixins


// The container path, mount options, and size of the `tmpfs` mount.
//
// > This object isn't applicable to jobs that are running on Fargate resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tmpfsProperty := &TmpfsProperty{
//   	ContainerPath: jsii.String("containerPath"),
//   	MountOptions: []*string{
//   		jsii.String("mountOptions"),
//   	},
//   	Size: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html
//
type CfnJobDefinitionPropsMixin_TmpfsProperty struct {
	// The absolute file path in the container where the `tmpfs` volume is mounted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html#cfn-batch-jobdefinition-tmpfs-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The list of `tmpfs` volume mount options.
	//
	// Valid values: " `defaults` " | " `ro` " | " `rw` " | " `suid` " | " `nosuid` " | " `dev` " | " `nodev` " | " `exec` " | " `noexec` " | " `sync` " | " `async` " | " `dirsync` " | " `remount` " | " `mand` " | " `nomand` " | " `atime` " | " `noatime` " | " `diratime` " | " `nodiratime` " | " `bind` " | " `rbind" | "unbindable" | "runbindable" | "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave" | "relatime` " | " `norelatime` " | " `strictatime` " | " `nostrictatime` " | " `mode` " | " `uid` " | " `gid` " | " `nr_inodes` " | " `nr_blocks` " | " `mpol` ".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html#cfn-batch-jobdefinition-tmpfs-mountoptions
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The size (in MiB) of the `tmpfs` volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-tmpfs.html#cfn-batch-jobdefinition-tmpfs-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

