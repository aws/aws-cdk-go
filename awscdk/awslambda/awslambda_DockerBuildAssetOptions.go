package awslambda


// Options when creating an asset from a Docker build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerBuildAssetOptions := &dockerBuildAssetOptions{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	file: jsii.String("file"),
//   	imagePath: jsii.String("imagePath"),
//   	outputPath: jsii.String("outputPath"),
//   	platform: jsii.String("platform"),
//   	targetStage: jsii.String("targetStage"),
//   }
//
type DockerBuildAssetOptions struct {
	// Build args.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Set build target for multi-stage container builds. Any stage defined afterwards will be ignored.
	//
	// Example value: `build-env`.
	TargetStage *string `field:"optional" json:"targetStage" yaml:"targetStage"`
	// The path in the Docker image where the asset is located after the build operation.
	ImagePath *string `field:"optional" json:"imagePath" yaml:"imagePath"`
	// The path on the local filesystem where the asset will be copied using `docker cp`.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
}

