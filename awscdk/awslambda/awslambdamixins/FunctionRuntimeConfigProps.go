package awslambdamixins


// Properties for the FunctionRuntimeConfig mixin.
//
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
type FunctionRuntimeConfigProps struct {
	// The name of the runtime alias that selects which runtime version the function runs on (for example `stable` or `latest`).
	//
	// With the default `RuntimeManagementMode.AUTO`, moving the alias to a new
	// version rolls the function forward on its next cold start, without a
	// function redeploy.
	RuntimeAlias *string `field:"required" json:"runtimeAlias" yaml:"runtimeAlias"`
	// The ARN of the customer-managed runtime that the function uses, for example `arn:aws:lambda:us-east-1:123456789012:runtime/my-runtime`.
	//
	// The specific runtime version is resolved by the service from `runtimeAlias`.
	RuntimeArn *string `field:"required" json:"runtimeArn" yaml:"runtimeArn"`
}

