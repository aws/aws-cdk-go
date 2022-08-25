package awsecr


// The tag mutability setting for your repository.
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &repositoryProps{
//   	imageTagMutability: ecr.tagMutability_IMMUTABLE,
//   })
//
type TagMutability string

const (
	// allow image tags to be overwritten.
	TagMutability_MUTABLE TagMutability = "MUTABLE"
	// all image tags within the repository will be immutable which will prevent them from being overwritten.
	TagMutability_IMMUTABLE TagMutability = "IMMUTABLE"
)

