package awssagemaker


// The infrastructure configuration for deploying the model to a real-time inference endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   realTimeInferenceConfigProperty := &RealTimeInferenceConfigProperty{
//   	InstanceCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   }
//
type CfnInferenceExperiment_RealTimeInferenceConfigProperty struct {
	// The number of instances of the type specified by `InstanceType` .
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The instance type the model is deployed to.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

