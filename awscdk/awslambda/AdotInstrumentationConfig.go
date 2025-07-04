package awslambda


// Properties for an ADOT instrumentation in Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
//   	AdotInstrumentation: &AdotInstrumentationConfig{
//   		LayerVersion: awscdk.AdotLayerVersion_FromJavaScriptSdkLayerVersion(awscdk.AdotLambdaLayerJavaScriptSdkVersion_LATEST()),
//   		ExecWrapper: awscdk.AdotLambdaExecWrapper_REGULAR_HANDLER,
//   	},
//   })
//
type AdotInstrumentationConfig struct {
	// The startup script to run, see ADOT documentation to pick the right script for your use case: https://aws-otel.github.io/docs/getting-started/lambda.
	ExecWrapper AdotLambdaExecWrapper `field:"required" json:"execWrapper" yaml:"execWrapper"`
	// The ADOT Lambda layer.
	LayerVersion AdotLayerVersion `field:"required" json:"layerVersion" yaml:"layerVersion"`
}

