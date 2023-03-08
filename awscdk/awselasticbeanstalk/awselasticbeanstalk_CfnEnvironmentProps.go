package awselasticbeanstalk

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEnvironmentProps := &cfnEnvironmentProps{
//   	applicationName: jsii.String("applicationName"),
//
//   	// the properties below are optional
//   	cnamePrefix: jsii.String("cnamePrefix"),
//   	description: jsii.String("description"),
//   	environmentName: jsii.String("environmentName"),
//   	operationsRole: jsii.String("operationsRole"),
//   	optionSettings: []interface{}{
//   		&optionSettingProperty{
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
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateName: jsii.String("templateName"),
//   	tier: &tierProperty{
//   		name: jsii.String("name"),
//   		type: jsii.String("type"),
//   		version: jsii.String("version"),
//   	},
//   	versionLabel: jsii.String("versionLabel"),
//   }
//
type CfnEnvironmentProps struct {
	// The name of the application that is associated with this environment.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL.
	//
	// If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
	CnamePrefix *string `field:"optional" json:"cnamePrefix" yaml:"cnamePrefix"`
	// Your description for this environment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A unique name for the environment.
	//
	// Constraint: Must be from 4 to 40 characters in length. The name can contain only letters, numbers, and hyphens. It can't start or end with a hyphen. This name must be unique within a region in your account.
	//
	// If you don't specify the `CNAMEPrefix` parameter, the environment name becomes part of the CNAME, and therefore part of the visible URL for your application.
	//
	// If you don't specify an environment name, AWS CloudFormation generates a unique physical ID and uses that ID for the environment name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// > The operations role feature of AWS Elastic Beanstalk is in beta release and is subject to change.
	//
	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role. If specified, Elastic Beanstalk uses the operations role for permissions to downstream services during this call and during subsequent calls acting on this environment. To specify an operations role, you must have the `iam:PassRole` permission for the role.
	OperationsRole *string `field:"optional" json:"operationsRole" yaml:"operationsRole"`
	// Key-value pairs defining configuration options for this environment, such as the instance type.
	//
	// These options override the values that are defined in the solution stack or the [configuration template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html) . If you remove any options during a stack update, the removed options retain their current values.
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
	//
	// For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the *AWS Elastic Beanstalk Developer Guide* .
	//
	// > If you specify `PlatformArn` , don't specify `SolutionStackName` .
	PlatformArn *string `field:"optional" json:"platformArn" yaml:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
	//
	// If specified, Elastic Beanstalk sets the configuration values to the default values associated with the specified solution stack. For a list of current solution stacks, see [Elastic Beanstalk Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platforms-supported.html) in the *AWS Elastic Beanstalk Platforms* guide.
	//
	// > If you specify `SolutionStackName` , don't specify `PlatformArn` or `TemplateName` .
	SolutionStackName *string `field:"optional" json:"solutionStackName" yaml:"solutionStackName"`
	// Specifies the tags applied to resources in the environment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the Elastic Beanstalk configuration template to use with the environment.
	//
	// > If you specify `TemplateName` , then don't specify `SolutionStackName` .
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// Specifies the tier to use in creating this environment.
	//
	// The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
	Tier interface{} `field:"optional" json:"tier" yaml:"tier"`
	// The name of the application version to deploy.
	//
	// Default: If not specified, Elastic Beanstalk attempts to deploy the sample application.
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

