// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Initialization properties for `CustomResourceProvider`.
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
type CustomResourceProviderProps struct {
	// A local file system directory with the provider's code.
	//
	// The code will be
	// bundled into a zip asset and wired to the provider's AWS Lambda function.
	CodeDirectory *string `field:"required" json:"codeDirectory" yaml:"codeDirectory"`
	// The AWS Lambda runtime and version to use for the provider.
	Runtime CustomResourceProviderRuntime `field:"required" json:"runtime" yaml:"runtime"`
	// A description of the function.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs that are passed to Lambda as Environment.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The amount of memory that your function has access to.
	//
	// Increasing the
	// function's memory also increases its CPU allocation.
	MemorySize Size `field:"optional" json:"memorySize" yaml:"memorySize"`
	// A set of IAM policy statements to include in the inline policy of the provider's lambda function.
	//
	// **Please note**: these are direct IAM JSON policy blobs, *not* `iam.PolicyStatement`
	// objects like you will see in the rest of the CDK.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
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
	PolicyStatements *[]interface{} `field:"optional" json:"policyStatements" yaml:"policyStatements"`
	// AWS Lambda timeout for the provider.
	Timeout Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

