package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPublicRepository`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repositoryCatalogData interface{}
//   var repositoryPolicyText interface{}
//
//   cfnPublicRepositoryProps := &cfnPublicRepositoryProps{
//   	repositoryCatalogData: repositoryCatalogData,
//   	repositoryName: jsii.String("repositoryName"),
//   	repositoryPolicyText: repositoryPolicyText,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPublicRepositoryProps struct {
	// The details about the repository that are publicly visible in the Amazon ECR Public Gallery.
	//
	// For more information, see [Amazon ECR Public repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the *Amazon ECR Public User Guide* .
	RepositoryCatalogData interface{} `field:"optional" json:"repositoryCatalogData" yaml:"repositoryCatalogData"`
	// The name to use for the public repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the public repository.
	//
	// For more information, see [Amazon ECR Public repository policies](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-policies.html) in the *Amazon ECR Public User Guide* .
	RepositoryPolicyText interface{} `field:"optional" json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

