package mixinsawscodestarconnections


// Properties for CfnSyncConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSyncConfigurationMixinProps := &CfnSyncConfigurationMixinProps{
//   	Branch: jsii.String("branch"),
//   	ConfigFile: jsii.String("configFile"),
//   	PublishDeploymentStatus: jsii.String("publishDeploymentStatus"),
//   	RepositoryLinkId: jsii.String("repositoryLinkId"),
//   	ResourceName: jsii.String("resourceName"),
//   	RoleArn: jsii.String("roleArn"),
//   	SyncType: jsii.String("syncType"),
//   	TriggerResourceUpdateOn: jsii.String("triggerResourceUpdateOn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html
//
type CfnSyncConfigurationMixinProps struct {
	// The branch associated with a specific sync configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-branch
	//
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The file path to the configuration file associated with a specific sync configuration.
	//
	// The path should point to an actual file in the sync configurations linked repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-configfile
	//
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	// Whether to enable or disable publishing of deployment status to source providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-publishdeploymentstatus
	//
	PublishDeploymentStatus *string `field:"optional" json:"publishDeploymentStatus" yaml:"publishDeploymentStatus"`
	// The ID of the repository link associated with a specific sync configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-repositorylinkid
	//
	RepositoryLinkId *string `field:"optional" json:"repositoryLinkId" yaml:"repositoryLinkId"`
	// The name of the connection resource associated with a specific sync configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-resourcename
	//
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// The Amazon Resource Name (ARN) of the IAM role associated with a specific sync configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The type of sync for a specific sync configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-synctype
	//
	SyncType *string `field:"optional" json:"syncType" yaml:"syncType"`
	// When to trigger Git sync to begin the stack update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-triggerresourceupdateon
	//
	TriggerResourceUpdateOn *string `field:"optional" json:"triggerResourceUpdateOn" yaml:"triggerResourceUpdateOn"`
}

