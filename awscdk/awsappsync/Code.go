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
//   	PipelineConfig: []iAppsyncFunction{
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

