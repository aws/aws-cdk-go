package awscodestarconnections


// Properties for defining a `CfnSyncConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSyncConfigurationProps := &CfnSyncConfigurationProps{
//   	Branch: jsii.String("branch"),
//   	ConfigFile: jsii.String("configFile"),
//   	RepositoryLinkId: jsii.String("repositoryLinkId"),
//   	ResourceName: jsii.String("resourceName"),
//   	RoleArn: jsii.String("roleArn"),
//   	SyncType: jsii.String("syncType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html
//
type CfnSyncConfigurationProps struct {
	// The name of the branch of the repository from which resources are to be synchronized,.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-branch
	//
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The source provider repository path of the sync configuration file of the respective SyncType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-configfile
	//
	ConfigFile *string `field:"required" json:"configFile" yaml:"configFile"`
	// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-repositorylinkid
	//
	RepositoryLinkId *string `field:"required" json:"repositoryLinkId" yaml:"repositoryLinkId"`
	// The name of the resource that is being synchronized to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-resourcename
	//
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html#cfn-codestarconnections-syncconfiguration-synctype
	//
	SyncType *string `field:"required" json:"syncType" yaml:"syncType"`
}

