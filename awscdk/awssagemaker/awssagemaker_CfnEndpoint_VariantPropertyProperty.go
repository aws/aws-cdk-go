package awssagemaker


// Specifies a production variant property type for an Endpoint.
//
// If you are updating an Endpoint with the [RetainAllVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) option set to `true` , the `VarientProperty` objects listed in [ExcludeRetainedVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-ExcludeRetainedVariantProperties) override the existing variant properties of the Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variantPropertyProperty := &variantPropertyProperty{
//   	variantPropertyType: jsii.String("variantPropertyType"),
//   }
//
type CfnEndpoint_VariantPropertyProperty struct {
	// The type of variant property. The supported values are:.
	//
	// - `DesiredInstanceCount` : Overrides the existing variant instance counts using the [InitialInstanceCount](https://docs.aws.amazon.com/sagemaker/latest/dg/API_ProductionVariant.html#SageMaker-Type-ProductionVariant-InitialInstanceCount) values in the [ProductionVariants](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html#SageMaker-CreateEndpointConfig-request-ProductionVariants) .
	// - `DesiredWeight` : Overrides the existing variant weights using the [InitialVariantWeight](https://docs.aws.amazon.com/sagemaker/latest/dg/API_ProductionVariant.html#SageMaker-Type-ProductionVariant-InitialVariantWeight) values in the [ProductionVariants](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html#SageMaker-CreateEndpointConfig-request-ProductionVariants) .
	// - `DataCaptureConfig` : (Not currently supported.)
	VariantPropertyType *string `field:"optional" json:"variantPropertyType" yaml:"variantPropertyType"`
}

