// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A table in a Redshift cluster.
//
// Example:
//   awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
//   	TableColumns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   			DistKey: jsii.Boolean(true),
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   	DistStyle: awscdkredshiftalpha.TableDistStyle_KEY,
//   })
//
// Experimental.
type Table interface {
	constructs.Construct
	ITable
	// The cluster where the table is located.
	// Experimental.
	Cluster() ICluster
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The columns of the table.
	// Experimental.
	TableColumns() *[]*Column
	// Name of the table.
	// Experimental.
	TableName() *string
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be destroyed (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	//
	// This resource is retained by default.
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Grant a user privilege to access this table.
	// Experimental.
	Grant(user IUser, actions ...TableAction)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Table
type jsiiProxy_Table struct {
	internal.Type__constructsConstruct
	jsiiProxy_ITable
}

func (j *jsiiProxy_Table) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableColumns() *[]*Column {
	var returns *[]*Column
	_jsii_.Get(
		j,
		"tableColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTable(scope constructs.Construct, id *string, props *TableProps) Table {
	_init_.Initialize()

	if err := validateNewTableParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Table{}

	_jsii_.Create(
		"@aws-cdk/aws-redshift-alpha.Table",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTable_Override(t Table, scope constructs.Construct, id *string, props *TableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-redshift-alpha.Table",
		[]interface{}{scope, id, props},
		t,
	)
}

// Specify a Redshift table using a table name and schema that already exists.
// Experimental.
func Table_FromTableAttributes(scope constructs.Construct, id *string, attrs *TableAttributes) ITable {
	_init_.Initialize()

	if err := validateTable_FromTableAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ITable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-redshift-alpha.Table",
		"fromTableAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func Table_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-redshift-alpha.Table",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_Table) Grant(user IUser, actions ...TableAction) {
	if err := t.validateGrantParameters(user); err != nil {
		panic(err)
	}
	args := []interface{}{user}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"grant",
		args,
	)
}

func (t *jsiiProxy_Table) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

