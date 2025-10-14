package awsecr


// The tag mutability setting for your repository.
//
// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	ImageTagMutability: ecr.TagMutability_IMMUTABLE,
//   })
//
type TagMutability string

const (
	// allow image tags to be overwritten.
	TagMutability_MUTABLE TagMutability = "MUTABLE"
	// all image tags within the repository will be immutable which will prevent them from being overwritten.
	TagMutability_IMMUTABLE TagMutability = "IMMUTABLE"
	// all image tags within the repository will be immutable, while allowing you to define some filters for tags that can be changed.
	TagMutability_IMMUTABLE_WITH_EXCLUSION TagMutability = "IMMUTABLE_WITH_EXCLUSION"
	// allow image tags to be overwritten while allowing you to define some filters for tags that should remain unchanged.
	TagMutability_MUTABLE_WITH_EXCLUSION TagMutability = "MUTABLE_WITH_EXCLUSION"
)

