package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsredshift/internal"
)

// Represents a table in a Redshift database.
// Experimental.
type ITable interface {
	awscdk.IConstruct
	// Grant a user privilege to access this table.
	// Experimental.
	Grant(user IUser, actions ...TableAction)
	// The cluster where the table is located.
	// Experimental.
	Cluster() ICluster
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName() *string
	// The columns of the table.
	// Experimental.
	TableColumns() *[]*Column
	// Name of the table.
	// Experimental.
	TableName() *string
}

// The jsii proxy for ITable
type jsiiProxy_ITable struct {
	internal.Type__awscdkIConstruct
}

func (i *jsiiProxy_ITable) Grant(user IUser, actions ...TableAction) {
	args := []interface{}{user}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"grant",
		args,
	)
}

func (j *jsiiProxy_ITable) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableColumns() *[]*Column {
	var returns *[]*Column
	_jsii_.Get(
		j,
		"tableColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

