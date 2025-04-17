package awslambda


// Result of the bind when an ECR image is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeImageConfig := &CodeImageConfig{
//   	ImageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	Cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	Entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
type CodeImageConfig struct {
	// URI to the Docker image.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	// Default: - use the CMD specified in the docker image or Dockerfile.
	//
	Cmd *[]*string `field:"optional" json:"cmd" yaml:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Default: - use the ENTRYPOINT in the docker image or Dockerfile.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	// Default: - use the WORKDIR in the docker image or Dockerfile.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

