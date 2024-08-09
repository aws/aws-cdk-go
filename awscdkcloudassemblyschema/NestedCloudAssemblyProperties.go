package awscdkcloudassemblyschema


// Artifact properties for nested cloud assemblies.
type NestedCloudAssemblyProperties struct {
	// Relative path to the nested cloud assembly.
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
	// Display name for the cloud assembly.
	// Default: - The artifact ID.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

