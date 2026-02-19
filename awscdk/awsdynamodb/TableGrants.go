package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
)

// A set of permissions to grant on a Table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryptedResource IEncryptedResource
//   var resourceWithPolicyV2 IResourceWithPolicyV2
//   var tableRef ITableRef
//
//   tableGrants := awscdk.Aws_dynamodb.NewTableGrants(&TableGrantsProps{
//   	Table: tableRef,
//
//   	// the properties below are optional
//   	EncryptedResource: encryptedResource,
//   	HasIndex: jsii.Boolean(false),
//   	PolicyResource: resourceWithPolicyV2,
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   })
//
type TableGrants interface {
	// Adds an IAM policy statement associated with this table to an IAM principal's policy.
	//
	// If `encryptionKey` is present, appropriate grants to the key needs to be added
	// separately using the `table.encryptionKey.grant*` methods.
	Actions(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits all DynamoDB operations ("dynamodb:*") to an IAM principal.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	FullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Grants permissions for this table to act as a destination for multi-account global table replication.
	MultiAccountReplicationFrom(sourceReplicaArn *string)
	// Grants permissions for this table to act as a source for multi-account global table replication.
	MultiAccountReplicationTo(destinationReplicaArn *string)
	// Permits an IAM principal all data read operations from this table: BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan, DescribeTable.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	ReadData(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal to all data read/write operations to this table.
	//
	// BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan,
	// BatchWriteItem, PutItem, UpdateItem, DeleteItem, DescribeTable
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	ReadWriteData(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all data write operations to this table: BatchWriteItem, PutItem, UpdateItem, DeleteItem, DescribeTable.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	WriteData(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for TableGrants
type jsiiProxy_TableGrants struct {
	_ byte // padding
}

func NewTableGrants(props *TableGrantsProps) TableGrants {
	_init_.Initialize()

	if err := validateNewTableGrantsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TableGrants{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.TableGrants",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewTableGrants_Override(t TableGrants, props *TableGrantsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.TableGrants",
		[]interface{}{props},
		t,
	)
}

// Creates a TableGrants object for a given table.
func TableGrants_FromTable(table interfacesawsdynamodb.ITableRef, regions *[]*string, hasIndex *bool) TableGrants {
	_init_.Initialize()

	if err := validateTableGrants_FromTableParameters(table); err != nil {
		panic(err)
	}
	var returns TableGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.TableGrants",
		"fromTable",
		[]interface{}{table, regions, hasIndex},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableGrants) Actions(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := t.validateActionsParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"actions",
		args,
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableGrants) FullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validateFullAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"fullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableGrants) MultiAccountReplicationFrom(sourceReplicaArn *string) {
	if err := t.validateMultiAccountReplicationFromParameters(sourceReplicaArn); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"multiAccountReplicationFrom",
		[]interface{}{sourceReplicaArn},
	)
}

func (t *jsiiProxy_TableGrants) MultiAccountReplicationTo(destinationReplicaArn *string) {
	if err := t.validateMultiAccountReplicationToParameters(destinationReplicaArn); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"multiAccountReplicationTo",
		[]interface{}{destinationReplicaArn},
	)
}

func (t *jsiiProxy_TableGrants) ReadData(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validateReadDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"readData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableGrants) ReadWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validateReadWriteDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"readWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableGrants) WriteData(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validateWriteDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"writeData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

