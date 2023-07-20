package awsecr


// The `LifecyclePolicy` property type specifies a lifecycle policy.
//
// For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) in the *Amazon ECR User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecyclePolicyProperty := &LifecyclePolicyProperty{
//   	LifecyclePolicyText: jsii.String("lifecyclePolicyText"),
//   	RegistryId: jsii.String("registryId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html
//
type CfnRepository_LifecyclePolicyProperty struct {
	// The JSON repository policy text to apply to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-lifecyclepolicytext
	//
	LifecyclePolicyText *string `field:"optional" json:"lifecyclePolicyText" yaml:"lifecyclePolicyText"`
	// The AWS account ID associated with the registry that contains the repository.
	//
	// If you do not specify a registry, the default registry is assumed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-registryid
	//
	RegistryId *string `field:"optional" json:"registryId" yaml:"registryId"`
}

