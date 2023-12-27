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
//   		CacheDisabled: jsii.Boolean(false),
//   		CacheFrom: []dockerCacheOption{
//   			&dockerCacheOption{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Params: map[string]*string{
//   					"paramsKey": jsii.String("params"),
//   				},
//   			},
//   		},
//   		CacheTo: &dockerCacheOption{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Params: map[string]*string{
//   				"paramsKey": jsii.String("params"),
//   			},
//   		},
//   		Directory: jsii.String("directory"),
//   		DockerBuildArgs: map[string]*string{
//   			"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   		},
//   		DockerBuildSecrets: map[string]*string{
//   			"dockerBuildSecretsKey": jsii.String("dockerBuildSecrets"),
//   		},
//   		DockerBuildSsh: jsii.String("dockerBuildSsh"),
//   		DockerBuildTarget: jsii.String("dockerBuildTarget"),
//   		DockerFile: jsii.String("dockerFile"),
//   		DockerOutputs: []*string{
//   			jsii.String("dockerOutputs"),
//   		},
//   		Executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		NetworkMode: jsii.String("networkMode"),
//   		Platform: jsii.String("platform"),
//   	},
//   }
//
type DockerImageAsset struct {
	// Destinations for this file asset.
	Destinations *map[string]*DockerImageDestination `field:"required" json:"destinations" yaml:"destinations"`
	// Source description for file assets.
	Source *DockerImageSource `field:"required" json:"source" yaml:"source"`
}

