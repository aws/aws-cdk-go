package awscdkservicecatalogappregistryalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties used to define targetapplication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import servicecatalogappregistry_alpha "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var permissionsBoundary permissionsBoundary
//   var propertyInjector iPropertyInjector
//   var stackSynthesizer stackSynthesizer
//
//   targetApplicationCommonOptions := &TargetApplicationCommonOptions{
//   	AnalyticsReporting: jsii.Boolean(false),
//   	AssociateCrossAccountStacks: jsii.Boolean(false),
//   	CrossRegionReferences: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	Env: &Environment{
//   		Account: jsii.String("account"),
//   		Region: jsii.String("region"),
//   	},
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	PermissionsBoundary: permissionsBoundary,
//   	PropertyInjectors: []*iPropertyInjector{
//   		propertyInjector,
//   	},
//   	StackId: jsii.String("stackId"),
//   	StackName: jsii.String("stackName"),
//   	SuppressTemplateIndentation: jsii.Boolean(false),
//   	Synthesizer: stackSynthesizer,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TerminationProtection: jsii.Boolean(false),
//   }
//
// Experimental.
type TargetApplicationCommonOptions struct {
	// Include runtime versioning information in this Stack.
	// Default: `analyticsReporting` setting of containing `App`, or value of
	// 'aws:cdk:version-reporting' context key.
	//
	// Experimental.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// Enable this flag to allow native cross region stack references.
	//
	// Enabling this will create a CloudFormation custom resource
	// in both the producing stack and consuming stack in order to perform the export/import
	//
	// This feature is currently experimental.
	// Default: false.
	//
	// Experimental.
	CrossRegionReferences *bool `field:"optional" json:"crossRegionReferences" yaml:"crossRegionReferences"`
	// A description of the stack.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS environment (account/region) where this stack will be deployed.
	//
	// Set the `region`/`account` fields of `env` to either a concrete value to
	// select the indicated environment (recommended for production stacks), or to
	// the values of environment variables
	// `CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
	// depend on the AWS credentials/configuration that the CDK CLI is executed
	// under (recommended for development stacks).
	//
	// If the `Stack` is instantiated inside a `Stage`, any undefined
	// `region`/`account` fields from `env` will default to the same field on the
	// encompassing `Stage`, if configured there.
	//
	// If either `region` or `account` are not set nor inherited from `Stage`, the
	// Stack will be considered "*environment-agnostic*"". Environment-agnostic
	// stacks can be deployed to any environment but may not be able to take
	// advantage of all features of the CDK. For example, they will not be able to
	// use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
	// automatically translate Service Principals to the right format based on the
	// environment's AWS partition, and other such enhancements.
	//
	// Example:
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   new Stack(app, 'Stack1', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     },
	//   });
	//
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   new Stack(app, 'Stack2', {
	//     env: {
	//       account: process.env.CDK_DEFAULT_ACCOUNT,
	//       region: process.env.CDK_DEFAULT_REGION
	//     },
	//   });
	//
	//   // Define multiple stacks stage associated with an environment
	//   const myStage = new Stage(app, 'MyStage', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     }
	//   });
	//
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   new MyStack(myStage, 'Stack1');
	//   new YourStack(myStage, 'Stack2');
	//
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   new MyStack(app, 'Stack1');
	//
	// Default: - The environment of the containing `Stage` if available,
	// otherwise create the stack will be environment-agnostic.
	//
	// Experimental.
	Env *awscdk.Environment `field:"optional" json:"env" yaml:"env"`
	// SNS Topic ARNs that will receive stack events.
	// Default: - no notification arns.
	//
	// Experimental.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Options for applying a permissions boundary to all IAM Roles and Users created within this Stage.
	// Default: - no permissions boundary is applied.
	//
	// Experimental.
	PermissionsBoundary awscdk.PermissionsBoundary `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// A list of IPropertyInjector attached to this Stack.
	// Default: - no PropertyInjectors.
	//
	// Experimental.
	PropertyInjectors *[]awscdk.IPropertyInjector `field:"optional" json:"propertyInjectors" yaml:"propertyInjectors"`
	// Name to deploy the stack with.
	// Default: - Derived from construct path.
	//
	// Experimental.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Enable this flag to suppress indentation in generated CloudFormation templates.
	//
	// If not specified, the value of the `@aws-cdk/core:suppressTemplateIndentation`
	// context key will be used. If that is not specified, then the
	// default value `false` will be used.
	// Default: - the value of `@aws-cdk/core:suppressTemplateIndentation`, or `false` if that is not set.
	//
	// Experimental.
	SuppressTemplateIndentation *bool `field:"optional" json:"suppressTemplateIndentation" yaml:"suppressTemplateIndentation"`
	// Synthesis method to use while deploying this stack.
	//
	// The Stack Synthesizer controls aspects of synthesis and deployment,
	// like how assets are referenced and what IAM roles to use. For more
	// information, see the README of the main CDK package.
	//
	// If not specified, the `defaultStackSynthesizer` from `App` will be used.
	// If that is not specified, `DefaultStackSynthesizer` is used if
	// `@aws-cdk/core:newStyleStackSynthesis` is set to `true` or the CDK major
	// version is v2. In CDK v1 `LegacyStackSynthesizer` is the default if no
	// other synthesizer is specified.
	// Default: - The synthesizer specified on `App`, or `DefaultStackSynthesizer` otherwise.
	//
	// Experimental.
	Synthesizer awscdk.IStackSynthesizer `field:"optional" json:"synthesizer" yaml:"synthesizer"`
	// Tags that will be applied to the Stack.
	//
	// These tags are applied to the CloudFormation Stack itself. They will not
	// appear in the CloudFormation template.
	//
	// However, at deployment time, CloudFormation will apply these tags to all
	// resources in the stack that support tagging. You will not be able to exempt
	// resources from tagging (using the `excludeResourceTypes` property of
	// `Tags.of(...).add()`) for tags applied in this way.
	// Default: {}.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	// Default: false.
	//
	// Experimental.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// Determines whether any cross-account stacks defined in the CDK app definition should be associated with the target application.
	//
	// If set to `true`, the application will first be shared with the accounts that own the stacks.
	// Default: - false.
	//
	// Experimental.
	AssociateCrossAccountStacks *bool `field:"optional" json:"associateCrossAccountStacks" yaml:"associateCrossAccountStacks"`
	// Stack ID in which application will be created or imported.
	//
	// The id of a stack is also the identifier that you use to
	// refer to it in the [AWS CDK Toolkit](https://docs.aws.amazon.com/cdk/v2/guide/cli.html).
	// Default: - The value of `stackName` will be used as stack id.
	//
	// Deprecated: - Use `stackName` instead to control the name and id of the stack.
	StackId *string `field:"optional" json:"stackId" yaml:"stackId"`
}

