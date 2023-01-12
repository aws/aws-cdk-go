package awslambda


// The wrapper script to be used for the Lambda function in order to enable auto instrumentation with ADOT.
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
type AdotLambdaExecWrapper string

const (
	// Wrapping regular Lambda handlers.
	AdotLambdaExecWrapper_REGULAR_HANDLER AdotLambdaExecWrapper = "REGULAR_HANDLER"
	// Wrapping regular handlers (implementing RequestHandler) proxied through API Gateway, enabling HTTP context propagation.
	AdotLambdaExecWrapper_PROXY_HANDLER AdotLambdaExecWrapper = "PROXY_HANDLER"
	// Wrapping streaming handlers (implementing RequestStreamHandler), enabling HTTP context propagation for HTTP requests.
	AdotLambdaExecWrapper_STREAM_HANDLER AdotLambdaExecWrapper = "STREAM_HANDLER"
)

