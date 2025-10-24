package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync resolver.
//
// Example:
//   var api GraphqlApi
//   var appsyncFunction AppsyncFunction
//
//
//   pipelineResolver := appsync.NewResolver(this, jsii.String("pipeline"), &ResolverProps{
//   	Api: Api,
//   	DataSource: api.AddNoneDataSource(jsii.String("none")),
//   	TypeName: jsii.String("typeName"),
//   	FieldName: jsii.String("fieldName"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("beforeRequest.vtl")),
//   	PipelineConfig: []IAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("afterResponse.vtl")),
//   })
//
type Resolver interface {
	constructs.Construct
	// the ARN of the resolver.
	Arn() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Resolver
type jsiiProxy_Resolver struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Resolver) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resolver) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewResolver(scope constructs.Construct, id *string, props *ResolverProps) Resolver {
	_init_.Initialize()

	if err := validateNewResolverParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Resolver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Resolver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewResolver_Override(r Resolver, scope constructs.Construct, id *string, props *ResolverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Resolver",
		[]interface{}{scope, id, props},
		r,
	)
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
func Resolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResolver_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Resolver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resolver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

