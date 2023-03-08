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
type DockerImageAssetInvalidationOptions struct {
	// Use `buildArgs` while calculating the asset hash.
	BuildArgs *bool `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Use `buildSecrets` while calculating the asset hash.
	BuildSecrets *bool `field:"optional" json:"buildSecrets" yaml:"buildSecrets"`
	// Use `extraHash` while calculating the asset hash.
	ExtraHash *bool `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Use `file` while calculating the asset hash.
	File *bool `field:"optional" json:"file" yaml:"file"`
	// Use `networkMode` while calculating the asset hash.
	NetworkMode *bool `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Use `outputs` while calculating the asset hash.
	Outputs *bool `field:"optional" json:"outputs" yaml:"outputs"`
	// Use `platform` while calculating the asset hash.
	Platform *bool `field:"optional" json:"platform" yaml:"platform"`
	// Use `repositoryName` while calculating the asset hash.
	RepositoryName *bool `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Use `target` while calculating the asset hash.
	Target *bool `field:"optional" json:"target" yaml:"target"`
}

