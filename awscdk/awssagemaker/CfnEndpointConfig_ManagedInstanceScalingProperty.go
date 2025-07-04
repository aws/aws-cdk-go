package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstanceScalingProperty := &ManagedInstanceScalingProperty{
//   	MaxInstanceCount: jsii.Number(123),
//   	MinInstanceCount: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html
//
type CfnEndpointConfig_ManagedInstanceScalingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-maxinstancecount
	//
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-mininstancecount
	//
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-managedinstancescaling.html#cfn-sagemaker-endpointconfig-managedinstancescaling-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

