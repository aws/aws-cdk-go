package cloudassemblyschema


// Metadata Entry spec for container images.
//
// Example:
//   entry := map[string]*string{
//   	"packaging": jsii.String("container-image"),
//   	"repositoryName": jsii.String("repository-name"),
//   	"imageTag": jsii.String("tag"),
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
	// Cache from options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	CacheFrom *[]*ContainerImageAssetCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	CacheTo *ContainerImageAssetCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
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

