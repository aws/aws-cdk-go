package awsecr


// Properties for defining a `CfnPullThroughCacheRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPullThroughCacheRuleProps := &CfnPullThroughCacheRuleProps{
//   	EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   }
//
type CfnPullThroughCacheRuleProps struct {
	// The Amazon ECR repository prefix associated with the pull through cache rule.
	EcrRepositoryPrefix *string `field:"optional" json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
	// The upstream registry URL associated with the pull through cache rule.
	UpstreamRegistryUrl *string `field:"optional" json:"upstreamRegistryUrl" yaml:"upstreamRegistryUrl"`
}

