package awsbackup


// Specifies an object containing resource type and backup options.
//
// This is only supported for Windows VSS backups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupOptions interface{}
//
//   advancedBackupSettingResourceTypeProperty := &advancedBackupSettingResourceTypeProperty{
//   	backupOptions: backupOptions,
//   	resourceType: jsii.String("resourceType"),
//   }
//
type CfnBackupPlan_AdvancedBackupSettingResourceTypeProperty struct {
	// The backup option for the resource.
	//
	// Each option is a key-value pair.
	BackupOptions interface{} `field:"required" json:"backupOptions" yaml:"backupOptions"`
	// The name of a resource type.
	//
	// The only supported resource type is EC2.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

