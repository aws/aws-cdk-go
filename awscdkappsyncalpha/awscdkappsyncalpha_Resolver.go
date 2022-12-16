// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync resolver.
//
// Example:
//   var api graphqlApi
//   var appsyncFunction appsyncFunction
//
//
//   pipelineResolver := appsync.NewResolver(this, jsii.String("pipeline"), &resolverProps{
//   	api: api,
//   	dataSource: api.addNoneDataSource(jsii.String("none")),
//   	typeName: jsii.String("typeName"),
//   	fieldName: jsii.String("fieldName"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("beforeRequest.vtl")),
//   	pipelineConfig: []iAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("afterResponse.vtl")),
//   })
//
// Experimental.
type Resolver interface {
	constructs.Construct
	// the ARN of the resolver.
	// Experimental.
	Arn() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
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


// Experimental.
func NewResolver(scope constructs.Construct, id *string, props *ResolverProps) Resolver {
	_init_.Initialize()

	if err := validateNewResolverParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Resolver{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Resolver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewResolver_Override(r Resolver, scope constructs.Construct, id *string, props *ResolverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.Resolver",
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
// Experimental.
func Resolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResolver_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.Resolver",
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

