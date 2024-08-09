package awscdkcloudassemblyschema


// Represents a missing piece of context.
type MissingContext struct {
	// The missing context key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// A set of provider-specific options.
	Props interface{} `field:"required" json:"props" yaml:"props"`
	// The provider from which we expect this context key to be obtained.
	Provider ContextProvider `field:"required" json:"provider" yaml:"provider"`
}

