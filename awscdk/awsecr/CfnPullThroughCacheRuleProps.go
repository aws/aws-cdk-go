package awsecr


// Properties for defining a `CfnPullThroughCacheRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPullThroughCacheRuleProps := &CfnPullThroughCacheRuleProps{
//   	CredentialArn: jsii.String("credentialArn"),
//   	EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	UpstreamRegistry: jsii.String("upstreamRegistry"),
//   	UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html
//
type CfnPullThroughCacheRuleProps struct {
	// The ARN of the Secrets Manager secret associated with the pull through cache rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html#cfn-ecr-pullthroughcacherule-credentialarn
	//
	CredentialArn *string `field:"optional" json:"credentialArn" yaml:"credentialArn"`
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
}

