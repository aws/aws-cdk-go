package interfacesawscodestar


// A reference to a GitHubRepository resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitHubRepositoryReference := &GitHubRepositoryReference{
//   	GitHubRepositoryId: jsii.String("gitHubRepositoryId"),
//   }
//
type GitHubRepositoryReference struct {
	// The Id of the GitHubRepository resource.
	GitHubRepositoryId *string `field:"required" json:"gitHubRepositoryId" yaml:"gitHubRepositoryId"`
}

