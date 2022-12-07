package awslambdadestinations


// Options for a Lambda destination.
//
// Example:
//   // Auto-extract response payload with a lambda destination
//   var destinationFn function
//
//
//   sourceFn := lambda.NewFunction(this, jsii.String("Source"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	// auto-extract on success
//   	onSuccess: destinations.NewLambdaDestination(destinationFn, &lambdaDestinationOptions{
//   		responseOnly: jsii.Boolean(true),
//   	}),
//   })
//
type LambdaDestinationOptions struct {
	// Whether the destination function receives only the `responsePayload` of the source function.
	//
	// When set to `true` and used as `onSuccess` destination, the destination
	// function will be invoked with the payload returned by the source function.
	//
	// When set to `true` and used as `onFailure` destination, the destination
	// function will be invoked with the error object returned by source function.
	//
	// See the README of this module to see a full explanation of this option.
	ResponseOnly *bool `field:"optional" json:"responseOnly" yaml:"responseOnly"`
}

