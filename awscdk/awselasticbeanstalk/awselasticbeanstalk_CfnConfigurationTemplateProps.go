package awselasticbeanstalk


// Properties for defining a `CfnConfigurationTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationTemplateProps := &cfnConfigurationTemplateProps{
//   	applicationName: jsii.String("applicationName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	environmentId: jsii.String("environmentId"),
//   	optionSettings: []interface{}{
//   		&configurationOptionSettingProperty{
//   			namespace: jsii.String("namespace"),
//   			optionName: jsii.String("optionName"),
//
//   			// the properties below are optional
//   			resourceName: jsii.String("resourceName"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	platformArn: jsii.String("platformArn"),
//   	solutionStackName: jsii.String("solutionStackName"),
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		applicationName: jsii.String("applicationName"),
//   		templateName: jsii.String("templateName"),
//   	},
//   }
//
type CfnConfigurationTemplateProps struct {
	// The name of the Elastic Beanstalk application to associate with this configuration template.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// An optional description for this configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of an environment whose settings you want to use to create the configuration template.
	//
	// You must specify `EnvironmentId` if you don't specify `PlatformArn` , `SolutionStackName` , or `SourceConfiguration` .
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// Option values for the Elastic Beanstalk configuration, such as the instance type.
	//
	// If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the *AWS Elastic Beanstalk Developer Guide* .
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform.
	//
	// For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the *AWS Elastic Beanstalk Developer Guide* .
	//
	// > If you specify `PlatformArn` , then don't specify `SolutionStackName` .
	PlatformArn *string `field:"optional" json:"platformArn" yaml:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses.
	//
	// For example, `64bit Amazon Linux 2013.09 running Tomcat 7 Java 7` . A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the *AWS Elastic Beanstalk Developer Guide* .
	//
	// You must specify `SolutionStackName` if you don't specify `PlatformArn` , `EnvironmentId` , or `SourceConfiguration` .
	//
	// Use the [`ListAvailableSolutionStacks`](https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ListAvailableSolutionStacks.html) API to obtain a list of available solution stacks.
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
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

