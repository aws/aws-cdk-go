package awscdkcloudassemblyschema


// A manifest for a single artifact within the cloud assembly.
type ArtifactManifest struct {
	// The type of artifact.
	Type ArtifactType `field:"required" json:"type" yaml:"type"`
	// IDs of artifacts that must be deployed before this artifact.
	// Default: - no dependencies.
	//
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// A string that represents this artifact.
	//
	// Should only be used in user interfaces.
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

