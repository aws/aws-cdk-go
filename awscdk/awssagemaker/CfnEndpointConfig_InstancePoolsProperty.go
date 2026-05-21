package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instancePoolsProperty := &InstancePoolsProperty{
//   	InstanceType: jsii.String("instanceType"),
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	ModelNameOverride: jsii.String("modelNameOverride"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html
//
type CfnEndpointConfig_InstancePoolsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html#cfn-sagemaker-endpointconfig-instancepools-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html#cfn-sagemaker-endpointconfig-instancepools-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html#cfn-sagemaker-endpointconfig-instancepools-modelnameoverride
	//
	ModelNameOverride *string `field:"optional" json:"modelNameOverride" yaml:"modelNameOverride"`
}

