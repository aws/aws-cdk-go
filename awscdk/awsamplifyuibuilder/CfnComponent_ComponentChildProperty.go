package awsamplifyuibuilder


// The `ComponentChild` property specifies a nested UI configuration within a parent `Component` .
//
// Example:
//
//
type CfnComponent_ComponentChildProperty struct {
	// The type of the child component.
	ComponentType *string `field:"required" json:"componentType" yaml:"componentType"`
	// The name of the child component.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the properties of the child component.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// The list of `ComponentChild` instances for this component.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// Describes the events that can be raised on the child component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
}

