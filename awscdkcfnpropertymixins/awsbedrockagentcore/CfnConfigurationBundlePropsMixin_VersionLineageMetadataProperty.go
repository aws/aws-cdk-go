package awsbedrockagentcore


// The version lineage metadata that tracks parent versions and creation source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   versionLineageMetadataProperty := &VersionLineageMetadataProperty{
//   	BranchName: jsii.String("branchName"),
//   	CommitMessage: jsii.String("commitMessage"),
//   	CreatedBy: &VersionCreatedBySourceProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//   	},
//   	ParentVersionIds: []*string{
//   		jsii.String("parentVersionIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versionlineagemetadata.html
//
type CfnConfigurationBundlePropsMixin_VersionLineageMetadataProperty struct {
	// The branch name for this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versionlineagemetadata.html#cfn-bedrockagentcore-configurationbundle-versionlineagemetadata-branchname
	//
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// A commit message describing the changes in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versionlineagemetadata.html#cfn-bedrockagentcore-configurationbundle-versionlineagemetadata-commitmessage
	//
	CommitMessage *string `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// The source that created a configuration bundle version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versionlineagemetadata.html#cfn-bedrockagentcore-configurationbundle-versionlineagemetadata-createdby
	//
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// A list of parent version identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versionlineagemetadata.html#cfn-bedrockagentcore-configurationbundle-versionlineagemetadata-parentversionids
	//
	ParentVersionIds *[]*string `field:"optional" json:"parentVersionIds" yaml:"parentVersionIds"`
}

