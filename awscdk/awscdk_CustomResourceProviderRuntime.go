// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The lambda runtime to use for the resource provider.
//
// This also indicates
// which language is used for the handler.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   provider := awscdk.CustomResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
//   	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	runtime: awscdk.CustomResourceProviderRuntime_NODEJS_14_X,
//   })
//   provider.addToRolePolicy(map[string]*string{
//   	"Effect": jsii.String("Allow"),
//   	"Action": jsii.String("s3:GetObject"),
//   	"Resource": jsii.String("*"),
//   })
//
type CustomResourceProviderRuntime string

const (
	// Node.js 12.x.
	CustomResourceProviderRuntime_NODEJS_12_X CustomResourceProviderRuntime = "NODEJS_12_X"
	// Node.js 14.x.
	CustomResourceProviderRuntime_NODEJS_14_X CustomResourceProviderRuntime = "NODEJS_14_X"
	// Node.js 16.x.
	CustomResourceProviderRuntime_NODEJS_16_X CustomResourceProviderRuntime = "NODEJS_16_X"
)

