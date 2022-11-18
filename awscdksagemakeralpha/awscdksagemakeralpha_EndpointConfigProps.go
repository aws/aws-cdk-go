// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Construction properties for a SageMaker EndpointConfig.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var modelA model
//   var modelB model
//
//
//   endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &endpointConfigProps{
//   	instanceProductionVariants: []instanceProductionVariantProps{
//   		&instanceProductionVariantProps{
//   			model: modelA,
//   			variantName: jsii.String("modelA"),
//   			initialVariantWeight: jsii.Number(2),
//   		},
//   		&instanceProductionVariantProps{
//   			model: modelB,
//   			variantName: jsii.String("variantB"),
//   			initialVariantWeight: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type EndpointConfigProps struct {
	// Optional KMS encryption key associated with this stream.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Name of the endpoint configuration.
	// Experimental.
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// A list of instance production variants.
	//
	// You can always add more variants later by calling
	// {@link EndpointConfig#addInstanceProductionVariant}.
	// Experimental.
	InstanceProductionVariants *[]*InstanceProductionVariantProps `field:"optional" json:"instanceProductionVariants" yaml:"instanceProductionVariants"`
}

