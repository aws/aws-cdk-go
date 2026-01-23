package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for an instance production variant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var acceleratorType AcceleratorType
//   var instanceType InstanceType
//   var model Model
//
//   instanceProductionVariantProps := &InstanceProductionVariantProps{
//   	Model: model,
//   	VariantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	AcceleratorType: acceleratorType,
//   	ContainerStartupHealthCheckTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	InitialInstanceCount: jsii.Number(123),
//   	InitialVariantWeight: jsii.Number(123),
//   	InstanceType: instanceType,
//   }
//
// Experimental.
type InstanceProductionVariantProps struct {
	// The model to host.
	// Experimental.
	Model IModel `field:"required" json:"model" yaml:"model"`
	// Name of the production variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	//
	// EI instances
	// provide on-demand GPU computing for inference.
	// Default: - none.
	//
	// Experimental.
	AcceleratorType AcceleratorType `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The timeout value, in seconds, for your inference container to pass health check.
	//
	// Range between 60 and 3600 seconds.
	// Default: - none.
	//
	// Experimental.
	ContainerStartupHealthCheckTimeout awscdk.Duration `field:"optional" json:"containerStartupHealthCheckTimeout" yaml:"containerStartupHealthCheckTimeout"`
	// Number of instances to launch initially.
	// Default: 1.
	//
	// Experimental.
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	//
	// The traffic to a production variant is determined by the ratio of the
	// variant weight to the sum of all variant weight values across all production variants.
	// Default: 1.0
	//
	// Experimental.
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
	// Instance type of the production variant.
	// Default: InstanceType.T2_MEDIUM
	//
	// Experimental.
	InstanceType InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
}

