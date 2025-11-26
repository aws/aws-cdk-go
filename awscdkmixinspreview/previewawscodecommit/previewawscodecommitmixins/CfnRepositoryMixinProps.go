package previewawscodecommitmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRepositoryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRepositoryMixinProps := &CfnRepositoryMixinProps{
//   	Code: &CodeProperty{
//   		BranchName: jsii.String("branchName"),
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	RepositoryDescription: jsii.String("repositoryDescription"),
//   	RepositoryName: jsii.String("repositoryName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Triggers: []interface{}{
//   		&RepositoryTriggerProperty{
//   			Branches: []*string{
//   				jsii.String("branches"),
//   			},
//   			CustomData: jsii.String("customData"),
//   			DestinationArn: jsii.String("destinationArn"),
//   			Events: []*string{
//   				jsii.String("events"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html
//
type CfnRepositoryMixinProps struct {
	// Information about code to be committed to a repository after it is created in an AWS CloudFormation stack.
	//
	// Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation.
	//
	// > You can only use this property to add code when creating a repository with a CloudFormation template at creation time. This property cannot be used for updating code to an existing repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-code
	//
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// The ID of the AWS Key Management Service encryption key used to encrypt and decrypt the repository.
	//
	// > The input can be the full ARN, the key ID, or the key alias. For more information, see [Finding the key ID and key ARN](https://docs.aws.amazon.com/kms/latest/developerguide/find-cmk-id-arn.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A comment or description about the new repository.
	//
	// > The description field for a repository accepts all HTML characters and all valid Unicode characters. Applications that do not HTML-encode the description and display it in a webpage can expose users to potentially malicious code. Make sure that you HTML-encode the description field in any application that uses this API to display the repository description on a webpage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-repositorydescription
	//
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// The name of the new repository to be created.
	//
	// > The repository name must be unique across the calling AWS account . Repository names are limited to 100 alphanumeric, dash, and underscore characters, and cannot include certain characters. For more information about the limits on repository names, see [Quotas](https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html) in the *AWS CodeCommit User Guide* . The suffix .git is prohibited.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-repositoryname
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// One or more tag key-value pairs to use when tagging this repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The JSON block of configuration information for each trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codecommit-repository.html#cfn-codecommit-repository-triggers
	//
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
}

