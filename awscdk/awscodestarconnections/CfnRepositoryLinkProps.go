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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html
//
type CfnRepositoryLinkProps struct {
	// The Amazon Resource Name (ARN) of the CodeStarConnection.
	//
	// The ARN is used as the connection reference when the connection is shared between AWS services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-connectionarn
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// the ID of the entity that owns the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-ownerid
	//
	OwnerId *string `field:"required" json:"ownerId" yaml:"ownerId"`
	// The repository for which the link is being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-repositoryname
	//
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties.
	//
	// If not specified, a default key will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Specifies the tags applied to a RepositoryLink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

