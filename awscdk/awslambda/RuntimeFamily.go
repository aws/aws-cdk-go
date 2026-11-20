package awslambda


// Example:
//   lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.NewRuntime(jsii.String("custom"), lambda.RuntimeFamily_OTHER),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   }).With(mixins.NewFunctionRuntimeConfig(&FunctionRuntimeConfigProps{
//   	RuntimeArn: jsii.String("arn:aws:lambda:us-east-1:123456789012:runtime/my-python-runtime"),
//   	RuntimeAlias: jsii.String("stable"),
//   }))
//
type RuntimeFamily string

const (
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	RuntimeFamily_JAVA RuntimeFamily = "JAVA"
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	RuntimeFamily_DOTNET_CORE RuntimeFamily = "DOTNET_CORE"
	RuntimeFamily_GO RuntimeFamily = "GO"
	RuntimeFamily_RUBY RuntimeFamily = "RUBY"
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

