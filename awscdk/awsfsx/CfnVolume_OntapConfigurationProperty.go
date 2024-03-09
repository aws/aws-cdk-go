package awsfsx


// Specifies the configuration of the ONTAP volume that you are creating.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ontapConfigurationProperty := &OntapConfigurationProperty{
//   	StorageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   	// the properties below are optional
//   	AggregateConfiguration: &AggregateConfigurationProperty{
//   		Aggregates: []*string{
//   			jsii.String("aggregates"),
//   		},
//   		ConstituentsPerAggregate: jsii.Number(123),
//   	},
//   	CopyTagsToBackups: jsii.String("copyTagsToBackups"),
//   	JunctionPath: jsii.String("junctionPath"),
//   	OntapVolumeType: jsii.String("ontapVolumeType"),
//   	SecurityStyle: jsii.String("securityStyle"),
//   	SizeInBytes: jsii.String("sizeInBytes"),
//   	SizeInMegabytes: jsii.String("sizeInMegabytes"),
//   	SnaplockConfiguration: &SnaplockConfigurationProperty{
//   		SnaplockType: jsii.String("snaplockType"),
//
//   		// the properties below are optional
//   		AuditLogVolume: jsii.String("auditLogVolume"),
//   		AutocommitPeriod: &AutocommitPeriodProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Value: jsii.Number(123),
//   		},
//   		PrivilegedDelete: jsii.String("privilegedDelete"),
//   		RetentionPeriod: &SnaplockRetentionPeriodProperty{
//   			DefaultRetention: &RetentionPeriodProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Value: jsii.Number(123),
//   			},
//   			MaximumRetention: &RetentionPeriodProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Value: jsii.Number(123),
//   			},
//   			MinimumRetention: &RetentionPeriodProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		VolumeAppendModeEnabled: jsii.String("volumeAppendModeEnabled"),
//   	},
//   	SnapshotPolicy: jsii.String("snapshotPolicy"),
//   	StorageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   	TieringPolicy: &TieringPolicyProperty{
//   		CoolingPeriod: jsii.Number(123),
//   		Name: jsii.String("name"),
//   	},
//   	VolumeStyle: jsii.String("volumeStyle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html
//
type CfnVolume_OntapConfigurationProperty struct {
	// Specifies the ONTAP SVM in which to create the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-storagevirtualmachineid
	//
	StorageVirtualMachineId *string `field:"required" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// Used to specify the configuration options for an FSx for ONTAP volume's storage aggregate or aggregates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-aggregateconfiguration
	//
	AggregateConfiguration interface{} `field:"optional" json:"aggregateConfiguration" yaml:"aggregateConfiguration"`
	// A boolean flag indicating whether tags for the volume should be copied to backups.
	//
	// This value defaults to false. If it's set to true, all tags for the volume are copied to all automatic and user-initiated backups where the user doesn't specify tags. If this value is true, and you specify one or more tags, only the specified tags are copied to backups. If you specify one or more tags when creating a user-initiated backup, no tags are copied from the volume, regardless of this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-copytagstobackups
	//
	CopyTagsToBackups *string `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Specifies the location in the SVM's namespace where the volume is mounted.
	//
	// This parameter is required. The `JunctionPath` must have a leading forward slash, such as `/vol3` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-junctionpath
	//
	JunctionPath *string `field:"optional" json:"junctionPath" yaml:"junctionPath"`
	// Specifies the type of volume you are creating. Valid values are the following:.
	//
	// - `RW` specifies a read/write volume. `RW` is the default.
	// - `DP` specifies a data-protection volume. A `DP` volume is read-only and can be used as the destination of a NetApp SnapMirror relationship.
	//
	// For more information, see [Volume types](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-types) in the Amazon FSx for NetApp ONTAP User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-ontapvolumetype
	//
	OntapVolumeType *string `field:"optional" json:"ontapVolumeType" yaml:"ontapVolumeType"`
	// Specifies the security style for the volume.
	//
	// If a volume's security style is not specified, it is automatically set to the root volume's security style. The security style determines the type of permissions that FSx for ONTAP uses to control data access. For more information, see [Volume security style](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-security-style) in the *Amazon FSx for NetApp ONTAP User Guide* . Specify one of the following values:
	//
	// - `UNIX` if the file system is managed by a UNIX administrator, the majority of users are NFS clients, and an application accessing the data uses a UNIX user as the service account.
	// - `NTFS` if the file system is managed by a Windows administrator, the majority of users are SMB clients, and an application accessing the data uses a Windows user as the service account.
	// - `MIXED` This is an advanced setting. For more information, see the topic [What the security styles and their effects are](https://docs.aws.amazon.com/https://docs.netapp.com/us-en/ontap/nfs-admin/security-styles-their-effects-concept.html) in the NetApp Documentation Center.
	//
	// For more information, see [Volume security style](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-security-style.html) in the FSx for ONTAP User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-securitystyle
	//
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Specifies the configured size of the volume, in bytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-sizeinbytes
	//
	SizeInBytes *string `field:"optional" json:"sizeInBytes" yaml:"sizeInBytes"`
	// Use `SizeInBytes` instead.
	//
	// Specifies the size of the volume, in megabytes (MB), that you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-sizeinmegabytes
	//
	SizeInMegabytes *string `field:"optional" json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// The SnapLock configuration object for an FSx for ONTAP SnapLock volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-snaplockconfiguration
	//
	SnaplockConfiguration interface{} `field:"optional" json:"snaplockConfiguration" yaml:"snaplockConfiguration"`
	// Specifies the snapshot policy for the volume. There are three built-in snapshot policies:.
	//
	// - `default` : This is the default policy. A maximum of six hourly snapshots taken five minutes past the hour. A maximum of two daily snapshots taken Monday through Saturday at 10 minutes after midnight. A maximum of two weekly snapshots taken every Sunday at 15 minutes after midnight.
	// - `default-1weekly` : This policy is the same as the `default` policy except that it only retains one snapshot from the weekly schedule.
	// - `none` : This policy does not take any snapshots. This policy can be assigned to volumes to prevent automatic snapshots from being taken.
	//
	// You can also provide the name of a custom policy that you created with the ONTAP CLI or REST API.
	//
	// For more information, see [Snapshot policies](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snapshots-ontap.html#snapshot-policies) in the Amazon FSx for NetApp ONTAP User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-snapshotpolicy
	//
	SnapshotPolicy *string `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume, or set to false to disable them.
	//
	// `StorageEfficiencyEnabled` is required when creating a `RW` volume ( `OntapVolumeType` set to `RW` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-storageefficiencyenabled
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-tieringpolicy
	//
	TieringPolicy interface{} `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
	// Use to specify the style of an ONTAP volume.
	//
	// FSx for ONTAP offers two styles of volumes that you can use for different purposes, FlexVol and FlexGroup volumes. For more information, see [Volume styles](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-styles.html) in the Amazon FSx for NetApp ONTAP User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-ontapconfiguration.html#cfn-fsx-volume-ontapconfiguration-volumestyle
	//
	VolumeStyle *string `field:"optional" json:"volumeStyle" yaml:"volumeStyle"`
}

