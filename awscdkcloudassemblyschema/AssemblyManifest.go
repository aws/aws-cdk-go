package awscdkcloudassemblyschema


// A manifest which describes the cloud assembly.
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

