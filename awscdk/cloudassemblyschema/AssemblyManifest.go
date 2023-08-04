package cloudassemblyschema


// A manifest which describes the cloud assembly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assemblyManifest := &AssemblyManifest{
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	Artifacts: map[string]artifactManifest{
//   		"artifactsKey": &artifactManifest{
//   			"type": awscdk.cloud_assembly_schema.ArtifactType_NONE,
//
//   			// the properties below are optional
//   			"dependencies": []*string{
//   				jsii.String("dependencies"),
//   			},
//   			"displayName": jsii.String("displayName"),
//   			"environment": jsii.String("environment"),
//   			"metadata": map[string][]MetadataEntry{
//   				"metadataKey": []MetadataEntry{
//   					&MetadataEntry{
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"data": jsii.String("data"),
//   						"trace": []*string{
//   							jsii.String("trace"),
//   						},
//   					},
//   				},
//   			},
//   			"properties": &AwsCloudFormationStackProperties{
//   				"templateFile": jsii.String("templateFile"),
//
//   				// the properties below are optional
//   				"assumeRoleArn": jsii.String("assumeRoleArn"),
//   				"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   				"bootstrapStackVersionSsmParameter": jsii.String("bootstrapStackVersionSsmParameter"),
//   				"cloudFormationExecutionRoleArn": jsii.String("cloudFormationExecutionRoleArn"),
//   				"lookupRole": &BootstrapRole{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"bootstrapStackVersionSsmParameter": jsii.String("bootstrapStackVersionSsmParameter"),
//   					"requiresBootstrapStackVersion": jsii.Number(123),
//   				},
//   				"parameters": map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				"requiresBootstrapStackVersion": jsii.Number(123),
//   				"stackName": jsii.String("stackName"),
//   				"stackTemplateAssetObjectUrl": jsii.String("stackTemplateAssetObjectUrl"),
//   				"tags": map[string]*string{
//   					"tagsKey": jsii.String("tags"),
//   				},
//   				"terminationProtection": jsii.Boolean(false),
//   				"validateOnSynth": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	Missing: []missingContext{
//   		&missingContext{
//   			Key: jsii.String("key"),
//   			Props: &AmiContextQuery{
//   				Account: jsii.String("account"),
//   				Filters: map[string][]*string{
//   					"filtersKey": []*string{
//   						jsii.String("filters"),
//   					},
//   				},
//   				Region: jsii.String("region"),
//
//   				// the properties below are optional
//   				LookupRoleArn: jsii.String("lookupRoleArn"),
//   				Owners: []interface{}{
//   					jsii.String("owners"),
//   				},
//   			},
//   			Provider: awscdk.Cloud_assembly_schema.ContextProvider_AMI_PROVIDER,
//   		},
//   	},
//   	Runtime: &RuntimeInfo{
//   		Libraries: map[string]*string{
//   			"librariesKey": jsii.String("libraries"),
//   		},
//   	},
//   }
//
type AssemblyManifest struct {
	// Protocol version.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The set of artifacts in this assembly.
	// Default: - no artifacts.
	//
	Artifacts *map[string]*ArtifactManifest `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Missing context information.
	//
	// If this field has values, it means that the
	// cloud assembly is not complete and should not be deployed.
	// Default: - no missing context.
	//
	Missing *[]*MissingContext `field:"optional" json:"missing" yaml:"missing"`
	// Runtime information.
	// Default: - no info.
	//
	Runtime *RuntimeInfo `field:"optional" json:"runtime" yaml:"runtime"`
}

