package awsec2


// Specifies the Elastic Inference Accelerator for the instance.
//
// `ElasticInferenceAccelerator` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticInferenceAcceleratorProperty := &ElasticInferenceAcceleratorProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Count: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-elasticinferenceaccelerator.html
//
type CfnInstance_ElasticInferenceAcceleratorProperty struct {
	// The type of elastic inference accelerator.
	//
	// The possible values are `eia1.medium` , `eia1.large` , `eia1.xlarge` , `eia2.medium` , `eia2.large` , and `eia2.xlarge` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-elasticinferenceaccelerator.html#cfn-ec2-instance-elasticinferenceaccelerator-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The number of elastic inference accelerators to attach to the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-elasticinferenceaccelerator.html#cfn-ec2-instance-elasticinferenceaccelerator-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

