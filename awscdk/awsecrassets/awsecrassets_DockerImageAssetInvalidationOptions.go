package awsecrassets


// Options to control invalidation of `DockerImageAsset` asset hashes.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &DockerImageAssetProps{
//   	Directory: path.join(__dirname, jsii.String("my-image")),
//   	BuildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	Invalidation: &DockerImageAssetInvalidationOptions{
//   		BuildArgs: jsii.Boolean(false),
//   	},
//   })
//
// Experimental.
type DockerImageAssetInvalidationOptions struct {
	// Use `buildArgs` while calculating the asset hash.
	// Experimental.
	BuildArgs *bool `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Use `extraHash` while calculating the asset hash.
	// Experimental.
	ExtraHash *bool `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Use `file` while calculating the asset hash.
	// Experimental.
	File *bool `field:"optional" json:"file" yaml:"file"`
	// Use `networkMode` while calculating the asset hash.
	// Experimental.
	NetworkMode *bool `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Use `platform` while calculating the asset hash.
	// Experimental.
	Platform *bool `field:"optional" json:"platform" yaml:"platform"`
	// Use `repositoryName` while calculating the asset hash.
	// Experimental.
	RepositoryName *bool `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Use `target` while calculating the asset hash.
	// Experimental.
	Target *bool `field:"optional" json:"target" yaml:"target"`
}

