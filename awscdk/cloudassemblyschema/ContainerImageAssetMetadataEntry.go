package cloudassemblyschema


// Metadata Entry spec for container images.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerImageAssetMetadataEntry := &ContainerImageAssetMetadataEntry{
//   	Id: jsii.String("id"),
//   	Packaging: jsii.String("packaging"),
//   	Path: jsii.String("path"),
//   	SourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	BuildSecrets: map[string]*string{
//   		"buildSecretsKey": jsii.String("buildSecrets"),
//   	},
//   	File: jsii.String("file"),
//   	ImageTag: jsii.String("imageTag"),
//   	NetworkMode: jsii.String("networkMode"),
//   	Outputs: []*string{
//   		jsii.String("outputs"),
//   	},
//   	Platform: jsii.String("platform"),
//   	RepositoryName: jsii.String("repositoryName"),
//   	Target: jsii.String("target"),
//   }
//
type ContainerImageAssetMetadataEntry struct {
	// Logical identifier for the asset.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Type of asset.
	Packaging *string `field:"required" json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The hash of the asset source.
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
	// Build args to pass to the `docker build` command.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Build secrets to pass to the `docker build` command.
	BuildSecrets *map[string]*string `field:"optional" json:"buildSecrets" yaml:"buildSecrets"`
	// Path to the Dockerfile (relative to the directory).
	File *string `field:"optional" json:"file" yaml:"file"`
	// The docker image tag to use for tagging pushed images.
	//
	// This field is
	// required if `imageParameterName` is ommited (otherwise, the app won't be
	// able to find the image).
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
	// Networking mode for the RUN commands during build.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Outputs to pass to the `docker build` command.
	// See: https://docs.docker.com/engine/reference/commandline/build/#custom-build-outputs
	//
	Outputs *[]*string `field:"optional" json:"outputs" yaml:"outputs"`
	// Platform to build for.
	//
	// _Requires Docker Buildx_.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// ECR repository name, if omitted a default name based on the asset's ID is used instead.
	//
	// Specify this property if you need to statically address the
	// image, e.g. from a Kubernetes Pod. Note, this is only the repository name,
	// without the registry and the tag parts.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

