package awslambda


// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("handler.zip"))),
//   	Runtime: lambda.Runtime_JAVA_11(),
//   	Handler: jsii.String("example.Handler::handleRequest"),
//   	RecursiveLoop: lambda.RecursiveLoop_TERMINATE,
//   })
//
type RecursiveLoop string

const (
	// Allows the recursive loop to happen and does not terminate it.
	RecursiveLoop_ALLOW RecursiveLoop = "ALLOW"
	// Terminates the recursive loop.
	RecursiveLoop_TERMINATE RecursiveLoop = "TERMINATE"
)

