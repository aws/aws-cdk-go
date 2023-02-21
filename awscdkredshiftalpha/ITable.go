// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a table in a Redshift database.
// Experimental.
type ITable interface {
	constructs.IConstruct
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
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_ITable) Grant(user IUser, actions ...TableAction) {
	if err := i.validateGrantParameters(user); err != nil {
		panic(err)
	}
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

