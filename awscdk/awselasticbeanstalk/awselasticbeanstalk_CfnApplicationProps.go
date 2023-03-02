package awselasticbeanstalk


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   	description: jsii.String("description"),
//   	resourceLifecycleConfig: &applicationResourceLifecycleConfigProperty{
//   		serviceRole: jsii.String("serviceRole"),
//   		versionLifecycleConfig: &applicationVersionLifecycleConfigProperty{
//   			maxAgeRule: &maxAgeRuleProperty{
//   				deleteSourceFromS3: jsii.Boolean(false),
//   				enabled: jsii.Boolean(false),
//   				maxAgeInDays: jsii.Number(123),
//   			},
//   			maxCountRule: &maxCountRuleProperty{
//   				deleteSourceFromS3: jsii.Boolean(false),
//   				enabled: jsii.Boolean(false),
//   				maxCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// A name for the Elastic Beanstalk application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Your description of the application.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	ResourceLifecycleConfig interface{} `field:"optional" json:"resourceLifecycleConfig" yaml:"resourceLifecycleConfig"`
}

