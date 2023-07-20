package awslambda


// The wrapper script to be used for the Lambda function in order to enable auto instrumentation with ADOT.
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
type AdotLambdaExecWrapper string

const (
	// Wrapping regular Lambda handlers.
	AdotLambdaExecWrapper_REGULAR_HANDLER AdotLambdaExecWrapper = "REGULAR_HANDLER"
	// Wrapping regular handlers (implementing RequestHandler) proxied through API Gateway, enabling HTTP context propagation.
	AdotLambdaExecWrapper_PROXY_HANDLER AdotLambdaExecWrapper = "PROXY_HANDLER"
	// Wrapping streaming handlers (implementing RequestStreamHandler), enabling HTTP context propagation for HTTP requests.
	AdotLambdaExecWrapper_STREAM_HANDLER AdotLambdaExecWrapper = "STREAM_HANDLER"
	// Wrapping python lambda handlers see https://aws-otel.github.io/docs/getting-started/lambda/lambda-python.
	AdotLambdaExecWrapper_INSTRUMENT_HANDLER AdotLambdaExecWrapper = "INSTRUMENT_HANDLER"
)

