package awssagemaker


// The configuration of the size measurements of the AMI update.
//
// Using this configuration, you can specify whether SageMaker should update your instance group by an amount or percentage of instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacitySizeConfigProperty := &CapacitySizeConfigProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-capacitysizeconfig.html
//
type CfnCluster_CapacitySizeConfigProperty struct {
	// Specifies whether SageMaker should process the update by amount or percentage of instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-capacitysizeconfig.html#cfn-sagemaker-cluster-capacitysizeconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the amount or percentage of instances SageMaker updates at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-capacitysizeconfig.html#cfn-sagemaker-cluster-capacitysizeconfig-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

