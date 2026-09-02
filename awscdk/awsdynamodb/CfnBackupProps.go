package awsdynamodb


// Properties for defining a `CfnBackup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBackupProps := &CfnBackupProps{
//   	BackupName: jsii.String("backupName"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-backup.html
//
type CfnBackupProps struct {
	// The name for the backup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-backup.html#cfn-dynamodb-backup-backupname
	//
	BackupName *string `field:"required" json:"backupName" yaml:"backupName"`
	// The name of the table to back up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-backup.html#cfn-dynamodb-backup-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

