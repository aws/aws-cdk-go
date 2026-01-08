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
//   	CustomRoleArn: jsii.String("customRoleArn"),
//   	Description: jsii.String("description"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	ImageTagMutability: jsii.String("imageTagMutability"),
//   	ImageTagMutabilityExclusionFilters: []interface{}{
//   		&ImageTagMutabilityExclusionFilterProperty{
//   			ImageTagMutabilityExclusionFilterType: jsii.String("imageTagMutabilityExclusionFilterType"),
//   			ImageTagMutabilityExclusionFilterValue: jsii.String("imageTagMutabilityExclusionFilterValue"),
//   		},
//   	},
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
	// A list of enumerable Strings representing the repository creation scenarios that this template will apply towards.
	//
	// The supported scenarios are PULL_THROUGH_CACHE, REPLICATION, and CREATE_ON_PUSH.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-appliedfor
	//
	AppliedFor *[]*string `field:"required" json:"appliedFor" yaml:"appliedFor"`
	// The repository namespace prefix associated with the repository creation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-prefix
	//
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// The ARN of the role to be assumed by Amazon ECR.
	//
	// Amazon ECR will assume your supplied role when the customRoleArn is specified. When this field isn't specified, Amazon ECR will use the service-linked role for the repository creation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-customrolearn
	//
	CustomRoleArn *string `field:"optional" json:"customRoleArn" yaml:"customRoleArn"`
	// The description associated with the repository creation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The encryption configuration associated with the repository creation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of `MUTABLE` will be used which will allow image tags to be overwritten. If `IMMUTABLE` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-imagetagmutability
	//
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// A list of filters that specify which image tags are excluded from the repository creation template's image tag mutability setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilters
	//
	ImageTagMutabilityExclusionFilters interface{} `field:"optional" json:"imageTagMutabilityExclusionFilters" yaml:"imageTagMutabilityExclusionFilters"`
	// The lifecycle policy to use for repositories created using the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-lifecyclepolicy
	//
	LifecyclePolicy *string `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// The repository policy to apply to repositories created using the template.
	//
	// A repository policy is a permissions policy associated with a repository to control access permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-repositorypolicy
	//
	RepositoryPolicy *string `field:"optional" json:"repositoryPolicy" yaml:"repositoryPolicy"`
	// The metadata to apply to the repository to help you categorize and organize.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repositorycreationtemplate.html#cfn-ecr-repositorycreationtemplate-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

