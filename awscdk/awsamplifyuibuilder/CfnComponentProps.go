package awsamplifyuibuilder


// Properties for defining a `CfnComponent`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html
//
type CfnComponentProps struct {
	// The unique ID of the Amplify app associated with the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-appid
	//
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The information to connect a component's properties to data at runtime.
	//
	// You can't specify `tags` as a valid property for `bindingProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-bindingproperties
	//
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// A list of the component's `ComponentChild` instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-children
	//
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-collectionproperties
	//
	CollectionProperties interface{} `field:"optional" json:"collectionProperties" yaml:"collectionProperties"`
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-componenttype
	//
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// The name of the backend environment that is a part of the Amplify app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-environmentname
	//
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Describes the events that can be raised on the component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-events
	//
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// The name of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Describes the component's properties that can be overriden in a customized instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// Describes the component's properties.
	//
	// You can't specify `tags` as a valid property for `properties` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The schema version of the component when it was imported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-schemaversion
	//
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// The unique ID of the component in its original source system, such as Figma.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-sourceid
	//
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// One or more key-value pairs to use when tagging the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-component.html#cfn-amplifyuibuilder-component-variants
	//
	Variants interface{} `field:"optional" json:"variants" yaml:"variants"`
}

