package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for a BackupSelection.
//
// Example:
//   var plan BackupPlan
//   var vpc Vpc
//
//   myTable := dynamodb.Table_FromTableName(this, jsii.String("Table"), jsii.String("myTableName"))
//   myDatabaseInstance := rds.NewDatabaseInstance(this, jsii.String("DatabaseInstance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_26(),
//   	}),
//   	Vpc: Vpc,
//   })
//   myDatabaseCluster := rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//   	InstanceProps: &InstanceProps{
//   		Vpc: *Vpc,
//   	},
//   })
//   myServerlessCluster := rds.NewServerlessCluster(this, jsii.String("ServerlessCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
//   	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql11")),
//   	Vpc: Vpc,
//   })
//   myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))
//
//   plan.AddSelection(jsii.String("Selection"), &BackupSelectionOptions{
//   	Resources: []BackupResource{
//   		backup.BackupResource_FromDynamoDbTable(myTable),
//   		backup.BackupResource_FromRdsDatabaseInstance(myDatabaseInstance),
//   		backup.BackupResource_FromRdsDatabaseCluster(myDatabaseCluster),
//   		backup.BackupResource_FromRdsServerlessCluster(myServerlessCluster),
//   		backup.BackupResource_FromTag(jsii.String("stage"), jsii.String("prod")),
//   		backup.BackupResource_FromConstruct(myCoolConstruct),
//   	},
//   })
//
type BackupSelectionOptions struct {
	// The resources to backup.
	//
	// Use the helper static methods defined on `BackupResource`.
	Resources *[]BackupResource `field:"required" json:"resources" yaml:"resources"`
	// Whether to automatically give restores permissions to the role that AWS Backup uses.
	//
	// If `true`, the `AWSBackupServiceRolePolicyForRestores` managed
	// policy will be attached to the role.
	// Default: false.
	//
	AllowRestores *bool `field:"optional" json:"allowRestores" yaml:"allowRestores"`
	// The name for this selection.
	// Default: - a CDK generated name.
	//
	BackupSelectionName *string `field:"optional" json:"backupSelectionName" yaml:"backupSelectionName"`
	// Whether to disable automatically assigning default backup permissions to the role that AWS Backup uses.
	//
	// If `false`, the `AWSBackupServiceRolePolicyForBackup` managed policy will be
	// attached to the role.
	// Default: false.
	//
	DisableDefaultBackupPolicy *bool `field:"optional" json:"disableDefaultBackupPolicy" yaml:"disableDefaultBackupPolicy"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role unless `disableDefaultBackupPolicy`
	// is set to `true`.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

