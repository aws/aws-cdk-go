package awslambda


// Properties to initialize a new EcrImageCode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrImageCodeProps := &ecrImageCodeProps{
//   	cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	tag: jsii.String("tag"),
//   	tagOrDigest: jsii.String("tagOrDigest"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type EcrImageCodeProps struct {
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	Cmd *[]*string `field:"optional" json:"cmd" yaml:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The image tag to use when pulling the image from ECR.
	// Deprecated: use `tagOrDigest`.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The image tag or digest to use when pulling the image from ECR (digests must start with `sha256:`).
	TagOrDigest *string `field:"optional" json:"tagOrDigest" yaml:"tagOrDigest"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

