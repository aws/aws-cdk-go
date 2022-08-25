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
//   myTable := dynamodb.table.fromTableName(this, jsii.String("Table"), jsii.String("myTableName"))
//   myDatabaseInstance := rds.NewDatabaseInstance(this, jsii.String("DatabaseInstance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
//   		version: rds.mysqlEngineVersion_VER_8_0_26(),
//   	}),
//   	vpc: vpc,
//   })
//   myDatabaseCluster := rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
//   		version: rds.auroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("clusteradmin")),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   	},
//   })
//   myServerlessCluster := rds.NewServerlessCluster(this, jsii.String("ServerlessCluster"), &serverlessClusterProps{
//   	engine: rds.*databaseClusterEngine_AURORA_POSTGRESQL(),
//   	parameterGroup: rds.parameterGroup.fromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql10")),
//   	vpc: vpc,
//   })
//   myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))
//
//   plan.addSelection(jsii.String("Selection"), &backupSelectionOptions{
//   	resources: []backupResource{
//   		backup.*backupResource.fromDynamoDbTable(myTable),
//   		backup.*backupResource.fromRdsDatabaseInstance(myDatabaseInstance),
//   		backup.*backupResource.fromRdsDatabaseCluster(myDatabaseCluster),
//   		backup.*backupResource.fromRdsServerlessCluster(myServerlessCluster),
//   		backup.*backupResource.fromTag(jsii.String("stage"), jsii.String("prod")),
//   		backup.*backupResource.fromConstruct(myCoolConstruct),
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
	AllowRestores *bool `field:"optional" json:"allowRestores" yaml:"allowRestores"`
	// The name for this selection.
	BackupSelectionName *string `field:"optional" json:"backupSelectionName" yaml:"backupSelectionName"`
	// The role that AWS Backup uses to authenticate when backuping or restoring the resources.
	//
	// The `AWSBackupServiceRolePolicyForBackup` managed policy
	// will be attached to this role.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

