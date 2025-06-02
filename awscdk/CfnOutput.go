package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   var cluster cluster
//
//   // add service account
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.GrantReadWrite(serviceAccount)
//
//   mypod := cluster.addManifest(jsii.String("mypod"), map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Pod"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("mypod"),
//   	},
//   	"spec": map[string]interface{}{
//   		"serviceAccountName": serviceAccount.serviceAccountName,
//   		"containers": []map[string]interface{}{
//   			map[string]interface{}{
//   				"name": jsii.String("hello"),
//   				"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   				"ports": []map[string]*f64{
//   					map[string]*f64{
//   						"containerPort": jsii.Number(8080),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // create the resource after the service account.
//   mypod.Node.AddDependency(serviceAccount)
//
//   // print the IAM role arn for this service account
//   // print the IAM role arn for this service account
//   awscdk.NewCfnOutput(this, jsii.String("ServiceAccountIamRole"), &CfnOutputProps{
//   	Value: serviceAccount.Role.RoleArn,
//   })
//
type CfnOutput interface {
	CfnElement
	// A condition to associate with this output value.
	//
	// If the condition evaluates
	// to `false`, this output value will not be included in the stack.
	// Default: - No condition is associated with the output.
	//
	Condition() CfnCondition
	SetCondition(val CfnCondition)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A String type that describes the output value.
	//
	// The description can be a maximum of 4 K in length.
	// Default: - No description.
	//
	Description() *string
	SetDescription(val *string)
	// The name used to export the value of this output across stacks.
	//
	// To use the value in another stack, pass the value of
	// `output.importValue` to it.
	// Default: - the output is not exported.
	//
	ExportName() *string
	SetExportName(val *string)
	// Return the `Fn.importValue` expression to import this value into another stack.
	//
	// The returned value should not be used in the same stack, but in a
	// different one. It must be deployed to the same environment, as
	// CloudFormation exports can only be imported in the same Region and
	// account.
	//
	// The is no automatic registration of dependencies between stacks when using
	// this mechanism, so you should make sure to deploy them in the right order
	// yourself.
	//
	// You can use this mechanism to share values across Stacks in different
	// Stages. If you intend to share the value to another Stack inside the same
	// Stage, the automatic cross-stack referencing mechanism is more convenient.
	ImportValue() *string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() Stack
	// The value of the property returned by the aws cloudformation describe-stacks command.
	//
	// The value of an output can include literals, parameter references, pseudo-parameters,
	// a mapping value, or intrinsic functions.
	Value() interface{}
	SetValue(val interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CfnOutput
type jsiiProxy_CfnOutput struct {
	jsiiProxy_CfnElement
}

func (j *jsiiProxy_CfnOutput) Condition() CfnCondition {
	var returns CfnCondition
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) ExportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) ImportValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOutput) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates an CfnOutput value for this stack.
func NewCfnOutput(scope constructs.Construct, id *string, props *CfnOutputProps) CfnOutput {
	_init_.Initialize()

	if err := validateNewCfnOutputParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOutput{}

	_jsii_.Create(
		"aws-cdk-lib.CfnOutput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an CfnOutput value for this stack.
func NewCfnOutput_Override(c CfnOutput, scope constructs.Construct, id *string, props *CfnOutputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CfnOutput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOutput)SetCondition(val CfnCondition) {
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_CfnOutput)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnOutput)SetExportName(val *string) {
	_jsii_.Set(
		j,
		"exportName",
		val,
	)
}

func (j *jsiiProxy_CfnOutput)SetValue(val interface{}) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOutput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOutput_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnOutput",
		"isCfnElement",
		[]interface{}{x},
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
func CfnOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOutput_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CfnOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOutput) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

