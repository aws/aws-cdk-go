package awscodestarconnections

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRepositoryLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRepositoryLinkProps := &CfnRepositoryLinkProps{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	OwnerId: jsii.String("ownerId"),
//   	RepositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html
//
type CfnRepositoryLinkProps struct {
	// The Amazon Resource Name (ARN) of the connection associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-connectionarn
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// The owner ID for the repository associated with the repository link, such as the owner ID in GitHub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-ownerid
	//
	OwnerId *string `field:"required" json:"ownerId" yaml:"ownerId"`
	// The name of the repository associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-repositoryname
	//
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The Amazon Resource Name (ARN) of the encryption key for the repository associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The tags for the repository to be associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

