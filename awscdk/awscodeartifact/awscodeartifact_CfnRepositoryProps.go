package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnRepositoryProps := &cfnRepositoryProps{
//   	domainName: jsii.String("domainName"),
//   	repositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	domainOwner: jsii.String("domainOwner"),
//   	externalConnections: []*string{
//   		jsii.String("externalConnections"),
//   	},
//   	permissionsPolicyDocument: permissionsPolicyDocument,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	upstreams: []*string{
//   		jsii.String("upstreams"),
//   	},
//   }
//
type CfnRepositoryProps struct {
	// The name of the domain that contains the repository.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of an upstream repository.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// A text description of the repository.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The 12-digit account number of the AWS account that owns the domain that contains the repository.
	//
	// It does not include dashes or spaces.
	DomainOwner *string `field:"optional" json:"domainOwner" yaml:"domainOwner"`
	// An array of external connections associated with the repository.
	ExternalConnections *[]*string `field:"optional" json:"externalConnections" yaml:"externalConnections"`
	// The document that defines the resource policy that is set on a repository.
	PermissionsPolicyDocument interface{} `field:"optional" json:"permissionsPolicyDocument" yaml:"permissionsPolicyDocument"`
	// A list of tags to be applied to the repository.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of upstream repositories to associate with the repository.
	//
	// The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. For more information, see [Working with upstream repositories](https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html) .
	Upstreams *[]*string `field:"optional" json:"upstreams" yaml:"upstreams"`
}

