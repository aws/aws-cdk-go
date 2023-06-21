package awsamplifyuibuilder


// Properties for defining a `CfnComponent`.
//
// Example:
//
//
type CfnComponentProps struct {
	// The information to connect a component's properties to data at runtime.
	//
	// You can't specify `tags` as a valid property for `bindingProperties` .
	BindingProperties interface{} `field:"required" json:"bindingProperties" yaml:"bindingProperties"`
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	ComponentType *string `field:"required" json:"componentType" yaml:"componentType"`
	// The name of the component.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the component's properties that can be overriden in a customized instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides interface{} `field:"required" json:"overrides" yaml:"overrides"`
	// Describes the component's properties.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	Variants interface{} `field:"required" json:"variants" yaml:"variants"`
	// The unique ID of the Amplify app associated with the component.
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// A list of the component's `ComponentChild` instances.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	CollectionProperties interface{} `field:"optional" json:"collectionProperties" yaml:"collectionProperties"`
	// The name of the backend environment that is a part of the Amplify app.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Describes the events that can be raised on the component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// The schema version of the component when it was imported.
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// One or more key-value pairs to use when tagging the component.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

