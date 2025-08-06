package awslambdanodejs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Bundling options.
//
// Example:
//   nodejs.NewNodejsFunction(this, jsii.String("my-handler"), &NodejsFunctionProps{
//   	Bundling: &BundlingOptions{
//   		DockerImage: awscdk.DockerImage_FromBuild(jsii.String("/path/to/Dockerfile")),
//   	},
//   })
//
type BundlingOptions struct {
	// The command to run in the container.
	// Default: - run the command defined in the image.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	// Default: - run the entrypoint defined in the image.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	// Default: - no environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	// Default: - no networking options.
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Default: - no platform specified.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Default: - no security options.
	//
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	// Default: - root or image default.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docker volumes to mount.
	// Default: - no volumes are mounted.
	//
	Volumes *[]*awscdk.DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	// Default: - no containers are specified to mount volumes from.
	//
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the container.
	// Default: - image default.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// Specify a custom hash for this asset.
	//
	// For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Default: - asset hash is calculated based on the bundled output.
	//
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Use this to insert an arbitrary string at the beginning of generated JavaScript files.
	//
	// This is similar to footer which inserts at the end instead of the beginning.
	//
	// This is commonly used to insert comments:.
	// Default: - no comments are passed.
	//
	Banner *string `field:"optional" json:"banner" yaml:"banner"`
	// Build arguments to pass when building the bundling image.
	// Default: - no build arguments are passed.
	//
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Includes AWS SDK in the bundle asset.
	// Default: - false
	// if `true` the `aws-sdk` will be included in the asset bundle and not be
	// resolved to the Lambda provided sdk.
	//
	BundleAwsSDK *bool `field:"optional" json:"bundleAwsSDK" yaml:"bundleAwsSDK"`
	// Which option to use to copy the source files to the docker container and output files back.
	// Default: - BundlingFileAccess.BIND_MOUNT
	//
	BundlingFileAccess awscdk.BundlingFileAccess `field:"optional" json:"bundlingFileAccess" yaml:"bundlingFileAccess"`
	// The charset to use for esbuild's output.
	//
	// By default esbuild's output is ASCII-only. Any non-ASCII characters are escaped
	// using backslash escape sequences. Using escape sequences makes the generated output
	// slightly bigger, and also makes it harder to read. If you would like for esbuild to print
	// the original characters without using escape sequences, use `Charset.UTF8`.
	// See: https://esbuild.github.io/api/#charset
	//
	// Default: Charset.ASCII
	//
	Charset Charset `field:"optional" json:"charset" yaml:"charset"`
	// Command hooks.
	// Default: - do not run additional commands.
	//
	CommandHooks ICommandHooks `field:"optional" json:"commandHooks" yaml:"commandHooks"`
	// Replace global identifiers with constant expressions.
	//
	// For example, `{ 'process.env.DEBUG': 'true' }`.
	//
	// Another example, `{ 'process.env.API_KEY': JSON.stringify('xxx-xxxx-xxx') }`.
	// Default: - no replacements are made.
	//
	Define *map[string]*string `field:"optional" json:"define" yaml:"define"`
	// A custom bundling Docker image.
	//
	// This image should have esbuild installed globally. If you plan to use `nodeModules`
	// it should also have `npm`, `yarn`, `bun` or `pnpm` depending on the lock file you're using.
	//
	// See https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/aws-lambda-nodejs/lib/Dockerfile
	// for the default image provided by aws-cdk-lib/aws-lambda-nodejs.
	// Default: - use the Docker image provided by aws-cdk-lib/aws-lambda-nodejs.
	//
	DockerImage awscdk.DockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Build arguments to pass into esbuild.
	//
	// For example, to add the [--log-limit](https://esbuild.github.io/api/#log-limit) flag:
	//
	// ```text
	// new NodejsFunction(scope, id, {
	//   ...
	//   bundling: {
	//     esbuildArgs: {
	//       "--log-limit": "0",
	//     }
	//   }
	// });
	// ```.
	// Default: - no additional esbuild arguments are passed.
	//
	EsbuildArgs *map[string]interface{} `field:"optional" json:"esbuildArgs" yaml:"esbuildArgs"`
	// The version of esbuild to use when running in a Docker container.
	// Default: - latest v0.
	//
	EsbuildVersion *string `field:"optional" json:"esbuildVersion" yaml:"esbuildVersion"`
	// A list of modules that should be considered as externals (already available in the runtime).
	// Default: - no replacements are made.
	//
	ExternalModules *[]*string `field:"optional" json:"externalModules" yaml:"externalModules"`
	// Use this to insert an arbitrary string at the end of generated JavaScript files.
	//
	// This is similar to banner which inserts at the beginning instead of the end.
	//
	// This is commonly used to insert comments.
	// Default: - no comments are passed.
	//
	Footer *string `field:"optional" json:"footer" yaml:"footer"`
	// Force bundling in a Docker container even if local bundling is possible.
	//
	// This is useful if your function relies on node modules
	// that should be installed (`nodeModules`) in a Lambda compatible
	// environment.
	// Default: false.
	//
	ForceDockerBundling *bool `field:"optional" json:"forceDockerBundling" yaml:"forceDockerBundling"`
	// Output format for the generated JavaScript files.
	// Default: OutputFormat.CJS
	//
	Format OutputFormat `field:"optional" json:"format" yaml:"format"`
	// This option allows you to automatically replace a global variable with an import from another file.
	// See: https://esbuild.github.io/api/#inject
	//
	// Default: - no code is injected.
	//
	Inject *[]*string `field:"optional" json:"inject" yaml:"inject"`
	// Whether to preserve the original `name` values even in minified code.
	//
	// In JavaScript the `name` property on functions and classes defaults to a
	// nearby identifier in the source code.
	//
	// However, minification renames symbols to reduce code size and bundling
	// sometimes need to rename symbols to avoid collisions. That changes value of
	// the `name` property for many of these cases. This is usually fine because
	// the `name` property is normally only used for debugging. However, some
	// frameworks rely on the `name` property for registration and binding purposes.
	// If this is the case, you can enable this option to preserve the original
	// `name` values even in minified code.
	// Default: false.
	//
	KeepNames *bool `field:"optional" json:"keepNames" yaml:"keepNames"`
	// Use loaders to change how a given input file is interpreted.
	//
	// Configuring a loader for a given file type lets you load that file type with
	// an `import` statement or a `require` call.
	//
	// For example, `{ '.png': 'dataurl' }`.
	// See: https://esbuild.github.io/api/#loader
	//
	// Default: - use esbuild default loaders.
	//
	Loader *map[string]*string `field:"optional" json:"loader" yaml:"loader"`
	// Log level for esbuild.
	//
	// This is also propagated to the package manager and
	// applies to its specific install command.
	// Default: LogLevel.WARNING
	//
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// How to determine the entry point for modules.
	//
	// Try ['module', 'main'] to default to ES module versions.
	// Default: [].
	//
	MainFields *[]*string `field:"optional" json:"mainFields" yaml:"mainFields"`
	// This option tells esbuild to write out a JSON file relative to output directory with metadata about the build.
	//
	// The metadata in this JSON file follows this schema (specified using TypeScript syntax):
	//
	// ```text
	// {
	//   outputs: {
	//     [path: string]: {
	//       bytes: number
	//       inputs: {
	//         [path: string]: { bytesInOutput: number }
	//       }
	//       imports: { path: string }[]
	//       exports: string[]
	//     }
	//   }
	// }
	// ```
	// This data can then be analyzed by other tools. For example,
	// bundle buddy can consume esbuild's metadata format and generates a treemap visualization
	// of the modules in your bundle and how much space each one takes up.
	// See: https://esbuild.github.io/api/#metafile
	//
	// Default: false.
	//
	Metafile *bool `field:"optional" json:"metafile" yaml:"metafile"`
	// Whether to minify files when bundling.
	// Default: false.
	//
	Minify *bool `field:"optional" json:"minify" yaml:"minify"`
	// A list of modules that should be installed instead of bundled.
	//
	// Modules are
	// installed in a Lambda compatible environment only when bundling runs in
	// Docker.
	// Default: - all modules are bundled.
	//
	NodeModules *[]*string `field:"optional" json:"nodeModules" yaml:"nodeModules"`
	// Run compilation using tsc before running file through bundling step.
	//
	// This usually is not required unless you are using new experimental features that
	// are only supported by typescript's `tsc` compiler.
	// One example of such feature is `emitDecoratorMetadata`.
	// Default: false.
	//
	PreCompilation *bool `field:"optional" json:"preCompilation" yaml:"preCompilation"`
	// Whether to include source maps when bundling.
	// Default: false.
	//
	SourceMap *bool `field:"optional" json:"sourceMap" yaml:"sourceMap"`
	// Source map mode to be used when bundling.
	// See: https://esbuild.github.io/api/#sourcemap
	//
	// Default: SourceMapMode.DEFAULT
	//
	SourceMapMode SourceMapMode `field:"optional" json:"sourceMapMode" yaml:"sourceMapMode"`
	// Whether to include original source code in source maps when bundling.
	// See: https://esbuild.github.io/api/#sources-content
	//
	// Default: true.
	//
	SourcesContent *bool `field:"optional" json:"sourcesContent" yaml:"sourcesContent"`
	// Target environment for the generated JavaScript code.
	// See: https://esbuild.github.io/api/#target
	//
	// Default: - the node version of the runtime.
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Normally the esbuild automatically discovers `tsconfig.json` files and reads their contents during a build.
	//
	// However, you can also configure a custom `tsconfig.json` file to use instead.
	//
	// This is similar to entry path, you need to provide path to your custom `tsconfig.json`.
	//
	// This can be useful if you need to do multiple builds of the same code with different settings.
	//
	// For example, `{ 'tsconfig': 'path/custom.tsconfig.json' }`.
	// Default: - automatically discovered by `esbuild`.
	//
	Tsconfig *string `field:"optional" json:"tsconfig" yaml:"tsconfig"`
}

