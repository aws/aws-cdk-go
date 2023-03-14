package awselasticbeanstalk


// Use the `ConfigurationOptionSetting` property type to specify an option for an AWS Elastic Beanstalk configuration template when defining an AWS::ElasticBeanstalk::ConfigurationTemplate resource in an AWS CloudFormation template.
//
// The `ConfigurationOptionSetting` property type specifies an option for an AWS Elastic Beanstalk configuration template.
//
// The `OptionSettings` property of the [AWS::ElasticBeanstalk::ConfigurationTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html) resource contains a list of `ConfigurationOptionSetting` property types.
//
// For a list of possible namespaces and option values, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationOptionSettingProperty := &configurationOptionSettingProperty{
//   	namespace: jsii.String("namespace"),
//   	optionName: jsii.String("optionName"),
//
//   	// the properties below are optional
//   	resourceName: jsii.String("resourceName"),
//   	value: jsii.String("value"),
//   }
//
type CfnConfigurationTemplate_ConfigurationOptionSettingProperty struct {
	// A unique namespace that identifies the option's associated AWS resource.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The name of the configuration option.
	OptionName *string `field:"required" json:"optionName" yaml:"optionName"`
	// A unique resource name for the option setting.
	//
	// Use it for a time–based scaling configuration option.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// The current value for the configuration option.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

