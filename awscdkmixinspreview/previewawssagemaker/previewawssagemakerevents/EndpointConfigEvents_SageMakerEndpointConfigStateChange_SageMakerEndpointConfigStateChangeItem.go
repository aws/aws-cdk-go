package previewawssagemakerevents


// Type definition for SageMakerEndpointConfigStateChangeItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerEndpointConfigStateChangeItem := &SageMakerEndpointConfigStateChangeItem{
//   	InitialInstanceCount: []*string{
//   		jsii.String("initialInstanceCount"),
//   	},
//   	InitialVariantWeight: []*string{
//   		jsii.String("initialVariantWeight"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	ModelName: []*string{
//   		jsii.String("modelName"),
//   	},
//   	VariantName: []*string{
//   		jsii.String("variantName"),
//   	},
//   }
//
// Experimental.
type EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeItem struct {
	// InitialInstanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of InitialInstanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitialInstanceCount *[]*string `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// InitialVariantWeight property.
	//
	// Specify an array of string values to match this event if the actual value of InitialVariantWeight is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitialVariantWeight *[]*string `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// InstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// ModelName property.
	//
	// Specify an array of string values to match this event if the actual value of ModelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelName *[]*string `field:"optional" json:"modelName" yaml:"modelName"`
	// VariantName property.
	//
	// Specify an array of string values to match this event if the actual value of VariantName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VariantName *[]*string `field:"optional" json:"variantName" yaml:"variantName"`
}

