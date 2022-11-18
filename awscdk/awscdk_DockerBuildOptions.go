// Version 2 of the AWS Cloud Development Kit library
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
type DockerBuildOptions struct {
	// Build args.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Set build target for multi-stage container builds. Any stage defined afterwards will be ignored.
	//
	// Example value: `build-env`.
	TargetStage *string `field:"optional" json:"targetStage" yaml:"targetStage"`
}

