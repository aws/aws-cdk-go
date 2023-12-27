package cloudassemblyschema


// Definitions for the asset manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetManifest := &AssetManifest{
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	DockerImages: map[string]dockerImageAsset{
//   		"dockerImagesKey": &dockerImageAsset{
//   			"destinations": map[string]DockerImageDestination{
//   				"destinationsKey": &DockerImageDestination{
//   					"imageTag": jsii.String("imageTag"),
//   					"repositoryName": jsii.String("repositoryName"),
//
//   					// the properties below are optional
//   					"assumeRoleArn": jsii.String("assumeRoleArn"),
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"region": jsii.String("region"),
//   				},
//   			},
//   			"source": &DockerImageSource{
//   				"cacheDisabled": jsii.Boolean(false),
//   				"cacheFrom": []DockerCacheOption{
//   					&DockerCacheOption{
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"params": map[string]*string{
//   							"paramsKey": jsii.String("params"),
//   						},
//   					},
//   				},
//   				"cacheTo": &DockerCacheOption{
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"params": map[string]*string{
//   						"paramsKey": jsii.String("params"),
//   					},
//   				},
//   				"directory": jsii.String("directory"),
//   				"dockerBuildArgs": map[string]*string{
//   					"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   				},
//   				"dockerBuildSecrets": map[string]*string{
//   					"dockerBuildSecretsKey": jsii.String("dockerBuildSecrets"),
//   				},
//   				"dockerBuildSsh": jsii.String("dockerBuildSsh"),
//   				"dockerBuildTarget": jsii.String("dockerBuildTarget"),
//   				"dockerFile": jsii.String("dockerFile"),
//   				"dockerOutputs": []*string{
//   					jsii.String("dockerOutputs"),
//   				},
//   				"executable": []*string{
//   					jsii.String("executable"),
//   				},
//   				"networkMode": jsii.String("networkMode"),
//   				"platform": jsii.String("platform"),
//   			},
//   		},
//   	},
//   	Files: map[string]fileAsset{
//   		"filesKey": &fileAsset{
//   			"destinations": map[string]FileDestination{
//   				"destinationsKey": &FileDestination{
//   					"bucketName": jsii.String("bucketName"),
//   					"objectKey": jsii.String("objectKey"),
//
//   					// the properties below are optional
//   					"assumeRoleArn": jsii.String("assumeRoleArn"),
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"region": jsii.String("region"),
//   				},
//   			},
//   			"source": &FileSource{
//   				"executable": []*string{
//   					jsii.String("executable"),
//   				},
//   				"packaging": awscdk.cloud_assembly_schema.FileAssetPackaging_FILE,
//   				"path": jsii.String("path"),
//   			},
//   		},
//   	},
//   }
//
type AssetManifest struct {
	// Version of the manifest.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The Docker image assets in this manifest.
	// Default: - No Docker images.
	//
	DockerImages *map[string]*DockerImageAsset `field:"optional" json:"dockerImages" yaml:"dockerImages"`
	// The file assets in this manifest.
	// Default: - No files.
	//
	Files *map[string]*FileAsset `field:"optional" json:"files" yaml:"files"`
}

