package awscodebuild


// `ProjectTriggers` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies webhooks that trigger an AWS CodeBuild build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectTriggersProperty := &projectTriggersProperty{
//   	buildType: jsii.String("buildType"),
//   	filterGroups: []interface{}{
//   		[]interface{}{
//   			&webhookFilterProperty{
//   				pattern: jsii.String("pattern"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				excludeMatchedPattern: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	webhook: jsii.Boolean(false),
//   }
//
type CfnProject_ProjectTriggersProperty struct {
	// Specifies the type of build this webhook will trigger. Allowed values are:.
	//
	// - **BUILD** - A single build
	// - **BUILD_BATCH** - A batch build.
	BuildType *string `field:"optional" json:"buildType" yaml:"buildType"`
	// A list of lists of `WebhookFilter` objects used to determine which webhook events are triggered.
	//
	// At least one `WebhookFilter` in the array must specify `EVENT` as its type.
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// Specifies whether or not to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	Webhook interface{} `field:"optional" json:"webhook" yaml:"webhook"`
}

