package awscdkcloudassemblyschema


// Information about the application's runtime components.
type RuntimeInfo struct {
	// The list of libraries loaded in the application, associated with their versions.
	Libraries *map[string]*string `field:"required" json:"libraries" yaml:"libraries"`
}

