package cloudassemblyschema


// A file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageAsset := &dockerImageAsset{
//   	destinations: map[string]dockerImageDestination{
//   		"destinationsKey": &dockerImageDestination{
//   			"imageTag": jsii.String("imageTag"),
//   			"repositoryName": jsii.String("repositoryName"),
//
//   			// the properties below are optional
//   			"assumeRoleArn": jsii.String("assumeRoleArn"),
//   			"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   			"region": jsii.String("region"),
//   		},
//   	},
//   	source: &dockerImageSource{
//   		directory: jsii.String("directory"),
//   		dockerBuildArgs: map[string]*string{
//   			"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   		},
//   		dockerBuildTarget: jsii.String("dockerBuildTarget"),
//   		dockerFile: jsii.String("dockerFile"),
//   		executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		networkMode: jsii.String("networkMode"),
//   		platform: jsii.String("platform"),
//   	},
//   }
//
type DockerImageAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*DockerImageDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *DockerImageSource `field:"required" json:"source" yaml:"source"`
}

