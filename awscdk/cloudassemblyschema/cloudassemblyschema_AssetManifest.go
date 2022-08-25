package cloudassemblyschema


// Definitions for the asset manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetManifest := &assetManifest{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	dockerImages: map[string]dockerImageAsset{
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
//   				"directory": jsii.String("directory"),
//   				"dockerBuildArgs": map[string]*string{
//   					"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   				},
//   				"dockerBuildTarget": jsii.String("dockerBuildTarget"),
//   				"dockerFile": jsii.String("dockerFile"),
//   				"executable": []*string{
//   					jsii.String("executable"),
//   				},
//   				"networkMode": jsii.String("networkMode"),
//   				"platform": jsii.String("platform"),
//   			},
//   		},
//   	},
//   	files: map[string]fileAsset{
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
	DockerImages *map[string]*DockerImageAsset `field:"optional" json:"dockerImages" yaml:"dockerImages"`
	// The file assets in this manifest.
	Files *map[string]*FileAsset `field:"optional" json:"files" yaml:"files"`
}

