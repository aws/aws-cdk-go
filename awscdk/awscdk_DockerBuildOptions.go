// An experiment to bundle the entire CDK into a single module
package awscdk


// Docker build options.
//
// Example:
//   lambda.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: lambda.code.fromAsset(jsii.String("/path/to/handler"), &assetOptions{
//   		bundling: &bundlingOptions{
//   			image: awscdk.DockerImage.fromBuild(jsii.String("/path/to/dir/with/DockerFile"), &dockerBuildOptions{
//   				buildArgs: map[string]*string{
//   					"ARG1": jsii.String("value1"),
//   				},
//   			}),
//   			command: []*string{
//   				jsii.String("my"),
//   				jsii.String("cool"),
//   				jsii.String("command"),
//   			},
//   		},
//   	}),
//   	runtime: lambda.runtime_PYTHON_3_9(),
//   	handler: jsii.String("index.handler"),
//   })
//
// Experimental.
type DockerBuildOptions struct {
	// Build args.
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

