package appstagingsynthesizeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Default Staging Stack Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bootstrapRole bootstrapRole
//   var permissionsBoundary permissionsBoundary
//   var propertyInjector iPropertyInjector
//   var stackSynthesizer stackSynthesizer
//
//   defaultStagingStackProps := &DefaultStagingStackProps{
//   	AppId: jsii.String("appId"),
//   	Qualifier: jsii.String("qualifier"),
//   	StagingBucketEncryption: awscdk.Aws_s3.BucketEncryption_UNENCRYPTED,
//
//   	// the properties below are optional
//   	AnalyticsReporting: jsii.Boolean(false),
//   	AutoDeleteStagingAssets: jsii.Boolean(false),
//   	CrossRegionReferences: jsii.Boolean(false),
//   	DeployRoleArn: jsii.String("deployRoleArn"),
//   	DeployTimeFileAssetLifetime: cdk.Duration_Minutes(jsii.Number(30)),
//   	Description: jsii.String("description"),
//   	Env: &Environment{
//   		Account: jsii.String("account"),
//   		Region: jsii.String("region"),
//   	},
//   	FileAssetPublishingRole: bootstrapRole,
//   	ImageAssetPublishingRole: bootstrapRole,
//   	ImageAssetVersionCount: jsii.Number(123),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	PermissionsBoundary: permissionsBoundary,
//   	PropertyInjectors: []*iPropertyInjector{
//   		propertyInjector,
//   	},
//   	StackName: jsii.String("stackName"),
//   	StagingBucketName: jsii.String("stagingBucketName"),
//   	StagingStackNamePrefix: jsii.String("stagingStackNamePrefix"),
//   	SuppressTemplateIndentation: jsii.Boolean(false),
//   	Synthesizer: stackSynthesizer,
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TerminationProtection: jsii.Boolean(false),
//   }
//
// Experimental.
type DefaultStagingStackProps struct {
	// A unique identifier for the application that the staging stack belongs to.
	//
	// This identifier will be used in the name of staging resources
	// created for this application, and should be unique across CDK apps.
	//
	// The identifier should include lowercase characters and dashes ('-') only
	// and have a maximum of 20 characters.
	// Experimental.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Encryption type for staging bucket.
	//
	// In future versions of this package, the default will be BucketEncryption.S3_MANAGED.
	//
	// In previous versions of this package, the default was to use KMS encryption for the staging bucket. KMS keys cost
	// $1/month, which could result in unexpected costs for users who are not aware of this. As we stabilize this module
	// we intend to make the default S3-managed encryption, which is free. However, the migration path from KMS to S3
	// managed encryption for existing buckets is not straightforward. Therefore, for now, this property is required.
	//
	// If you have an existing staging bucket encrypted with a KMS key, you will likely want to set this property to
	// BucketEncryption.KMS. If you are creating a new staging bucket, you can set this property to
	// BucketEncryption.S3_MANAGED to avoid the cost of a KMS key.
	// Experimental.
	StagingBucketEncryption awss3.BucketEncryption `field:"required" json:"stagingBucketEncryption" yaml:"stagingBucketEncryption"`
	// Auto deletes objects in the staging S3 bucket and images in the staging ECR repositories.
	// Default: true.
	//
	// Experimental.
	AutoDeleteStagingAssets *bool `field:"optional" json:"autoDeleteStagingAssets" yaml:"autoDeleteStagingAssets"`
	// The lifetime for deploy time file assets.
	//
	// Assets that are only necessary at deployment time (for instance,
	// CloudFormation templates and Lambda source code bundles) will be
	// automatically deleted after this many days. Assets that may be
	// read from the staging bucket during your application's run time
	// will not be deleted.
	//
	// Set this to the length of time you wish to be able to roll back to
	// previous versions of your application without having to do a new
	// `cdk synth` and re-upload of assets.
	// Default: - Duration.days(30)
	//
	// Experimental.
	DeployTimeFileAssetLifetime awscdk.Duration `field:"optional" json:"deployTimeFileAssetLifetime" yaml:"deployTimeFileAssetLifetime"`
	// Pass in an existing role to be used as the file publishing role.
	// Default: - a new role will be created.
	//
	// Experimental.
	FileAssetPublishingRole BootstrapRole `field:"optional" json:"fileAssetPublishingRole" yaml:"fileAssetPublishingRole"`
	// Pass in an existing role to be used as the image publishing role.
	// Default: - a new role will be created.
	//
	// Experimental.
	ImageAssetPublishingRole BootstrapRole `field:"optional" json:"imageAssetPublishingRole" yaml:"imageAssetPublishingRole"`
	// The maximum number of image versions to store in a repository.
	//
	// Previous versions of an image can be stored for rollback purposes.
	// Once a repository has more than 3 image versions stored, the oldest
	// version will be discarded. This allows for sensible garbage collection
	// while maintaining a few previous versions for rollback scenarios.
	// Default: - up to 3 versions stored.
	//
	// Experimental.
	ImageAssetVersionCount *float64 `field:"optional" json:"imageAssetVersionCount" yaml:"imageAssetVersionCount"`
	// Explicit name for the staging bucket.
	// Default: - a well-known name unique to this app/env.
	//
	// Experimental.
	StagingBucketName *string `field:"optional" json:"stagingBucketName" yaml:"stagingBucketName"`
	// Specify a custom prefix to be used as the staging stack name and construct ID.
	//
	// The prefix will be appended before the appId, which
	// is required to be part of the stack name and construct ID to
	// ensure uniqueness.
	// Default: 'StagingStack'.
	//
	// Experimental.
	StagingStackNamePrefix *string `field:"optional" json:"stagingStackNamePrefix" yaml:"stagingStackNamePrefix"`
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
	// Default: - no notfication arns.
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
	// Stack tags that will be applied to all the taggable resources and the stack itself.
	// Default: {}.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	// Default: false.
	//
	// Experimental.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// The qualifier used to specialize strings.
	//
	// Shouldn't be necessary but who knows what people might do.
	// Experimental.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// The ARN of the deploy action role, if given.
	//
	// This role will need permissions to read from to the staging resources.
	// Default: - The CLI credentials are assumed, no additional permissions are granted.
	//
	// Experimental.
	DeployRoleArn *string `field:"optional" json:"deployRoleArn" yaml:"deployRoleArn"`
}

