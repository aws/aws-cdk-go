package awsfsx


// Specifies the configuration of the ONTAP volume that you are creating.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ontapConfigurationProperty := &ontapConfigurationProperty{
//   	sizeInMegabytes: jsii.String("sizeInMegabytes"),
//   	storageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   	// the properties below are optional
//   	copyTagsToBackups: jsii.String("copyTagsToBackups"),
//   	junctionPath: jsii.String("junctionPath"),
//   	ontapVolumeType: jsii.String("ontapVolumeType"),
//   	securityStyle: jsii.String("securityStyle"),
//   	snapshotPolicy: jsii.String("snapshotPolicy"),
//   	storageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   	tieringPolicy: &tieringPolicyProperty{
//   		coolingPeriod: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnVolume_OntapConfigurationProperty struct {
	// Specifies the size of the volume, in megabytes (MB), that you are creating.
	SizeInMegabytes *string `field:"required" json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// Specifies the ONTAP SVM in which to create the volume.
	StorageVirtualMachineId *string `field:"required" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// `CfnVolume.OntapConfigurationProperty.CopyTagsToBackups`.
	CopyTagsToBackups *string `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Specifies the location in the SVM's namespace where the volume is mounted.
	//
	// The `JunctionPath` must have a leading forward slash, such as `/vol3` .
	JunctionPath *string `field:"optional" json:"junctionPath" yaml:"junctionPath"`
	// `CfnVolume.OntapConfigurationProperty.OntapVolumeType`.
	OntapVolumeType *string `field:"optional" json:"ontapVolumeType" yaml:"ontapVolumeType"`
	// The security style for the volume. Specify one of the following values:.
	//
	// - `UNIX` if the file system is managed by a UNIX administrator, the majority of users are NFS clients, and an application accessing the data uses a UNIX user as the service account. `UNIX` is the default.
	// - `NTFS` if the file system is managed by a Windows administrator, the majority of users are SMB clients, and an application accessing the data uses a Windows user as the service account.
	// - `MIXED` if the file system is managed by both UNIX and Windows administrators and users consist of both NFS and SMB clients.
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// `CfnVolume.OntapConfigurationProperty.SnapshotPolicy`.
	SnapshotPolicy *string `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume.
	StorageEfficiencyEnabled *string `field:"optional" json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// Describes the data tiering policy for an ONTAP volume.
	//
	// When enabled, Amazon FSx for ONTAP's intelligent tiering automatically transitions a volume's data between the file system's primary storage and capacity pool storage based on your access patterns.
	//
	// Valid tiering policies are the following:
	//
	// - `SNAPSHOT_ONLY` - (Default value) moves cold snapshots to the capacity pool storage tier.
	//
	// - `AUTO` - moves cold user data and snapshots to the capacity pool storage tier based on your access patterns.
	//
	// - `ALL` - moves all user data blocks in both the active file system and Snapshot copies to the storage pool tier.
	//
	// - `NONE` - keeps a volume's data in the primary storage tier, preventing it from being moved to the capacity pool tier.
	TieringPolicy interface{} `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
}

