package cloudassemblyschema


// Options for configuring the Docker cache backend.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerImageAssetCacheOption := &ContainerImageAssetCacheOption{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Params: map[string]*string{
//   		"paramsKey": jsii.String("params"),
//   	},
//   }
//
type ContainerImageAssetCacheOption struct {
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

