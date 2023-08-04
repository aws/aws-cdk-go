package awsecrassets


// Options for configuring the Docker cache backend.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &DockerImageAssetProps{
//   	Directory: path.join(__dirname, jsii.String("my-image")),
//   	CacheFrom: []dockerCacheOption{
//   		&dockerCacheOption{
//   			Type: jsii.String("registry"),
//   			Params: map[string]*string{
//   				"ref": jsii.String("ghcr.io/myorg/myimage:cache"),
//   			},
//   		},
//   	},
//   	CacheTo: &dockerCacheOption{
//   		Type: jsii.String("registry"),
//   		Params: map[string]*string{
//   			"ref": jsii.String("ghcr.io/myorg/myimage:cache"),
//   			"mode": jsii.String("max"),
//   			"compression": jsii.String("zstd"),
//   		},
//   	},
//   })
//
type DockerCacheOption struct {
	// The type of cache to use.
	//
	// Refer to https://docs.docker.com/build/cache/backends/ for full list of backends.
	//
	// Example:
	//   "registry"
	//
	// Default: - unspecified.
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Any parameters to pass into the docker cache backend configuration.
	//
	// Refer to https://docs.docker.com/build/cache/backends/ for cache backend configuration.
	//
	// Example:
	//   var branch string
	//
	//
	//   params := map[string]interface{}{
	//   	"ref": fmt.Sprintf("12345678.dkr.ecr.us-west-2.amazonaws.com/cache:%v", branch),
	//   	"mode": jsii.String("max"),
	//   }
	//
	// Default: {} No options provided.
	//
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

