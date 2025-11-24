package mixinsawscodebuild


// `ProjectTriggers` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies webhooks that trigger an AWS CodeBuild build.
//
// > The Webhook feature isn't available in AWS CloudFormation for GitHub Enterprise projects. Use the AWS CLI or AWS CodeBuild console to create the webhook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   projectTriggersProperty := &ProjectTriggersProperty{
//   	BuildType: jsii.String("buildType"),
//   	FilterGroups: []interface{}{
//   		[]interface{}{
//   			&WebhookFilterProperty{
//   				ExcludeMatchedPattern: jsii.Boolean(false),
//   				Pattern: jsii.String("pattern"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	PullRequestBuildPolicy: &PullRequestBuildPolicyProperty{
//   		ApproverRoles: []*string{
//   			jsii.String("approverRoles"),
//   		},
//   		RequiresCommentApproval: jsii.String("requiresCommentApproval"),
//   	},
//   	ScopeConfiguration: &ScopeConfigurationProperty{
//   		Domain: jsii.String("domain"),
//   		Name: jsii.String("name"),
//   		Scope: jsii.String("scope"),
//   	},
//   	Webhook: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html
//
type CfnProjectPropsMixin_ProjectTriggersProperty struct {
	// Specifies the type of build this webhook will trigger. Allowed values are:.
	//
	// - **BUILD** - A single build
	// - **BUILD_BATCH** - A batch build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-buildtype
	//
	BuildType *string `field:"optional" json:"buildType" yaml:"buildType"`
	// A list of lists of `WebhookFilter` objects used to determine which webhook events are triggered.
	//
	// At least one `WebhookFilter` in the array must specify `EVENT` as its type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-filtergroups
	//
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-pullrequestbuildpolicy
	//
	PullRequestBuildPolicy interface{} `field:"optional" json:"pullRequestBuildPolicy" yaml:"pullRequestBuildPolicy"`
	// Contains configuration information about the scope for a webhook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"optional" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// Specifies whether or not to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-webhook
	//
	Webhook interface{} `field:"optional" json:"webhook" yaml:"webhook"`
}

