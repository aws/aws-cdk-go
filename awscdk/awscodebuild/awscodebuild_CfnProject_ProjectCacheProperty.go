package awscodebuild


// `ProjectCache` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies information about the cache for the build project. If `ProjectCache` is not specified, then both of its properties default to `NO_CACHE` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectCacheProperty := &projectCacheProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	location: jsii.String("location"),
//   	modes: []*string{
//   		jsii.String("modes"),
//   	},
//   }
//
type CfnProject_ProjectCacheProperty struct {
	// The type of cache used by the build project. Valid values include:.
	//
	// - `NO_CACHE` : The build project does not use any cache.
	// - `S3` : The build project reads and writes from and to S3.
	// - `LOCAL` : The build project stores a cache locally on a build host that is only available to that build host.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Information about the cache location:.
	//
	// - `NO_CACHE` or `LOCAL` : This value is ignored.
	// - `S3` : This is the S3 bucket name/prefix.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// An array of strings that specify the local cache modes.
	//
	// You can use one or more local cache modes at the same time. This is only used for `LOCAL` cache types.
	//
	// Possible values are:
	//
	// - **LOCAL_SOURCE_CACHE** - Caches Git metadata for primary and secondary sources. After the cache is created, subsequent builds pull only the change between commits. This mode is a good choice for projects with a clean working directory and a source that is a large Git repository. If you choose this option and your project does not use a Git repository (GitHub, GitHub Enterprise, or Bitbucket), the option is ignored.
	// - **LOCAL_DOCKER_LAYER_CACHE** - Caches existing Docker layers. This mode is a good choice for projects that build or pull large Docker images. It can prevent the performance issues caused by pulling large Docker images down from the network.
	//
	// > - You can use a Docker layer cache in the Linux environment only.
	// > - The `privileged` flag must be set so that your project has the required Docker permissions.
	// > - You should consider the security implications before you use a Docker layer cache.
	// - **LOCAL_CUSTOM_CACHE** - Caches directories you specify in the buildspec file. This mode is a good choice if your build scenario is not suited to one of the other three local cache modes. If you use a custom cache:
	//
	// - Only directories can be specified for caching. You cannot specify individual files.
	// - Symlinks are used to reference cached directories.
	// - Cached directories are linked to your build before it downloads its project sources. Cached items are overridden if a source item has the same name. Directories are specified using cache paths in the buildspec file.
	Modes *[]*string `field:"optional" json:"modes" yaml:"modes"`
}

