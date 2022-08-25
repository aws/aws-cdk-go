package awscodebuild


// The type returned from {@link ISource#bind}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfig := &sourceConfig{
//   	sourceProperty: &sourceProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		auth: &sourceAuthProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			resource: jsii.String("resource"),
//   		},
//   		buildSpec: jsii.String("buildSpec"),
//   		buildStatusConfig: &buildStatusConfigProperty{
//   			context: jsii.String("context"),
//   			targetUrl: jsii.String("targetUrl"),
//   		},
//   		gitCloneDepth: jsii.Number(123),
//   		gitSubmodulesConfig: &gitSubmodulesConfigProperty{
//   			fetchSubmodules: jsii.Boolean(false),
//   		},
//   		insecureSsl: jsii.Boolean(false),
//   		location: jsii.String("location"),
//   		reportBuildStatus: jsii.Boolean(false),
//   		sourceIdentifier: jsii.String("sourceIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	buildTriggers: &projectTriggersProperty{
//   		buildType: jsii.String("buildType"),
//   		filterGroups: []interface{}{
//   			[]interface{}{
//   				&webhookFilterProperty{
//   					pattern: jsii.String("pattern"),
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					excludeMatchedPattern: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		webhook: jsii.Boolean(false),
//   	},
//   	sourceVersion: jsii.String("sourceVersion"),
//   }
//
type SourceConfig struct {
	SourceProperty *CfnProject_SourceProperty `field:"required" json:"sourceProperty" yaml:"sourceProperty"`
	BuildTriggers *CfnProject_ProjectTriggersProperty `field:"optional" json:"buildTriggers" yaml:"buildTriggers"`
	// `AWS::CodeBuild::Project.SourceVersion`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-sourceversion
	//
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

