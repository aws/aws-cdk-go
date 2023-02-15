package awslambda


// Properties for an ADOT instrumentation in Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_18_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
//   	adotInstrumentation: &adotInstrumentationConfig{
//   		layerVersion: awscdk.AdotLayerVersion.fromJavaScriptSdkLayerVersion(awscdk.AdotLambdaLayerJavaScriptSdkVersion_LATEST()),
//   		execWrapper: awscdk.AdotLambdaExecWrapper_REGULAR_HANDLER,
//   	},
//   })
//
type AdotInstrumentationConfig struct {
	// The startup script to run, see ADOT documentation to pick the right script for your use case: https://aws-otel.github.io/docs/getting-started/lambda.
	ExecWrapper AdotLambdaExecWrapper `field:"required" json:"execWrapper" yaml:"execWrapper"`
	// The ADOT Lambda layer.
	LayerVersion AdotLayerVersion `field:"required" json:"layerVersion" yaml:"layerVersion"`
}

