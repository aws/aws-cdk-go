package awsfsx


// Specifies the SnapLock configuration for an FSx for ONTAP SnapLock volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snaplockConfigurationProperty := &SnaplockConfigurationProperty{
//   	SnaplockType: jsii.String("snaplockType"),
//
//   	// the properties below are optional
//   	AuditLogVolume: jsii.String("auditLogVolume"),
//   	AutocommitPeriod: &AutocommitPeriodProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Value: jsii.Number(123),
//   	},
//   	PrivilegedDelete: jsii.String("privilegedDelete"),
//   	RetentionPeriod: &SnaplockRetentionPeriodProperty{
//   		DefaultRetention: &RetentionPeriodProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Value: jsii.Number(123),
//   		},
//   		MaximumRetention: &RetentionPeriodProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Value: jsii.Number(123),
//   		},
//   		MinimumRetention: &RetentionPeriodProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	VolumeAppendModeEnabled: jsii.String("volumeAppendModeEnabled"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html
//
type CfnVolume_SnaplockConfigurationProperty struct {
	// Specifies the retention mode of an FSx for ONTAP SnapLock volume.
	//
	// After it is set, it can't be changed. You can choose one of the following retention modes:
	//
	// - `COMPLIANCE` : Files transitioned to write once, read many (WORM) on a Compliance volume can't be deleted until their retention periods expire. This retention mode is used to address government or industry-specific mandates or to protect against ransomware attacks. For more information, see [SnapLock Compliance](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snaplock-compliance.html) .
	// - `ENTERPRISE` : Files transitioned to WORM on an Enterprise volume can be deleted by authorized users before their retention periods expire using privileged delete. This retention mode is used to advance an organization's data integrity and internal compliance or to test retention settings before using SnapLock Compliance. For more information, see [SnapLock Enterprise](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snaplock-enterprise.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html#cfn-fsx-volume-snaplockconfiguration-snaplocktype
	//
	SnaplockType *string `field:"required" json:"snaplockType" yaml:"snaplockType"`
	// Enables or disables the audit log volume for an FSx for ONTAP SnapLock volume.
	//
	// The default value is `false` . If you set `AuditLogVolume` to `true` , the SnapLock volume is created as an audit log volume. The minimum retention period for an audit log volume is six months.
	//
	// For more information, see [SnapLock audit log volumes](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/how-snaplock-works.html#snaplock-audit-log-volume) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html#cfn-fsx-volume-snaplockconfiguration-auditlogvolume
	//
	AuditLogVolume *string `field:"optional" json:"auditLogVolume" yaml:"auditLogVolume"`
	// The configuration object for setting the autocommit period of files in an FSx for ONTAP SnapLock volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html#cfn-fsx-volume-snaplockconfiguration-autocommitperiod
	//
	AutocommitPeriod interface{} `field:"optional" json:"autocommitPeriod" yaml:"autocommitPeriod"`
	// Enables, disables, or permanently disables privileged delete on an FSx for ONTAP SnapLock Enterprise volume.
	//
	// Enabling privileged delete allows SnapLock administrators to delete write once, read many (WORM) files even if they have active retention periods. `PERMANENTLY_DISABLED` is a terminal state. If privileged delete is permanently disabled on a SnapLock volume, you can't re-enable it. The default value is `DISABLED` .
	//
	// For more information, see [Privileged delete](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snaplock-enterprise.html#privileged-delete) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html#cfn-fsx-volume-snaplockconfiguration-privilegeddelete
	//
	PrivilegedDelete *string `field:"optional" json:"privilegedDelete" yaml:"privilegedDelete"`
	// Specifies the retention period of an FSx for ONTAP SnapLock volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html#cfn-fsx-volume-snaplockconfiguration-retentionperiod
	//
	RetentionPeriod interface{} `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Enables or disables volume-append mode on an FSx for ONTAP SnapLock volume.
	//
	// Volume-append mode allows you to create WORM-appendable files and write data to them incrementally. The default value is `false` .
	//
	// For more information, see [Volume-append mode](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/worm-state.html#worm-state-append) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockconfiguration.html#cfn-fsx-volume-snaplockconfiguration-volumeappendmodeenabled
	//
	VolumeAppendModeEnabled *string `field:"optional" json:"volumeAppendModeEnabled" yaml:"volumeAppendModeEnabled"`
}

