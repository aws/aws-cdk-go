package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents source code for an AppSync Function or Resolver.
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
type Code interface {
	// Bind source code to an AppSync Function or resolver.
	Bind(scope constructs.Construct) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.Code",
		nil, // no parameters
		c,
	)
}

// Loads the function code from a local disk path.
func Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Inline code for AppSync function.
//
// Returns: `InlineCode` with inline code.
func Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

