package awscodecommit

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
//   cfnRepositoryProps := &cfnRepositoryProps{
//   	repositoryName: jsii.String("repositoryName"),
//
//   	// the properties below are optional
//   	code: &codeProperty{
//   		s3: &s3Property{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//
//   		// the properties below are optional
//   		branchName: jsii.String("branchName"),
//   	},
//   	repositoryDescription: jsii.String("repositoryDescription"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggers: []interface{}{
//   		&repositoryTriggerProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			events: []*string{
//   				jsii.String("events"),
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			branches: []*string{
//   				jsii.String("branches"),
//   			},
//   			customData: jsii.String("customData"),
//   		},
//   	},
//   }
//
type CfnRepositoryProps struct {
	// The name of the new repository to be created.
	//
	// > The repository name must be unique across the calling AWS account . Repository names are limited to 100 alphanumeric, dash, and underscore characters, and cannot include certain characters. For more information about the limits on repository names, see [Quotas](https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html) in the *AWS CodeCommit User Guide* . The suffix .git is prohibited.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	//
	// Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation.
	//
	// > You can only use this property to add code when creating a repository with a AWS CloudFormation template at creation time. This property cannot be used for updating code to an existing repository.
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// A comment or description about the new repository.
	//
	// > The description field for a repository accepts all HTML characters and all valid Unicode characters. Applications that do not HTML-encode the description and display it in a webpage can expose users to potentially malicious code. Make sure that you HTML-encode the description field in any application that uses this API to display the repository description on a webpage.
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// One or more tag key-value pairs to use when tagging this repository.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The JSON block of configuration information for each trigger.
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
}

