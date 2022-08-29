// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a user in a Redshift database.
// Experimental.
type IUser interface {
	constructs.IConstruct
	// Grant this user privilege to access a table.
	// Experimental.
	AddTablePrivileges(table ITable, actions ...TableAction)
	// The cluster where the table is located.
	// Experimental.
	Cluster() ICluster
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName() *string
	// The password of the user.
	// Experimental.
	Password() awscdk.SecretValue
	// The name of the user.
	// Experimental.
	Username() *string
}

// The jsii proxy for IUser
type jsiiProxy_IUser struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IUser) AddTablePrivileges(table ITable, actions ...TableAction) {
	args := []interface{}{table}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTablePrivileges",
		args,
	)
}

func (j *jsiiProxy_IUser) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

