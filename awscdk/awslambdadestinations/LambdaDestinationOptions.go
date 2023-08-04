package awslambdadestinations


// Options for a Lambda destination.
//
// Example:
//   // Auto-extract response payload with a lambda destination
//   var destinationFn function
//
//
//   sourceFn := lambda.NewFunction(this, jsii.String("Source"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	// auto-extract on success
//   	OnSuccess: destinations.NewLambdaDestination(destinationFn, &LambdaDestinationOptions{
//   		ResponseOnly: jsii.Boolean(true),
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
	// Default: false The destination function receives the full invocation record.
	//
	ResponseOnly *bool `field:"optional" json:"responseOnly" yaml:"responseOnly"`
}

