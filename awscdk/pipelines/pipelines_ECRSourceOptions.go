package pipelines


// Options for ECR sources.
//
// Example:
//   var repository iRepository
//
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
//
// Experimental.
type ECRSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Experimental.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The image tag that will be checked for changes.
	// Experimental.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

