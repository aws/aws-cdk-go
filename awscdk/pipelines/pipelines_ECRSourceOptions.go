package pipelines


// Options for ECR sources.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var repository iRepository
//
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
//
type ECRSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The image tag that will be checked for changes.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

