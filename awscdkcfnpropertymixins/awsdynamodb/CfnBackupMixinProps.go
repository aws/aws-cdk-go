package awsdynamodb


// Properties for CfnBackupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBackupMixinProps := &CfnBackupMixinProps{
//   	BackupName: jsii.String("backupName"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-backup.html
//
type CfnBackupMixinProps struct {
	// The name for the backup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-backup.html#cfn-dynamodb-backup-backupname
	//
	BackupName *string `field:"optional" json:"backupName" yaml:"backupName"`
	// The name of the table to back up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-backup.html#cfn-dynamodb-backup-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

