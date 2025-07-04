package cloudassemblyschema


// A manifest for a single artifact within the cloud assembly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRoleAdditionalOptions interface{}
//
//   artifactManifest := &ArtifactManifest{
//   	Type: awscdk.Cloud_assembly_schema.ArtifactType_NONE,
//
//   	// the properties below are optional
//   	Dependencies: []*string{
//   		jsii.String("dependencies"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Environment: jsii.String("environment"),
//   	Metadata: map[string][]metadataEntry{
//   		"metadataKey": []*metadataEntry{
//   			&metadataEntry{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"data": jsii.String("data"),
//   				"trace": []*string{
//   					jsii.String("trace"),
//   				},
//   			},
//   		},
//   	},
//   	Properties: &AwsCloudFormationStackProperties{
//   		TemplateFile: jsii.String("templateFile"),
//
//   		// the properties below are optional
//   		AssumeRoleAdditionalOptions: map[string]interface{}{
//   			"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   		},
//   		AssumeRoleArn: jsii.String("assumeRoleArn"),
//   		AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		CloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   		LookupRole: &BootstrapRole{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			AssumeRoleAdditionalOptions: map[string]interface{}{
//   				"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   			},
//   			AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   			BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   			RequiresBootstrapStackVersion: jsii.Number(123),
//   		},
//   		NotificationArns: []*string{
//   			jsii.String("notificationArns"),
//   		},
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		RequiresBootstrapStackVersion: jsii.Number(123),
//   		StackName: jsii.String("stackName"),
//   		StackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		TerminationProtection: jsii.Boolean(false),
//   		ValidateOnSynth: jsii.Boolean(false),
//   	},
//   }
//
type ArtifactManifest struct {
	// The type of artifact.
	Type ArtifactType `field:"required" json:"type" yaml:"type"`
	// IDs of artifacts that must be deployed before this artifact.
	// Default: - no dependencies.
	//
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// A string that can be shown to a user to uniquely identify this artifact inside a cloud assembly tree.
	//
	// Is used by the CLI to present a list of stacks to the user in a way that
	// makes sense to them. Even though the property name "display name" doesn't
	// imply it, this field is used to select stacks as well, so all stacks should
	// have a unique display name.
	// Default: - no display name.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The environment into which this artifact is deployed.
	// Default: - no envrionment.
	//
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Associated metadata.
	// Default: - no metadata.
	//
	Metadata *map[string]*[]*MetadataEntry `field:"optional" json:"metadata" yaml:"metadata"`
	// The set of properties for this artifact (depends on type).
	// Default: - no properties.
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

