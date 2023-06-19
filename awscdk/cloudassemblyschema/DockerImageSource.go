package cloudassemblyschema


// Properties for how to produce a Docker image from a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageSource := &DockerImageSource{
//   	Directory: jsii.String("directory"),
//   	DockerBuildArgs: map[string]*string{
//   		"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   	},
//   	DockerBuildTarget: jsii.String("dockerBuildTarget"),
//   	DockerFile: jsii.String("dockerFile"),
//   	Executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	NetworkMode: jsii.String("networkMode"),
//   	Platform: jsii.String("platform"),
//   }
//
// Experimental.
type DockerImageSource struct {
	// The directory containing the Docker image build instructions.
	//
	// This path is relative to the asset manifest location.
	// Experimental.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Additional build arguments.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerBuildArgs *map[string]*string `field:"optional" json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Target build stage in a Dockerfile with multiple build stages.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerBuildTarget *string `field:"optional" json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// The name of the file with build instructions.
	//
	// Only allowed when `directory` is set.
	// Experimental.
	DockerFile *string `field:"optional" json:"dockerFile" yaml:"dockerFile"`
	// A command-line executable that returns the name of a local Docker image on stdout after being run.
	// Experimental.
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	// Experimental.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Platform to build for. _Requires Docker Buildx_.
	//
	// Specify this property to build images on a specific platform/architecture.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

