package awslambda


// Options when creating an asset from a Docker build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerBuildAssetOptions := &DockerBuildAssetOptions{
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	File: jsii.String("file"),
//   	ImagePath: jsii.String("imagePath"),
//   	OutputPath: jsii.String("outputPath"),
//   	Platform: jsii.String("platform"),
//   	TargetStage: jsii.String("targetStage"),
//   }
//
type DockerBuildAssetOptions struct {
	// Build args.
	// Default: - no build args.
	//
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	// Default: `Dockerfile`.
	//
	File *string `field:"optional" json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Default: - no platform specified.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Set build target for multi-stage container builds. Any stage defined afterwards will be ignored.
	//
	// Example value: `build-env`.
	// Default: - Build all stages defined in the Dockerfile.
	//
	TargetStage *string `field:"optional" json:"targetStage" yaml:"targetStage"`
	// The path in the Docker image where the asset is located after the build operation.
	// Default: /asset.
	//
	ImagePath *string `field:"optional" json:"imagePath" yaml:"imagePath"`
	// The path on the local filesystem where the asset will be copied using `docker cp`.
	// Default: - a unique temporary directory in the system temp directory.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
}

