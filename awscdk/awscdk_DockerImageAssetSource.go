// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageAssetSource := &dockerImageAssetSource{
//   	sourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	directoryName: jsii.String("directoryName"),
//   	dockerBuildArgs: map[string]*string{
//   		"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   	},
//   	dockerBuildTarget: jsii.String("dockerBuildTarget"),
//   	dockerFile: jsii.String("dockerFile"),
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	networkMode: jsii.String("networkMode"),
//   	platform: jsii.String("platform"),
//   	repositoryName: jsii.String("repositoryName"),
//   }
//
// Experimental.
type DockerImageAssetSource struct {
	// The hash of the contents of the docker build context.
	//
	// This hash is used
	// throughout the system to identify this image and avoid duplicate work
	// in case the source did not change.
	//
	// NOTE: this means that if you wish to update your docker image, you
	// must make a modification to the source (e.g. add some metadata to your Dockerfile).
	// Experimental.
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
	// The directory where the Dockerfile is stored, must be relative to the cloud assembly root.
	// Experimental.
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	//
	// Only allowed when `directoryName` is specified.
	// Experimental.
	DockerBuildArgs *map[string]*string `field:"optional" json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Docker target to build to.
	//
	// Only allowed when `directoryName` is specified.
	// Experimental.
	DockerBuildTarget *string `field:"optional" json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// Path to the Dockerfile (relative to the directory).
	//
	// Only allowed when `directoryName` is specified.
	// Experimental.
	DockerFile *string `field:"optional" json:"dockerFile" yaml:"dockerFile"`
	// An external command that will produce the packaged asset.
	//
	// The command should produce the name of a local Docker image on `stdout`.
	// Experimental.
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	// Experimental.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Platform to build for. _Requires Docker Buildx_.
	//
	// Specify this property to build images on a specific platform.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// ECR repository name.
	//
	// Specify this property if you need to statically address the image, e.g.
	// from a Kubernetes Pod. Note, this is only the repository name, without the
	// registry and the tag parts.
	// Deprecated: repository name should be specified at the environment-level and not at the image level.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

