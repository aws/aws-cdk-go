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
//   	File: jsii.String("file"),
//   	ImageNameParameter: jsii.String("imageNameParameter"),
//   	ImageTag: jsii.String("imageTag"),
//   	NetworkMode: jsii.String("networkMode"),
//   	Platform: jsii.String("platform"),
//   	RepositoryName: jsii.String("repositoryName"),
//   	Target: jsii.String("target"),
//   }
//
// Experimental.
type ContainerImageAssetMetadataEntry struct {
	// Logical identifier for the asset.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Type of asset.
	// Experimental.
	Packaging *string `field:"required" json:"packaging" yaml:"packaging"`
	// Path on disk to the asset.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The hash of the asset source.
	// Experimental.
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
	// Build args to pass to the `docker build` command.
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// ECR Repository name and repo digest (separated by "@sha256:") where this image is stored.
	// Deprecated: specify `repositoryName` and `imageTag` instead, and then you
	// know where the image will go.
	ImageNameParameter *string `field:"optional" json:"imageNameParameter" yaml:"imageNameParameter"`
	// The docker image tag to use for tagging pushed images.
	//
	// This field is
	// required if `imageParameterName` is ommited (otherwise, the app won't be
	// able to find the image).
	// Experimental.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
	// Networking mode for the RUN commands during build.
	// Experimental.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Platform to build for.
	//
	// _Requires Docker Buildx_.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// ECR repository name, if omitted a default name based on the asset's ID is used instead.
	//
	// Specify this property if you need to statically address the
	// image, e.g. from a Kubernetes Pod. Note, this is only the repository name,
	// without the registry and the tag parts.
	// Experimental.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

