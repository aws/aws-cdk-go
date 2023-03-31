// An experiment to bundle the entire CDK into a single module
package awscdk


// The lambda runtime to use for the resource provider.
//
// This also indicates
// which language is used for the handler.
//
// Example:
//   provider := awscdk.CustomResourceProvider.getOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &customResourceProviderProps{
//   	codeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	runtime: awscdk.CustomResourceProviderRuntime_NODEJS_14_X,
//   	policyStatements: []interface{}{
//   		map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("s3:PutObject*"),
//   			"Resource": jsii.String("*"),
//   		},
//   	},
//   })
//
// Experimental.
type CustomResourceProviderRuntime string

const (
	// Node.js 12.x.
	// Deprecated: Use {@link NODEJS_14_X}.
	CustomResourceProviderRuntime_NODEJS_12 CustomResourceProviderRuntime = "NODEJS_12"
	// Node.js 12.x.
	// Experimental.
	CustomResourceProviderRuntime_NODEJS_12_X CustomResourceProviderRuntime = "NODEJS_12_X"
	// Node.js 14.x.
	// Experimental.
	CustomResourceProviderRuntime_NODEJS_14_X CustomResourceProviderRuntime = "NODEJS_14_X"
	// Node.js 16.x.
	// Experimental.
	CustomResourceProviderRuntime_NODEJS_16_X CustomResourceProviderRuntime = "NODEJS_16_X"
)

