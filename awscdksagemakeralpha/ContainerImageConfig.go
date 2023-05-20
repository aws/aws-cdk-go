package awscdksagemakeralpha


// The configuration for creating a container image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   containerImageConfig := &ContainerImageConfig{
//   	ImageName: jsii.String("imageName"),
//   }
//
// Experimental.
type ContainerImageConfig struct {
	// The image name. Images in Amazon ECR repositories can be specified by either using the full registry/repository:tag or registry/repository@digest.
	//
	// For example, `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest` or
	// `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE`.
	// Experimental.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
}

