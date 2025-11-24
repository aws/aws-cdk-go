package mixinsawscodestarconnections

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRepositoryLinkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRepositoryLinkMixinProps := &CfnRepositoryLinkMixinProps{
//   	ConnectionArn: jsii.String("connectionArn"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	OwnerId: jsii.String("ownerId"),
//   	RepositoryName: jsii.String("repositoryName"),
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
type CfnRepositoryLinkMixinProps struct {
	// The Amazon Resource Name (ARN) of the connection associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-connectionarn
	//
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// The Amazon Resource Name (ARN) of the encryption key for the repository associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The owner ID for the repository associated with the repository link, such as the owner ID in GitHub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-ownerid
	//
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// The name of the repository associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-repositoryname
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The tags for the repository to be associated with the repository link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

