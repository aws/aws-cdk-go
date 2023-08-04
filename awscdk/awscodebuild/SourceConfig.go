package awscodebuild


// The type returned from `ISource#bind`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfig := &SourceConfig{
//   	SourceProperty: &SourceProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Auth: &SourceAuthProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Resource: jsii.String("resource"),
//   		},
//   		BuildSpec: jsii.String("buildSpec"),
//   		BuildStatusConfig: &BuildStatusConfigProperty{
//   			Context: jsii.String("context"),
//   			TargetUrl: jsii.String("targetUrl"),
//   		},
//   		GitCloneDepth: jsii.Number(123),
//   		GitSubmodulesConfig: &GitSubmodulesConfigProperty{
//   			FetchSubmodules: jsii.Boolean(false),
//   		},
//   		InsecureSsl: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   		ReportBuildStatus: jsii.Boolean(false),
//   		SourceIdentifier: jsii.String("sourceIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	BuildTriggers: &ProjectTriggersProperty{
//   		BuildType: jsii.String("buildType"),
//   		FilterGroups: []interface{}{
//   			[]interface{}{
//   				&WebhookFilterProperty{
//   					Pattern: jsii.String("pattern"),
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					ExcludeMatchedPattern: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		Webhook: jsii.Boolean(false),
//   	},
//   	SourceVersion: jsii.String("sourceVersion"),
//   }
//
type SourceConfig struct {
	SourceProperty *CfnProject_SourceProperty `field:"required" json:"sourceProperty" yaml:"sourceProperty"`
	BuildTriggers *CfnProject_ProjectTriggersProperty `field:"optional" json:"buildTriggers" yaml:"buildTriggers"`
	// `AWS::CodeBuild::Project.SourceVersion`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-sourceversion
	//
	// Default: the latest version.
	//
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

