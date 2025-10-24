package pipelines


// Options for ECR sources.
//
// Example:
//   var repository IRepository
//
//   pipelines.CodePipelineSource_Ecr(repository, &ECRSourceOptions{
//   	ImageTag: jsii.String("latest"),
//   })
//
type ECRSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Default: - The repository name.
	//
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The image tag that will be checked for changes.
	// Default: latest.
	//
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

