package awscodegurureviewer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRepositoryAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRepositoryAssociationProps := &cfnRepositoryAssociationProps{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	bucketName: jsii.String("bucketName"),
//   	connectionArn: jsii.String("connectionArn"),
//   	owner: jsii.String("owner"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRepositoryAssociationProps struct {
	// The name of the repository.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of repository that contains the source code to be reviewed. The valid values are:.
	//
	// - `CodeCommit`
	// - `Bitbucket`
	// - `GitHubEnterpriseServer`
	// - `S3Bucket`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the bucket.
	//
	// This is required for your S3Bucket repositoryThe name must start with the prefix, `codeguru-reviewer-*` .
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
	//
	// Its format is `arn:aws:codestar-connections:region-id:aws-account_id:connection/connection-id` . For more information, see [Connection](https://docs.aws.amazon.com/codestar-connections/latest/APIReference/API_Connection.html) in the *AWS CodeStar Connections API Reference* .
	//
	// `ConnectionArn` must be specified for Bitbucket and GitHub Enterprise Server repositories. It has no effect if it is specified for an AWS CodeCommit repository.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// The owner of the repository.
	//
	// For a GitHub Enterprise Server or Bitbucket repository, this is the username for the account that owns the repository.
	//
	// `Owner` must be specified for Bitbucket and GitHub Enterprise Server repositories. It has no effect if it is specified for an AWS CodeCommit repository.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// An array of key-value pairs used to tag an associated repository.
	//
	// A tag is a custom attribute label with two parts:
	//
	// - A *tag key* (for example, `CostCenter` , `Environment` , `Project` , or `Secret` ). Tag keys are case sensitive.
	// - An optional field known as a *tag value* (for example, `111122223333` , `Production` , or a team name). Omitting the tag value is the same as using an empty string. Like tag keys, tag values are case sensitive.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

