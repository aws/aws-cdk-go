package awsdevopsagent


// GitLab project integration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitLabConfigurationProperty := &GitLabConfigurationProperty{
//   	ProjectId: jsii.String("projectId"),
//   	ProjectPath: jsii.String("projectPath"),
//
//   	// the properties below are optional
//   	EnableWebhookUpdates: jsii.Boolean(false),
//   	InstanceIdentifier: jsii.String("instanceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html
//
type CfnAssociation_GitLabConfigurationProperty struct {
	// GitLab numeric project ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-projectid
	//
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Full GitLab project path (e.g., namespace/project-name).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-projectpath
	//
	ProjectPath *string `field:"required" json:"projectPath" yaml:"projectPath"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-enablewebhookupdates
	//
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// GitLab instance identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-gitlabconfiguration.html#cfn-devopsagent-association-gitlabconfiguration-instanceidentifier
	//
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
}

