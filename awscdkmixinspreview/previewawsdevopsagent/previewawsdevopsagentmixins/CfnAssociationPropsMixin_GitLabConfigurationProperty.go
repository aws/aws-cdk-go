package previewawsdevopsagentmixins


// Configuration for GitLab project integration.
//
// Defines the numeric project ID, full project path (namespace/project-name), GitLab instance identifier, and webhook update settings required for the Agent Space to access and interact with the GitLab project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gitLabConfigurationProperty := &GitLabConfigurationProperty{
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	InstanceIdentifier: jsii.String("instanceIdentifier"),
//   	ProjectId: jsii.String("projectId"),
//   	ProjectPath: jsii.String("projectPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html
//
type CfnAssociationPropsMixin_GitLabConfigurationProperty struct {
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// GitLab instance identifier (e.g., gitlab.com).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-instanceidentifier
	//
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// GitLab numeric project ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-projectid
	//
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Full GitLab project path (e.g., namespace/project-name).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-projectpath
	//
	ProjectPath *string `field:"optional" json:"projectPath" yaml:"projectPath"`
}

