package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsefs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A resource to backup.
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
func BackupResource_FromDynamoDbTable(table interfacesawsdynamodb.ITableRef) BackupResource {
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
func BackupResource_FromEc2Instance(instance interfacesawsec2.IInstanceRef) BackupResource {
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
func BackupResource_FromEfsFileSystem(fileSystem interfacesawsefs.IFileSystemRef) BackupResource {
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

