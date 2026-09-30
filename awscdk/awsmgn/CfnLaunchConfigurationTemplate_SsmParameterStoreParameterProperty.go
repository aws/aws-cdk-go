package awsmgn


// An AWS Systems Manager Parameter Store parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmParameterStoreParameterProperty := &SsmParameterStoreParameterProperty{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterType: jsii.String("parameterType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmparameterstoreparameter.html
//
type CfnLaunchConfigurationTemplate_SsmParameterStoreParameterProperty struct {
	// AWS Systems Manager Parameter Store parameter name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmparameterstoreparameter.html#cfn-mgn-launchconfigurationtemplate-ssmparameterstoreparameter-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// AWS Systems Manager Parameter Store parameter type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-ssmparameterstoreparameter.html#cfn-mgn-launchconfigurationtemplate-ssmparameterstoreparameter-parametertype
	//
	ParameterType *string `field:"required" json:"parameterType" yaml:"parameterType"`
}

