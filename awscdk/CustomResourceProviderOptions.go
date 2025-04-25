package awscdk


// Initialization options for custom resource providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyStatements interface{}
//   var size size
//
//   customResourceProviderOptions := &CustomResourceProviderOptions{
//   	Description: jsii.String("description"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	MemorySize: size,
//   	PolicyStatements: []interface{}{
//   		policyStatements,
//   	},
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	UseCfnResponseWrapper: jsii.Boolean(false),
//   }
//
type CustomResourceProviderOptions struct {
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
	//   provider := awscdk.CustomResourceProvider_GetOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &CustomResourceProviderProps{
	//   	CodeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
	//   	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
	//   	PolicyStatements: []interface{}{
	//   		map[string]*string{
	//   			"Effect": jsii.String("Allow"),
	//   			"Action": jsii.String("s3:PutObject*"),
	//   			"Resource": jsii.String("*"),
	//   		},
	//   	},
	//   })
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
}

