package previewawsbackupmixins


// Specifies an object containing resource type and backup options.
//
// This is only supported for Windows VSS backups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var backupOptions interface{}
//
//   advancedBackupSettingResourceTypeProperty := &AdvancedBackupSettingResourceTypeProperty{
//   	BackupOptions: backupOptions,
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-advancedbackupsettingresourcetype.html
//
type CfnBackupPlanPropsMixin_AdvancedBackupSettingResourceTypeProperty struct {
	// The backup option for the resource.
	//
	// Each option is a key-value pair. This option is only available for Windows VSS backup jobs.
	//
	// Valid values:
	//
	// Set to `"WindowsVSS":"enabled"` to enable the `WindowsVSS` backup option and create a Windows VSS backup.
	//
	// Set to `"WindowsVSS":"disabled"` to create a regular backup. The `WindowsVSS` option is not enabled by default.
	//
	// If you specify an invalid option, you get an `InvalidParameterValueException` exception.
	//
	// For more information about Windows VSS backups, see [Creating a VSS-Enabled Windows Backup](https://docs.aws.amazon.com/aws-backup/latest/devguide/windows-backups.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-advancedbackupsettingresourcetype.html#cfn-backup-backupplan-advancedbackupsettingresourcetype-backupoptions
	//
	BackupOptions interface{} `field:"optional" json:"backupOptions" yaml:"backupOptions"`
	// The name of a resource type.
	//
	// The only supported resource type is EC2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-advancedbackupsettingresourcetype.html#cfn-backup-backupplan-advancedbackupsettingresourcetype-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

