package awsimagebuilderalpha


// The rendered Dockerfile value, for use in CloudFormation.
//
// - For inline dockerfiles, dockerfileTemplateData is the Dockerfile template text
// - For S3-backed dockerfiles, dockerfileTemplateUri is the S3 URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   dockerfileTemplateConfig := &DockerfileTemplateConfig{
//   	DockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	DockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   }
//
// Experimental.
type DockerfileTemplateConfig struct {
	// The rendered Dockerfile data, for use in CloudFormation.
	// Default: - none if dockerfileTemplateUri is set.
	//
	// Experimental.
	DockerfileTemplateData *string `field:"optional" json:"dockerfileTemplateData" yaml:"dockerfileTemplateData"`
	// The rendered Dockerfile URI, for use in CloudFormation.
	// Default: - none if dockerfileTemplateData is set.
	//
	// Experimental.
	DockerfileTemplateUri *string `field:"optional" json:"dockerfileTemplateUri" yaml:"dockerfileTemplateUri"`
}

