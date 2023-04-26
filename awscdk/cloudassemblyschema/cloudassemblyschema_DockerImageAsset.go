package cloudassemblyschema


// A file asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageAsset := &DockerImageAsset{
//   	Destinations: map[string]dockerImageDestination{
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
//   	Source: &DockerImageSource{
//   		Directory: jsii.String("directory"),
//   		DockerBuildArgs: map[string]*string{
//   			"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   		},
//   		DockerBuildTarget: jsii.String("dockerBuildTarget"),
//   		DockerFile: jsii.String("dockerFile"),
//   		Executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		NetworkMode: jsii.String("networkMode"),
//   		Platform: jsii.String("platform"),
//   	},
//   }
//
// Experimental.
type DockerImageAsset struct {
	// Destinations for this file asset.
	// Experimental.
	Destinations *map[string]*DockerImageDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	// Experimental.
	Source *DockerImageSource `field:"required" json:"source" yaml:"source"`
}

