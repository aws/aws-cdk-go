package previewawselasticbeanstalkmixins


// Use the `SourceConfiguration` property type to specify another AWS Elastic Beanstalk configuration template as the base to creating a new AWS::ElasticBeanstalk::ConfigurationTemplate resource in an AWS CloudFormation template.
//
// An AWS Elastic Beanstalk configuration template to base a new one on. You can use it to define a [AWS::ElasticBeanstalk::ConfigurationTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceConfigurationProperty := &SourceConfigurationProperty{
//   	ApplicationName: jsii.String("applicationName"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html
//
type CfnConfigurationTemplatePropsMixin_SourceConfigurationProperty struct {
	// The name of the application associated with the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The name of the configuration template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

