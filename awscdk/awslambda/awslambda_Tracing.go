package awslambda


// X-Ray Tracing Modes (https://docs.aws.amazon.com/lambda/latest/dg/API_TracingConfig.html).
//
// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_16_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = function(event, ctx, cb) { return cb(null, \"hi\"); }")),
//   	tracing: lambda.tracing_ACTIVE,
//   })
//
type Tracing string

const (
	// Lambda will respect any tracing header it receives from an upstream service.
	//
	// If no tracing header is received, Lambda will call X-Ray for a tracing decision.
	Tracing_ACTIVE Tracing = "ACTIVE"
	// Lambda will only trace the request from an upstream service if it contains a tracing header with "sampled=1".
	Tracing_PASS_THROUGH Tracing = "PASS_THROUGH"
	// Lambda will not trace any request.
	Tracing_DISABLED Tracing = "DISABLED"
)

