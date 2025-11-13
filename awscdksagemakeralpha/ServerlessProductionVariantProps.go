package awscdksagemakeralpha


// Construction properties for a serverless production variant.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var model Model
//
//
//   endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("ServerlessEndpointConfig"), &EndpointConfigProps{
//   	ServerlessProductionVariant: &ServerlessProductionVariantProps{
//   		Model: model,
//   		VariantName: jsii.String("serverlessVariant"),
//   		MaxConcurrency: jsii.Number(10),
//   		MemorySizeInMB: jsii.Number(2048),
//   		ProvisionedConcurrency: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type ServerlessProductionVariantProps struct {
	// The maximum number of concurrent invocations your serverless endpoint can process.
	//
	// Valid range: 1-200.
	// Experimental.
	MaxConcurrency *float64 `field:"required" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The memory size of your serverless endpoint.
	//
	// Valid values are in 1 GB increments:
	// 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.
	// Experimental.
	MemorySizeInMB *float64 `field:"required" json:"memorySizeInMB" yaml:"memorySizeInMB"`
	// The model to host.
	// Experimental.
	Model IModel `field:"required" json:"model" yaml:"model"`
	// Name of the production variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// The traffic to a production variant is determined by the ratio of the
	// variant weight to the sum of all variant weight values across all production variants.
	// Default: 1.0
	//
	// Experimental.
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// The number of concurrent invocations that are provisioned and ready to respond to your endpoint.
	//
	// Valid range: 1-200, must be less than or equal to maxConcurrency.
	// Default: - none.
	//
	// Experimental.
	ProvisionedConcurrency *float64 `field:"optional" json:"provisionedConcurrency" yaml:"provisionedConcurrency"`
}

