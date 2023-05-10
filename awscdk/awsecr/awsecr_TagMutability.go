package awsecr


// The tag mutability setting for your repository.
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	ImageTagMutability: ecr.TagMutability_IMMUTABLE,
//   })
//
// Experimental.
type TagMutability string

const (
	// allow image tags to be overwritten.
	// Experimental.
	TagMutability_MUTABLE TagMutability = "MUTABLE"
	// all image tags within the repository will be immutable which will prevent them from being overwritten.
	// Experimental.
	TagMutability_IMMUTABLE TagMutability = "IMMUTABLE"
)

