package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class for specifying specific appsync runtime versions.
//
// Example:
//   var api graphqlApi
//
//
//   myJsFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &appsyncFunctionProps{
//   	name: jsii.String("my_js_function"),
//   	api: api,
//   	dataSource: api.addNoneDataSource(jsii.String("none")),
//   	code: appsync.code.fromAsset(jsii.String("directory/function_code.js")),
//   	runtime: appsync.functionRuntime_JS_1_0_0(),
//   })
//
//   appsync.NewResolver(this, jsii.String("PipelineResolver"), &resolverProps{
//   	api: api,
//   	typeName: jsii.String("typeName"),
//   	fieldName: jsii.String("fieldName"),
//   	code: appsync.*code.fromInline(jsii.String("\n    // The before step\n    export function request(...args) {\n      console.log(args);\n      return {}\n    }\n\n    // The after step\n    export function response(ctx) {\n      return ctx.prev.result\n    }\n  ")),
//   	runtime: appsync.*functionRuntime_JS_1_0_0(),
//   	pipelineConfig: []iAppsyncFunction{
//   		myJsFunction,
//   	},
//   })
//
type FunctionRuntime interface {
	// The name of the runtime.
	Name() *string
	// The runtime version.
	Version() *string
	// Convert to Cfn runtime configuration property format.
	ToProperties() *RuntimeConfig
}

// The jsii proxy struct for FunctionRuntime
type jsiiProxy_FunctionRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_FunctionRuntime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionRuntime) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewFunctionRuntime(family FunctionRuntimeFamily, version *string) FunctionRuntime {
	_init_.Initialize()

	if err := validateNewFunctionRuntimeParameters(family, version); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionRuntime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.FunctionRuntime",
		[]interface{}{family, version},
		&j,
	)

	return &j
}

func NewFunctionRuntime_Override(f FunctionRuntime, family FunctionRuntimeFamily, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.FunctionRuntime",
		[]interface{}{family, version},
		f,
	)
}

func FunctionRuntime_JS_1_0_0() FunctionRuntime {
	_init_.Initialize()
	var returns FunctionRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.FunctionRuntime",
		"JS_1_0_0",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FunctionRuntime) ToProperties() *RuntimeConfig {
	var returns *RuntimeConfig

	_jsii_.Invoke(
		f,
		"toProperties",
		nil, // no parameters
		&returns,
	)

	return returns
}

