package previewawsecrmixins


// Properties for CfnPullThroughCacheRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPullThroughCacheRuleMixinProps := &CfnPullThroughCacheRuleMixinProps{
//   	CredentialArn: jsii.String("credentialArn"),
//   	CustomRoleArn: jsii.String("customRoleArn"),
//   	EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	UpstreamRegistry: jsii.String("upstreamRegistry"),
//   	UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   	UpstreamRepositoryPrefix: jsii.String("upstreamRepositoryPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html
//
type CfnPullThroughCacheRuleMixinProps struct {
	// The ARN of the Secrets Manager secret associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-credentialarn
	//
	CredentialArn *string `field:"optional" json:"credentialArn" yaml:"credentialArn"`
	// The ARN of the IAM role associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-customrolearn
	//
	CustomRoleArn *string `field:"optional" json:"customRoleArn" yaml:"customRoleArn"`
	// The Amazon ECR repository prefix associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-ecrrepositoryprefix
	//
	EcrRepositoryPrefix *string `field:"optional" json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
	// The name of the upstream source registry associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-upstreamregistry
	//
	UpstreamRegistry *string `field:"optional" json:"upstreamRegistry" yaml:"upstreamRegistry"`
	// The upstream registry URL associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-upstreamregistryurl
	//
	UpstreamRegistryUrl *string `field:"optional" json:"upstreamRegistryUrl" yaml:"upstreamRegistryUrl"`
	// The upstream repository prefix associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-upstreamrepositoryprefix
	//
	UpstreamRepositoryPrefix *string `field:"optional" json:"upstreamRepositoryPrefix" yaml:"upstreamRepositoryPrefix"`
}

