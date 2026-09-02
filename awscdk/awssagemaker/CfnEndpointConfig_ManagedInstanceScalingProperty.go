package awssagemaker


// Settings that control the range in the number of instances that the endpoint provisions as it scales up or down to accommodate traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstanceScalingProperty := &ManagedInstanceScalingProperty{
//   	MaxInstanceCount: jsii.Number(123),
//   	MinInstanceCount: jsii.Number(123),
//   	ScaleInPolicy: &ScaleInPolicyProperty{
//   		Strategy: jsii.String("strategy"),
//
//   		// the properties below are optional
//   		CooldownInMinutes: jsii.Number(123),
//   		MaximumStepSize: jsii.Number(123),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html
//
type CfnEndpointConfig_ManagedInstanceScalingProperty struct {
	// The maximum number of instances that the endpoint can provision when it scales up to accommodate an increase in traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-maxinstancecount
	//
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// The minimum number of instances that the endpoint must retain when it scales down to accommodate a decrease in traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-mininstancecount
	//
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// Specifies how the endpoint releases instances when managed instance scaling scales in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-scaleinpolicy
	//
	ScaleInPolicy interface{} `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// Indicates whether managed instance scaling is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

