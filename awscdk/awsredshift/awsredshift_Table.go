package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A table in a Redshift cluster.
//
// Example:
//   awscdk.NewTable(this, jsii.String("Table"), &TableProps{
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
//   	DistStyle: awscdk.TableDistStyle_KEY,
//   })
//
// Experimental.
type Table interface {
	awscdk.Construct
	ITable
	// The cluster where the table is located.
	// Experimental.
	Cluster() ICluster
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for Table
type jsiiProxy_Table struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_Table) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_redshift.Table",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTable_Override(t Table, scope constructs.Construct, id *string, props *TableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_redshift.Table",
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
		"monocdk.aws_redshift.Table",
		"fromTableAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Table_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshift.Table",
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

func (t *jsiiProxy_Table) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Table) OnSynthesize(session constructs.ISynthesisSession) {
	if err := t.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Table) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Table) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Table) Synthesize(session awscdk.ISynthesisSession) {
	if err := t.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
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

func (t *jsiiProxy_Table) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

