package awslambda


// X-Ray Tracing Modes (https://docs.aws.amazon.com/lambda/latest/dg/API_TracingConfig.html).
//
// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
//   	Tracing: lambda.Tracing_ACTIVE,
//   })
//
type Tracing string

const (
	// Lambda will respect any tracing header it receives from an upstream service.
	//
	// If no tracing header is received, Lambda will sample the request based on a fixed rate. Please see the [Using AWS Lambda with AWS X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) documentation for details on this sampling behavior.
	Tracing_ACTIVE Tracing = "ACTIVE"
	// Lambda will only trace the request from an upstream service if it contains a tracing header with "sampled=1".
	Tracing_PASS_THROUGH Tracing = "PASS_THROUGH"
	// Lambda will not trace any request.
	Tracing_DISABLED Tracing = "DISABLED"
)

