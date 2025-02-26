package awscdklambdagoalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Bundling options.
//
// Example:
//   go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
//   	Entry: jsii.String("app/cmd/api"),
//   	Bundling: &BundlingOptions{
//   		Environment: map[string]*string{
//   			"HELLO": jsii.String("WORLD"),
//   		},
//   	},
//   })
//
// Experimental.
type BundlingOptions struct {
	// The command to run in the container.
	// Default: - run the command defined in the image.
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	// Default: - run the entrypoint defined in the image.
	//
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	// Default: - no environment variables.
	//
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	// Default: - no networking options.
	//
	// Experimental.
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Default: - no platform specified.
	//
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Default: - no security options.
	//
	// Experimental.
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	// Default: - root or image default.
	//
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docker volumes to mount.
	// Default: - no volumes are mounted.
	//
	// Experimental.
	Volumes *[]*awscdk.DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	// Default: - no containers are specified to mount volumes from.
	//
	// Experimental.
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the container.
	// Default: - image default.
	//
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Default: - based on `assetHashType`.
	//
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Determines how the asset hash is calculated. Assets will get rebuilt and uploaded only if their hash has changed.
	//
	// If the asset hash is set to `OUTPUT` (default), the hash is calculated
	// after bundling. This means that any change in the output will cause
	// the asset to be invalidated and uploaded. Bear in mind that the
	// go binary that is output can be different depending on the environment
	// that it was compiled in. If you want to control when the output is changed
	// it is recommended that you use immutable build images such as
	// `public.ecr.aws/bitnami/golang:1.16.3-debian-10-r16`.
	//
	// If the asset hash is set to `SOURCE`, then only changes to the source
	// directory will cause the asset to rebuild. If your go project has multiple
	// Lambda functions this means that an update to any one function could cause
	// all the functions to be rebuilt and uploaded.
	// Default: - AssetHashType.OUTPUT. If `assetHash` is also specified,
	// the default is `CUSTOM`.
	//
	// Experimental.
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Build arguments to pass when building the bundling image.
	// Default: - no build arguments are passed.
	//
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Which option to use to copy the source files to the docker container and output files back.
	// Default: - BundlingFileAccess.BIND_MOUNT
	//
	// Experimental.
	BundlingFileAccess awscdk.BundlingFileAccess `field:"optional" json:"bundlingFileAccess" yaml:"bundlingFileAccess"`
	// Whether or not to enable cgo during go build.
	//
	// This will set the CGO_ENABLED environment variable.
	// Default: - false.
	//
	// Experimental.
	CgoEnabled *bool `field:"optional" json:"cgoEnabled" yaml:"cgoEnabled"`
	// Command hooks.
	// Default: - do not run additional commands.
	//
	// Experimental.
	CommandHooks ICommandHooks `field:"optional" json:"commandHooks" yaml:"commandHooks"`
	// A custom bundling Docker image.
	// Default: - use the Docker image provided by.
	//
	// Experimental.
	DockerImage awscdk.DockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Force bundling in a Docker container even if local bundling is possible.
	// Default: - false.
	//
	// Experimental.
	ForcedDockerBundling *bool `field:"optional" json:"forcedDockerBundling" yaml:"forcedDockerBundling"`
	// List of additional flags to use while building.
	//
	// For example:
	// ['ldflags "-s -w"'].
	// Default: - none.
	//
	// Experimental.
	GoBuildFlags *[]*string `field:"optional" json:"goBuildFlags" yaml:"goBuildFlags"`
	// What Go proxies to use to fetch the packages involved in the build.
	//
	// Pass a list of proxy addresses in order, and/or the string `'direct'` to
	// attempt direct access.
	//
	// By default this construct uses no proxies, but a standard Go install would
	// use the Google proxy by default. To recreate that behavior, do the following:
	//
	// ```ts
	// new go.GoFunction(this, 'GoFunction', {
	//   entry: 'app/cmd/api',
	//   bundling: {
	//     goProxies: [go.GoFunction.GOOGLE_GOPROXY, 'direct'],
	//   },
	// });
	// ```.
	// Default: - Direct access.
	//
	// Experimental.
	GoProxies *[]*string `field:"optional" json:"goProxies" yaml:"goProxies"`
}

