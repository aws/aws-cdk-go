package awsecr


// Properties for defining a `CfnRepositoryCreationTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRepositoryCreationTemplateProps := &CfnRepositoryCreationTemplateProps{
//   	AppliedFor: []*string{
//   		jsii.String("appliedFor"),
//   	},
//   	Prefix: jsii.String("prefix"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	ImageTagMutability: jsii.String("imageTagMutability"),
//   	LifecyclePolicy: jsii.String("lifecyclePolicy"),
//   	RepositoryPolicy: jsii.String("repositoryPolicy"),
//   	ResourceTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html
//
type CfnRepositoryCreationTemplateProps struct {
	// A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-appliedfor
	//
	AppliedFor *[]*string `field:"required" json:"appliedFor" yaml:"appliedFor"`
	// The prefix use to match the repository name and apply the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-prefix
	//
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// The description of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
	//
	// By default, when no encryption configuration is set or the `AES256` encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.
	//
	// For more control over the encryption of the contents of your repository, you can use server-side encryption with AWS Key Management Service key stored in AWS Key Management Service ( AWS KMS ) to encrypt your images. For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The image tag mutability setting for the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-imagetagmutability
	//
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// The JSON lifecycle policy text to apply to the repository.
	//
	// For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-lifecyclepolicy
	//
	LifecyclePolicy *string `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-repositorypolicy
	//
	RepositoryPolicy *string `field:"optional" json:"repositoryPolicy" yaml:"repositoryPolicy"`
	// The tags attached to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

