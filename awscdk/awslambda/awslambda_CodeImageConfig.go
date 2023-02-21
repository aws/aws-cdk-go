package awslambda


// Result of the bind when an ECR image is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeImageConfig := &codeImageConfig{
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
// Experimental.
type CodeImageConfig struct {
	// URI to the Docker image.
	// Experimental.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	// Experimental.
	Cmd *[]*string `field:"optional" json:"cmd" yaml:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

