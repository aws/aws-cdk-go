package awsecr

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
//   var repositoryPolicyText interface{}
//
//   cfnRepositoryProps := &CfnRepositoryProps{
//   	EmptyOnDelete: jsii.Boolean(false),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	ImageScanningConfiguration: &ImageScanningConfigurationProperty{
//   		ScanOnPush: jsii.Boolean(false),
//   	},
//   	ImageTagMutability: jsii.String("imageTagMutability"),
//   	LifecyclePolicy: &LifecyclePolicyProperty{
//   		LifecyclePolicyText: jsii.String("lifecyclePolicyText"),
//   		RegistryId: jsii.String("registryId"),
//   	},
//   	RepositoryName: jsii.String("repositoryName"),
//   	RepositoryPolicyText: repositoryPolicyText,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html
//
type CfnRepositoryProps struct {
	// If true, deleting the repository force deletes the contents of the repository.
	//
	// If false, the repository must be empty before attempting to delete it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-emptyondelete
	//
	EmptyOnDelete interface{} `field:"optional" json:"emptyOnDelete" yaml:"emptyOnDelete"`
	// The encryption configuration for the repository.
	//
	// This determines how the contents of your repository are encrypted at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The image scanning configuration for the repository.
	//
	// This determines whether images are scanned for known vulnerabilities after being pushed to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-imagescanningconfiguration
	//
	ImageScanningConfiguration interface{} `field:"optional" json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of `MUTABLE` will be used which will allow image tags to be overwritten. If `IMMUTABLE` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-imagetagmutability
	//
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// Creates or updates a lifecycle policy.
	//
	// For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-lifecyclepolicy
	//
	LifecyclePolicy interface{} `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// The name to use for the repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext
	//
	RepositoryPolicyText interface{} `field:"optional" json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

