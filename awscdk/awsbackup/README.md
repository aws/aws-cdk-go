# AWS Backup Construct Library

AWS Backup is a fully managed backup service that makes it easy to centralize and automate the
backup of data across AWS services in the cloud and on premises. Using AWS Backup, you can
configure backup policies and monitor backup activity for your AWS resources in one place.

## Backup plan and selection

In AWS Backup, a *backup plan* is a policy expression that defines when and how you want to back up
your AWS resources, such as Amazon DynamoDB tables or Amazon Elastic File System (Amazon EFS) file
systems. You can assign resources to backup plans, and AWS Backup automatically backs up and retains
backups for those resources according to the backup plan. You can create multiple backup plans if you
have workloads with different backup requirements.

This module provides ready-made backup plans (similar to the console experience):

```go
// Daily, weekly and monthly with 5 year retention
plan := backup.BackupPlan_DailyWeeklyMonthly5YearRetention(this, jsii.String("Plan"))
```

Assigning resources to a plan can be done with `addSelection()`:

```go
var plan backupPlan
var vpc vpc

myTable := dynamodb.Table_FromTableName(this, jsii.String("Table"), jsii.String("myTableName"))
myDatabaseInstance := rds.NewDatabaseInstance(this, jsii.String("DatabaseInstance"), &DatabaseInstanceProps{
	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
		Version: rds.MysqlEngineVersion_VER_8_0_26(),
	}),
	Vpc: Vpc,
})
myDatabaseCluster := rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &DatabaseClusterProps{
	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
	}),
	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
	InstanceProps: &InstanceProps{
		Vpc: *Vpc,
	},
})
myServerlessCluster := rds.NewServerlessCluster(this, jsii.String("ServerlessCluster"), &ServerlessClusterProps{
	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql10")),
	Vpc: Vpc,
})
myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))

plan.AddSelection(jsii.String("Selection"), &BackupSelectionOptions{
	Resources: []backupResource{
		backup.*backupResource_FromDynamoDbTable(myTable),
		backup.*backupResource_FromRdsDatabaseInstance(myDatabaseInstance),
		backup.*backupResource_FromRdsDatabaseCluster(myDatabaseCluster),
		backup.*backupResource_FromRdsServerlessCluster(myServerlessCluster),
		backup.*backupResource_FromTag(jsii.String("stage"), jsii.String("prod")),
		backup.*backupResource_FromConstruct(myCoolConstruct),
	},
})
```

If not specified, a new IAM role with a managed policy for backup will be
created for the selection. The `BackupSelection` implements `IGrantable`.

To add rules to a plan, use `addRule()`:

```go
var plan backupPlan

plan.AddRule(backup.NewBackupPlanRule(&BackupPlanRuleProps{
	CompletionWindow: awscdk.Duration_Hours(jsii.Number(2)),
	StartWindow: awscdk.Duration_*Hours(jsii.Number(1)),
	ScheduleExpression: events.Schedule_Cron(&CronOptions{
		 // Only cron expressions are supported
		Day: jsii.String("15"),
		Hour: jsii.String("3"),
		Minute: jsii.String("30"),
	}),
	MoveToColdStorageAfter: awscdk.Duration_Days(jsii.Number(30)),
}))
```

Continuous backup and point-in-time restores (PITR) can be configured.
Property `deleteAfter` defines the retention period for the backup. It is mandatory if PITR is enabled.
If no value is specified, the retention period is set to 35 days which is the maximum retention period supported by PITR.
Property `moveToColdStorageAfter` must not be specified because PITR does not support this option.
This example defines an AWS Backup rule with PITR and a retention period set to 14 days:

```go
var plan backupPlan

plan.AddRule(backup.NewBackupPlanRule(&BackupPlanRuleProps{
	EnableContinuousBackup: jsii.Boolean(true),
	DeleteAfter: awscdk.Duration_Days(jsii.Number(14)),
}))
```

Rules can also specify to copy recovery points to another Backup Vault using `copyActions`. Copied recovery points can
optionally have `moveToColdStorageAfter` and `deleteAfter` configured.

```go
var plan backupPlan
var secondaryVault backupVault

plan.AddRule(backup.NewBackupPlanRule(&BackupPlanRuleProps{
	CopyActions: []backupPlanCopyActionProps{
		&backupPlanCopyActionProps{
			DestinationBackupVault: secondaryVault,
			MoveToColdStorageAfter: awscdk.Duration_Days(jsii.Number(30)),
			DeleteAfter: awscdk.Duration_*Days(jsii.Number(120)),
		},
	},
}))
```

Ready-made rules are also available:

```go
var plan backupPlan

plan.AddRule(backup.BackupPlanRule_Daily())
plan.AddRule(backup.BackupPlanRule_Weekly())
```

By default a new [vault](#Backup-vault) is created when creating a plan.
It is also possible to specify a vault either at the plan level or at the
rule level.

```go
myVault := backup.BackupVault_FromBackupVaultName(this, jsii.String("Vault1"), jsii.String("myVault"))
otherVault := backup.BackupVault_FromBackupVaultName(this, jsii.String("Vault2"), jsii.String("otherVault"))

plan := backup.BackupPlan_Daily35DayRetention(this, jsii.String("Plan"), myVault) // Use `myVault` for all plan rules
plan.AddRule(backup.BackupPlanRule_Monthly1Year(otherVault))
```

You can [backup](https://docs.aws.amazon.com/aws-backup/latest/devguide/windows-backups.html)
VSS-enabled Windows applications running on Amazon EC2 instances by setting the `windowsVss`
parameter to `true`. If the application has VSS writer registered with Windows VSS,
then AWS Backup creates a snapshot that will be consistent for that application.

```go
plan := backup.NewBackupPlan(this, jsii.String("Plan"), &BackupPlanProps{
	WindowsVss: jsii.Boolean(true),
})
```

## Backup vault

In AWS Backup, a *backup vault* is a container that you organize your backups in. You can use backup
vaults to set the AWS Key Management Service (AWS KMS) encryption key that is used to encrypt backups
in the backup vault and to control access to the backups in the backup vault. If you require different
encryption keys or access policies for different groups of backups, you can optionally create multiple
backup vaults.

```go
myKey := kms.Key_FromKeyArn(this, jsii.String("MyKey"), jsii.String("aaa"))
myTopic := sns.Topic_FromTopicArn(this, jsii.String("MyTopic"), jsii.String("bbb"))

vault := backup.NewBackupVault(this, jsii.String("Vault"), &BackupVaultProps{
	EncryptionKey: myKey,
	 // Custom encryption key
	NotificationTopic: myTopic,
})
```

A vault has a default `RemovalPolicy` set to `RETAIN`. Note that removing a vault
that contains recovery points will fail.

You can assign policies to backup vaults and the resources they contain. Assigning policies allows
you to do things like grant access to users to create backup plans and on-demand backups, but limit
their ability to delete recovery points after they're created.

Use the `accessPolicy` property to create a backup vault policy:

```go
vault := backup.NewBackupVault(this, jsii.String("Vault"), &BackupVaultProps{
	AccessPolicy: iam.NewPolicyDocument(&PolicyDocumentProps{
		Statements: []policyStatement{
			iam.NewPolicyStatement(&PolicyStatementProps{
				Effect: iam.Effect_DENY,
				Principals: []iPrincipal{
					iam.NewAnyPrincipal(),
				},
				Actions: []*string{
					jsii.String("backup:DeleteRecoveryPoint"),
				},
				Resources: []*string{
					jsii.String("*"),
				},
				Conditions: map[string]interface{}{
					"StringNotLike": map[string][]*string{
						"aws:userId": []*string{
							jsii.String("user1"),
							jsii.String("user2"),
						},
					},
				},
			}),
		},
	}),
})
```

Alternativately statements can be added to the vault policy using `addToAccessPolicy()`.

Use the `blockRecoveryPointDeletion` property or the `blockRecoveryPointDeletion()` method to add
a statement to the vault access policy that prevents recovery point deletions in your vault:

```go
var backupVault backupVault
backup.NewBackupVault(this, jsii.String("Vault"), &BackupVaultProps{
	BlockRecoveryPointDeletion: jsii.Boolean(true),
})
backupVault.BlockRecoveryPointDeletion()
```

By default access is not restricted.

Use the `lockConfiguration` property to enable [AWS Backup Vault Lock](https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html):

```go
// Example automatically generated from non-compiling source. May contain errors.
NewBackupVault(stack, jsii.String("Vault"), map[string]map[string]duration{
	"lockConfiguration": map[string]duration{
		"minRetention": awscdk.Duration_days(jsii.Number(30)),
	},
})
```

## Importing existing backup vault

To import an existing backup vault into your CDK application, use the `BackupVault.fromBackupVaultArn` or `BackupVault.fromBackupVaultName`
static method. Here is an example of giving an IAM Role permission to start a backup job:

```go
importedVault := backup.BackupVault_FromBackupVaultName(this, jsii.String("Vault"), jsii.String("myVaultName"))

role := iam.NewRole(this, jsii.String("Access Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

importedVault.Grant(role, jsii.String("backup:StartBackupJob"))
```
