package awselasticbeanstalk


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	Description: jsii.String("description"),
//   	ResourceLifecycleConfig: &ApplicationResourceLifecycleConfigProperty{
//   		ServiceRole: jsii.String("serviceRole"),
//   		VersionLifecycleConfig: &ApplicationVersionLifecycleConfigProperty{
//   			MaxAgeRule: &MaxAgeRuleProperty{
//   				DeleteSourceFromS3: jsii.Boolean(false),
//   				Enabled: jsii.Boolean(false),
//   				MaxAgeInDays: jsii.Number(123),
//   			},
//   			MaxCountRule: &MaxCountRuleProperty{
//   				DeleteSourceFromS3: jsii.Boolean(false),
//   				Enabled: jsii.Boolean(false),
//   				MaxCount: jsii.Number(123),
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

