package awsecr

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
//   var repositoryPolicyText interface{}
//
//   cfnRepositoryProps := &cfnRepositoryProps{
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		encryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		kmsKey: jsii.String("kmsKey"),
//   	},
//   	imageScanningConfiguration: &imageScanningConfigurationProperty{
//   		scanOnPush: jsii.Boolean(false),
//   	},
//   	imageTagMutability: jsii.String("imageTagMutability"),
//   	lifecyclePolicy: &lifecyclePolicyProperty{
//   		lifecyclePolicyText: jsii.String("lifecyclePolicyText"),
//   		registryId: jsii.String("registryId"),
//   	},
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
type CfnRepositoryProps struct {
	// The encryption configuration for the repository.
	//
	// This determines how the contents of your repository are encrypted at rest.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The image scanning configuration for the repository.
	//
	// This determines whether images are scanned for known vulnerabilities after being pushed to the repository.
	ImageScanningConfiguration interface{} `field:"optional" json:"imageScanningConfiguration" yaml:"imageScanningConfiguration"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of `MUTABLE` will be used which will allow image tags to be overwritten. If `IMMUTABLE` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// Creates or updates a lifecycle policy.
	//
	// For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) .
	LifecyclePolicy interface{} `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// The name to use for the repository.
	//
	// The repository name may be specified on its own (such as `nginx-web-app` ) or it can be prepended with a namespace to group the repository into a category (such as `project-a/nginx-web-app` ). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// The JSON repository policy text to apply to the repository.
	//
	// For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide* .
	RepositoryPolicyText interface{} `field:"optional" json:"repositoryPolicyText" yaml:"repositoryPolicyText"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

