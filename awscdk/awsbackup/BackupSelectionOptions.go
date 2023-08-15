package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for a BackupSelection.
//
// Example:
//   var plan backupPlan
//   var vpc vpc
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
//   	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql10")),
//   	Vpc: Vpc,
//   })
//   myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))
//
//   plan.AddSelection(jsii.String("Selection"), &BackupSelectionOptions{
//   	Resources: []backupResource{
//   		backup.*backupResource_FromDynamoDbTable(myTable),
//   		backup.*backupResource_FromRdsDatabaseInstance(myDatabaseInstance),
//   		backup.*backupResource_FromRdsDatabaseCluster(myDatabaseCluster),
//   		backup.*backupResource_FromRdsServerlessCluster(myServerlessCluster),
//   		backup.*backupResource_FromTag(jsii.String("stage"), jsii.String("prod")),
//   		backup.*backupResource_FromConstruct(myCoolConstruct),
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
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

