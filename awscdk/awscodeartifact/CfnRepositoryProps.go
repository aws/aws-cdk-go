package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRepository`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var permissionsPolicyDocument interface{}
//
//   cfnRepositoryProps := &CfnRepositoryProps{
//   	DomainName: jsii.String("domainName"),
//   	RepositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DomainOwner: jsii.String("domainOwner"),
//   	ExternalConnections: []*string{
//   		jsii.String("externalConnections"),
//   	},
//   	PermissionsPolicyDocument: permissionsPolicyDocument,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Upstreams: []*string{
//   		jsii.String("upstreams"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html
//
type CfnRepositoryProps struct {
	// The name of the domain that contains the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of an upstream repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-repositoryname
	//
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// A text description of the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The 12-digit account number of the AWS account that owns the domain that contains the repository.
	//
	// It does not include dashes or spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-domainowner
	//
	DomainOwner *string `field:"optional" json:"domainOwner" yaml:"domainOwner"`
	// An array of external connections associated with the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-externalconnections
	//
	ExternalConnections *[]*string `field:"optional" json:"externalConnections" yaml:"externalConnections"`
	// The document that defines the resource policy that is set on a repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-permissionspolicydocument
	//
	PermissionsPolicyDocument interface{} `field:"optional" json:"permissionsPolicyDocument" yaml:"permissionsPolicyDocument"`
	// A list of tags to be applied to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of upstream repositories to associate with the repository.
	//
	// The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. For more information, see [Working with upstream repositories](https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-repository.html#cfn-codeartifact-repository-upstreams
	//
	Upstreams *[]*string `field:"optional" json:"upstreams" yaml:"upstreams"`
}

