package awselasticbeanstalk


// Properties for defining a `CfnConfigurationTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationTemplateProps := &CfnConfigurationTemplateProps{
//   	ApplicationName: jsii.String("applicationName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EnvironmentId: jsii.String("environmentId"),
//   	OptionSettings: []interface{}{
//   		&ConfigurationOptionSettingProperty{
//   			Namespace: jsii.String("namespace"),
//   			OptionName: jsii.String("optionName"),
//
//   			// the properties below are optional
//   			ResourceName: jsii.String("resourceName"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	PlatformArn: jsii.String("platformArn"),
//   	SolutionStackName: jsii.String("solutionStackName"),
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		ApplicationName: jsii.String("applicationName"),
//   		TemplateName: jsii.String("templateName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html
//
type CfnConfigurationTemplateProps struct {
	// The name of the Elastic Beanstalk application to associate with this configuration template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-applicationname
	//
	ApplicationName interface{} `field:"required" json:"applicationName" yaml:"applicationName"`
	// An optional description for this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of an environment whose settings you want to use to create the configuration template.
	//
	// You must specify `EnvironmentId` if you don't specify `PlatformArn` , `SolutionStackName` , or `SourceConfiguration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-environmentid
	//
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// Option values for the Elastic Beanstalk configuration, such as the instance type.
	//
	// If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-optionsettings
	//
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform.
	//
	// For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the *AWS Elastic Beanstalk Developer Guide* .
	//
	// > If you specify `PlatformArn` , then don't specify `SolutionStackName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-platformarn
	//
	PlatformArn *string `field:"optional" json:"platformArn" yaml:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses.
	//
	// For example, `64bit Amazon Linux 2013.09 running Tomcat 7 Java 7` . A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the *AWS Elastic Beanstalk Developer Guide* .
	//
	// You must specify `SolutionStackName` if you don't specify `PlatformArn` , `EnvironmentId` , or `SourceConfiguration` .
	//
	// Use the [`ListAvailableSolutionStacks`](https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ListAvailableSolutionStacks.html) API to obtain a list of available solution stacks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-solutionstackname
	//
	SolutionStackName *string `field:"optional" json:"solutionStackName" yaml:"solutionStackName"`
	// An Elastic Beanstalk configuration template to base this one on.
	//
	// If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.
	//
	// Values specified in `OptionSettings` override any values obtained from the `SourceConfiguration` .
	//
	// You must specify `SourceConfiguration` if you don't specify `PlatformArn` , `EnvironmentId` , or `SolutionStackName` .
	//
	// Constraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

