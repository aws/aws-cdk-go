package awssagemaker


// Specifies an instance type and its priority for a heterogeneous endpoint.
//
// Use instance pools to configure a production variant with multiple instance types, enabling the endpoint to provision instances across different types based on priority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instancePoolProperty := &InstancePoolProperty{
//   	InstanceType: jsii.String("instanceType"),
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	ModelNameOverride: jsii.String("modelNameOverride"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepool.html
//
type CfnEndpointConfig_InstancePoolProperty struct {
	// The ML compute instance type for the instance pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepool.html#cfn-sagemaker-endpointconfig-instancepool-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The priority for the instance pool.
	//
	// SageMaker attempts to provision instances in order of priority, starting with the lowest value. If instances for a higher-priority pool are unavailable, SageMaker attempts to provision from the next pool. Valid values: 1 to 5, where 1 is the highest priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepool.html#cfn-sagemaker-endpointconfig-instancepool-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The name of a SageMaker model to use for this instance pool instead of the model specified for the production variant.
	//
	// Use this to deploy a different model optimized for the instance type in this pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepool.html#cfn-sagemaker-endpointconfig-instancepool-modelnameoverride
	//
	ModelNameOverride *string `field:"optional" json:"modelNameOverride" yaml:"modelNameOverride"`
}

