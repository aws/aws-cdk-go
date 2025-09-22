package awscdk


// The lambda runtime to use for the resource provider.
//
// This also indicates
// which language is used for the handler.
//
// Example:
//   provider := awscdk.CustomResourceProvider_GetOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &CustomResourceProviderProps{
//   	CodeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
//   })
//   provider.AddToRolePolicy(map[string]*string{
//   	"Effect": jsii.String("Allow"),
//   	"Action": jsii.String("s3:GetObject"),
//   	"Resource": jsii.String("*"),
//   })
//
type CustomResourceProviderRuntime string

const (
	// Node.js 12.x.
	// Deprecated: Use latest version.
	CustomResourceProviderRuntime_NODEJS_12_X CustomResourceProviderRuntime = "NODEJS_12_X"
	// Node.js 14.x.
	// Deprecated: Use latest version.
	CustomResourceProviderRuntime_NODEJS_14_X CustomResourceProviderRuntime = "NODEJS_14_X"
	// Node.js 16.x.
	// Deprecated: Use latest version.
	CustomResourceProviderRuntime_NODEJS_16_X CustomResourceProviderRuntime = "NODEJS_16_X"
	// Node.js 18.x.
	CustomResourceProviderRuntime_NODEJS_18_X CustomResourceProviderRuntime = "NODEJS_18_X"
	// Node.js 20.x.
	CustomResourceProviderRuntime_NODEJS_20_X CustomResourceProviderRuntime = "NODEJS_20_X"
	// Node.js 22.x.
	CustomResourceProviderRuntime_NODEJS_22_X CustomResourceProviderRuntime = "NODEJS_22_X"
)

