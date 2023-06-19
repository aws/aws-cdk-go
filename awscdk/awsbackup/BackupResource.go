package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsefs"
	"github.com/aws/aws-cdk-go/awscdk/awsrds"
	"github.com/aws/constructs-go/constructs/v3"
)

// A resource to backup.
//
// Example:
//   var plan backupPlan
//
//   myTable := dynamodb.Table_FromTableName(this, jsii.String("Table"), jsii.String("myTableName"))
//   myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))
//
//   plan.AddSelection(jsii.String("Selection"), &BackupSelectionOptions{
//   	Resources: []backupResource{
//   		backup.*backupResource_FromDynamoDbTable(myTable),
//   		backup.*backupResource_FromTag(jsii.String("stage"), jsii.String("prod")),
//   		backup.*backupResource_FromConstruct(myCoolConstruct),
//   	},
//   })
//
// Experimental.
type BackupResource interface {
	// A construct.
	// Experimental.
	Construct() awscdk.Construct
	// A resource.
	// Experimental.
	Resource() *string
	// A condition on a tag.
	// Experimental.
	TagCondition() *TagCondition
}

// The jsii proxy struct for BackupResource
type jsiiProxy_BackupResource struct {
	_ byte // padding
}

func (j *jsiiProxy_BackupResource) Construct() awscdk.Construct {
	var returns awscdk.Construct
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


// Experimental.
func NewBackupResource(resource *string, tagCondition *TagCondition, construct constructs.Construct) BackupResource {
	_init_.Initialize()

	if err := validateNewBackupResourceParameters(tagCondition); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupResource{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupResource",
		[]interface{}{resource, tagCondition, construct},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupResource_Override(b BackupResource, resource *string, tagCondition *TagCondition, construct constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupResource",
		[]interface{}{resource, tagCondition, construct},
		b,
	)
}

// A list of ARNs or match patterns such as `arn:aws:ec2:us-east-1:123456789012:volume/*`.
// Experimental.
func BackupResource_FromArn(arn *string) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromArnParameters(arn); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

// Adds all supported resources in a construct.
// Experimental.
func BackupResource_FromConstruct(construct constructs.Construct) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromConstructParameters(construct); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromConstruct",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// A DynamoDB table.
// Experimental.
func BackupResource_FromDynamoDbTable(table awsdynamodb.ITable) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromDynamoDbTableParameters(table); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromDynamoDbTable",
		[]interface{}{table},
		&returns,
	)

	return returns
}

// An EC2 instance.
// Experimental.
func BackupResource_FromEc2Instance(instance awsec2.IInstance) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromEc2InstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromEc2Instance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// An EFS file system.
// Experimental.
func BackupResource_FromEfsFileSystem(fileSystem awsefs.IFileSystem) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromEfsFileSystemParameters(fileSystem); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromEfsFileSystem",
		[]interface{}{fileSystem},
		&returns,
	)

	return returns
}

// A RDS database instance.
// Experimental.
func BackupResource_FromRdsDatabaseInstance(instance awsrds.IDatabaseInstance) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromRdsDatabaseInstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromRdsDatabaseInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// A tag condition.
// Experimental.
func BackupResource_FromTag(key *string, value *string, operation TagOperation) BackupResource {
	_init_.Initialize()

	if err := validateBackupResource_FromTagParameters(key, value); err != nil {
		panic(err)
	}
	var returns BackupResource

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupResource",
		"fromTag",
		[]interface{}{key, value, operation},
		&returns,
	)

	return returns
}

