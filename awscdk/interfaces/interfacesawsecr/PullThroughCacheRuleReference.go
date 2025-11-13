package interfacesawsecr


// A reference to a PullThroughCacheRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pullThroughCacheRuleReference := &PullThroughCacheRuleReference{
//   	EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   }
//
type PullThroughCacheRuleReference struct {
	// The EcrRepositoryPrefix of the PullThroughCacheRule resource.
	EcrRepositoryPrefix *string `field:"required" json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
}

