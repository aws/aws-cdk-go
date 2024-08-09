package awscdkcloudassemblyschema


// A file asset.
type DockerImageAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*DockerImageDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *DockerImageSource `field:"required" json:"source" yaml:"source"`
}

