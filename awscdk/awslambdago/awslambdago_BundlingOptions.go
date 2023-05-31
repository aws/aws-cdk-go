package awslambdago

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Bundling options.
//
// Example:
//   lambda.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
//   	Entry: jsii.String("app/cmd/api"),
//   	Bundling: &BundlingOptions{
//   		DockerImage: awscdk.DockerImage_FromBuild(jsii.String("/path/to/Dockerfile")),
//   	},
//   })
//
// Experimental.
type BundlingOptions struct {
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
	// Experimental.
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Build arguments to pass when building the bundling image.
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Whether or not to enable cgo during go build.
	//
	// This will set the CGO_ENABLED environment variable.
	// Experimental.
	CgoEnabled *bool `field:"optional" json:"cgoEnabled" yaml:"cgoEnabled"`
	// Command hooks.
	// Experimental.
	CommandHooks ICommandHooks `field:"optional" json:"commandHooks" yaml:"commandHooks"`
	// A custom bundling Docker image.
	// Experimental.
	DockerImage awscdk.DockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Environment variables defined when go runs.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Force bundling in a Docker container even if local bundling is possible.
	// Experimental.
	ForcedDockerBundling *bool `field:"optional" json:"forcedDockerBundling" yaml:"forcedDockerBundling"`
	// List of additional flags to use while building.
	//
	// For example:
	// ['ldflags "-s -w"'].
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
	// new lambda.GoFunction(this, 'GoFunction', {
	//    entry: 'app/cmd/api',
	//    bundling: {
	//      goProxies: [lambda.GoFunction.GOOGLE_GOPROXY, 'direct'],
	//    },
	// });
	// ```.
	// Experimental.
	GoProxies *[]*string `field:"optional" json:"goProxies" yaml:"goProxies"`
}

