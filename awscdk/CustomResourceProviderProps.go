package awscdk


// Initialization properties for `CustomResourceProvider`.
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
type CustomResourceProviderProps struct {
	// A description of the function.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs that are passed to Lambda as Environment.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The amount of memory that your function has access to.
	//
	// Increasing the
	// function's memory also increases its CPU allocation.
	// Default: Size.mebibytes(128)
	//
	MemorySize Size `field:"optional" json:"memorySize" yaml:"memorySize"`
	// A set of IAM policy statements to include in the inline policy of the provider's lambda function.
	//
	// **Please note**: these are direct IAM JSON policy blobs, *not* `iam.PolicyStatement`
	// objects like you will see in the rest of the CDK.
	//
	// Example:
	//   const provider = CustomResourceProvider.getOrCreateProvider(this, 'Custom::MyCustomResourceType', {
	//     codeDirectory: `${__dirname}/my-handler`,
	//     runtime: CustomResourceProviderRuntime.NODEJS_18_X,
	//     policyStatements: [
	//       {
	//         Effect: 'Allow',
	//         Action: 's3:PutObject*',
	//         Resource: '*',
	//       }
	//     ],
	//   });
	//
	// Default: - no additional inline policy.
	//
	PolicyStatements *[]interface{} `field:"optional" json:"policyStatements" yaml:"policyStatements"`
	// AWS Lambda timeout for the provider.
	// Default: Duration.minutes(15)
	//
	Timeout Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Whether or not the cloudformation response wrapper (`nodejs-entrypoint.ts`) is used. If set to `true`, `nodejs-entrypoint.js` is bundled in the same asset as the custom resource and set as the entrypoint. If set to `false`, the custom resource provided is the entrypoint.
	// Default: - `true` if `inlineCode: false` and `false` otherwise.
	//
	UseCfnResponseWrapper *bool `field:"optional" json:"useCfnResponseWrapper" yaml:"useCfnResponseWrapper"`
	// A local file system directory with the provider's code.
	//
	// The code will be
	// bundled into a zip asset and wired to the provider's AWS Lambda function.
	CodeDirectory *string `field:"required" json:"codeDirectory" yaml:"codeDirectory"`
	// The AWS Lambda runtime and version to use for the provider.
	Runtime CustomResourceProviderRuntime `field:"required" json:"runtime" yaml:"runtime"`
}

