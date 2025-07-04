package awscodebuild


// Specifies the visibility of the project's builds.
//
// Example:
//   codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	Visibility: codebuild.ProjectVisibility_PUBLIC_READ,
//   })
//
type ProjectVisibility string

const (
	// The project builds are visible to the public.
	ProjectVisibility_PUBLIC_READ ProjectVisibility = "PUBLIC_READ"
	// The project builds are not visible to the public.
	ProjectVisibility_PRIVATE ProjectVisibility = "PRIVATE"
)

