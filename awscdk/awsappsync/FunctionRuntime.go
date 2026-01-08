package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class for specifying specific appsync runtime versions.
//
// Example:
//   var api GraphqlApi
//
//
//   myJsFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &AppsyncFunctionProps{
//   	Name: jsii.String("my_js_function"),
//   	Api: Api,
//   	DataSource: api.AddNoneDataSource(jsii.String("none")),
//   	Code: appsync.Code_FromAsset(jsii.String("directory/function_code.js")),
//   	Runtime: appsync.FunctionRuntime_JS_1_0_0(),
//   })
//
//   appsync.NewResolver(this, jsii.String("PipelineResolver"), &ResolverProps{
//   	Api: Api,
//   	TypeName: jsii.String("typeName"),
//   	FieldName: jsii.String("fieldName"),
//   	Code: appsync.Code_FromInline(jsii.String(`
//   	    // The before step
//   	    export function request(...args) {
//   	      console.log(args);
//   	      return {}
//   	    }
//
//   	    // The after step
//   	    export function response(ctx) {
//   	      return ctx.prev.result
//   	    }
//   	  `)),
//   	Runtime: appsync.FunctionRuntime_JS_1_0_0(),
//   	PipelineConfig: []IFunctionConfigurationRef{
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

