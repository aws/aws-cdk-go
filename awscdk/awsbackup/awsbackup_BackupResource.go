package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/constructs-go/constructs/v10"
)

// A resource to backup.
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
type BackupResource interface {
	// A construct.
	Construct() constructs.Construct
	// A resource.
	Resource() *string
	// A condition on a tag.
	TagCondition() *TagCondition
}

// The jsii proxy struct for BackupResource
type jsiiProxy_BackupResource struct {
	_ byte // padding
}

func (j *jsiiProxy_BackupResource) Construct() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"construct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupResource) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupResource) TagCondition() *TagCondition {
	var returns *TagCondition
	_jsii_.Get(
		j,
		"tagCondition",
		&returns,
	)
	return returns
}


func NewBackupResource(resource *string, tagCondition *TagCondition, construct constructs.Construct) BackupResource {
	_init_.Initialize()

	if err := validateNewBackupResourceParameters(tagCondition); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupResource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.BackupResource",
		[]interface{}{resource, tagCondition, construct},
		&j,
	)

	return &j
}

func NewBackupResource_Override(b BackupResource, resource *string, tagCondition *TagCondition, construct constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.BackupResource",
		[]interface{}{resource, tagCondition, construct},
		b,
	)
}

// A list of ARNs or match patterns such as `arn:aws:ec2:us-east-1:123456789012:volume/*`.
func BackupResource_FromArn(arn *string) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromArnParameters(arn); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

// Adds all supported resources in a construct.
func BackupResource_FromConstruct(construct constructs.Construct) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromConstructParameters(construct); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromConstruct",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// A DynamoDB table.
func BackupResource_FromDynamoDbTable(table awsdynamodb.ITable) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromDynamoDbTableParameters(table); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromDynamoDbTable",
		[]interface{}{table},
		&returns,
	)

	return returns
}

// An EC2 instance.
func BackupResource_FromEc2Instance(instance awsec2.IInstance) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromEc2InstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromEc2Instance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// An EFS file system.
func BackupResource_FromEfsFileSystem(fileSystem awsefs.IFileSystem) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromEfsFileSystemParameters(fileSystem); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromEfsFileSystem",
		[]interface{}{fileSystem},
		&returns,
	)

	return returns
}

// A RDS database cluter.
func BackupResource_FromRdsDatabaseCluster(cluster awsrds.IDatabaseCluster) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromRdsDatabaseClusterParameters(cluster); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromRdsDatabaseCluster",
		[]interface{}{cluster},
		&returns,
	)

	return returns
}

// A RDS database instance.
func BackupResource_FromRdsDatabaseInstance(instance awsrds.IDatabaseInstance) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromRdsDatabaseInstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromRdsDatabaseInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// An Aurora database instance.
func BackupResource_FromRdsServerlessCluster(cluster awsrds.IServerlessCluster) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromRdsServerlessClusterParameters(cluster); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromRdsServerlessCluster",
		[]interface{}{cluster},
		&returns,
	)

	return returns
}

// A tag condition.
func BackupResource_FromTag(key *string, value *string, operation TagOperation) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromTagParameters(key, value); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupResource",
		"fromTag",
		[]interface{}{key, value, operation},
		&returns,
	)

	return returns
}

