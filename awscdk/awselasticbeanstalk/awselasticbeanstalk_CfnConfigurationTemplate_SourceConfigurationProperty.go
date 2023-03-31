package awselasticbeanstalk


// Use the `SourceConfiguration` property type to specify another AWS Elastic Beanstalk configuration template as the base to creating a new AWS::ElasticBeanstalk::ConfigurationTemplate resource in an AWS CloudFormation template.
//
// An AWS Elastic Beanstalk configuration template to base a new one on. You can use it to define a [AWS::ElasticBeanstalk::ConfigurationTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigurationProperty := &sourceConfigurationProperty{
//   	applicationName: jsii.String("applicationName"),
//   	templateName: jsii.String("templateName"),
//   }
//
type CfnConfigurationTemplate_SourceConfigurationProperty struct {
	// The name of the application associated with the configuration.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The name of the configuration template.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}

