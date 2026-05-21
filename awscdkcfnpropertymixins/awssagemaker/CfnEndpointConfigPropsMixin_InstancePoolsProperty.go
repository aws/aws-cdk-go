package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   instancePoolsProperty := &InstancePoolsProperty{
//   	InstanceType: jsii.String("instanceType"),
//   	ModelNameOverride: jsii.String("modelNameOverride"),
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html
//
type CfnEndpointConfigPropsMixin_InstancePoolsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html#cfn-sagemaker-endpointconfig-instancepools-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html#cfn-sagemaker-endpointconfig-instancepools-modelnameoverride
	//
	ModelNameOverride *string `field:"optional" json:"modelNameOverride" yaml:"modelNameOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-instancepools.html#cfn-sagemaker-endpointconfig-instancepools-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

