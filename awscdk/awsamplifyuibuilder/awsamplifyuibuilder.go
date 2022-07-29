package awsamplifyuibuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AmplifyUIBuilder::Component`.
//
// The AWS::AmplifyUIBuilder::Component resource specifies a component within an Amplify app. A component is a user interface (UI) element that you can customize. Use `ComponentChild` to configure an instance of a `Component` . A `ComponentChild` instance inherits the configuration of the main `Component` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentChildProperty_ componentChildProperty
//   var componentPropertyProperty_ componentPropertyProperty
//   var events interface{}
//   var fields interface{}
//   var overrides interface{}
//   var predicateProperty_ predicateProperty
//   var properties interface{}
//   var variantValues interface{}
//
//   cfnComponent := awscdk.Aws_amplifyuibuilder.NewCfnComponent(this, jsii.String("MyCfnComponent"), &cfnComponentProps{
//   	bindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &ComponentBindingPropertiesValueProperty{
//   			"bindingProperties": &ComponentBindingPropertiesValuePropertiesProperty{
//   				"bucket": jsii.String("bucket"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"field": jsii.String("field"),
//   				"key": jsii.String("key"),
//   				"model": jsii.String("model"),
//   				"predicates": []interface{}{
//   					&predicateProperty{
//   						"and": []interface{}{
//   							predicateProperty_,
//   						},
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operator": jsii.String("operator"),
//   						"or": []interface{}{
//   							predicateProperty_,
//   						},
//   					},
//   				},
//   				"userAttribute": jsii.String("userAttribute"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	componentType: jsii.String("componentType"),
//   	name: jsii.String("name"),
//   	overrides: map[string]interface{}{
//   		"overridesKey": overrides,
//   	},
//   	properties: map[string]interface{}{
//   		"propertiesKey": &componentPropertyProperty{
//   			"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"bindings": bindings,
//   			"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &ComponentConditionPropertyProperty{
//   				"else": componentPropertyProperty_,
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operandType": jsii.String("operandType"),
//   				"operator": jsii.String("operator"),
//   				"property": jsii.String("property"),
//   				"then": componentPropertyProperty_,
//   			},
//   			"configured": jsii.Boolean(false),
//   			"defaultValue": jsii.String("defaultValue"),
//   			"event": jsii.String("event"),
//   			"importedValue": jsii.String("importedValue"),
//   			"model": jsii.String("model"),
//   			"property": jsii.String("property"),
//   			"type": jsii.String("type"),
//   			"userAttribute": jsii.String("userAttribute"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	variants: []interface{}{
//   		&componentVariantProperty{
//   			overrides: overrides,
//   			variantValues: variantValues,
//   		},
//   	},
//
//   	// the properties below are optional
//   	children: []interface{}{
//   		&componentChildProperty{
//   			componentType: jsii.String("componentType"),
//   			name: jsii.String("name"),
//   			properties: properties,
//
//   			// the properties below are optional
//   			children: []interface{}{
//   				componentChildProperty_,
//   			},
//   			events: events,
//   		},
//   	},
//   	collectionProperties: map[string]interface{}{
//   		"collectionPropertiesKey": &ComponentDataConfigurationProperty{
//   			"model": jsii.String("model"),
//
//   			// the properties below are optional
//   			"identifiers": []*string{
//   				jsii.String("identifiers"),
//   			},
//   			"predicate": &predicateProperty{
//   				"and": []interface{}{
//   					predicateProperty_,
//   				},
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operator": jsii.String("operator"),
//   				"or": []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   			"sort": []interface{}{
//   				&SortPropertyProperty{
//   					"direction": jsii.String("direction"),
//   					"field": jsii.String("field"),
//   				},
//   			},
//   		},
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &ComponentEventProperty{
//   			"action": jsii.String("action"),
//   			"parameters": &ActionParametersProperty{
//   				"anchor": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"fields": fields,
//   				"global": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"id": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"model": jsii.String("model"),
//   				"state": &MutationActionSetStateParameterProperty{
//   					"componentName": jsii.String("componentName"),
//   					"property": jsii.String("property"),
//   					"set": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": bindings,
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"target": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"type": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"url": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	schemaVersion: jsii.String("schemaVersion"),
//   	sourceId: jsii.String("sourceId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnComponent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique ID for the Amplify app.
	AttrAppId() *string
	// The name of the backend environment that is a part of the Amplify app.
	AttrEnvironmentName() *string
	// The unique ID of the component.
	AttrId() *string
	// The information to connect a component's properties to data at runtime.
	//
	// You can't specify `tags` as a valid property for `bindingProperties` .
	BindingProperties() interface{}
	SetBindingProperties(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A list of the component's `ComponentChild` instances.
	Children() interface{}
	SetChildren(val interface{})
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	CollectionProperties() interface{}
	SetCollectionProperties(val interface{})
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	ComponentType() *string
	SetComponentType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Describes the events that can be raised on the component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events() interface{}
	SetEvents(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the component.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Describes the component's properties that can be overriden in a customized instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides() interface{}
	SetOverrides(val interface{})
	// Describes the component's properties.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties() interface{}
	SetProperties(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The schema version of the component when it was imported.
	SchemaVersion() *string
	SetSchemaVersion(val *string)
	// The unique ID of the component in its original source system, such as Figma.
	SourceId() *string
	SetSourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// One or more key-value pairs to use when tagging the component.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	Variants() interface{}
	SetVariants(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComponent
type jsiiProxy_CfnComponent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponent) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) BindingProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bindingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Children() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"children",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CollectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) ComponentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Overrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Properties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Variants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variants",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent(scope awscdk.Construct, id *string, props *CfnComponentProps) CfnComponent {
	_init_.Initialize()

	j := jsiiProxy_CfnComponent{}

	_jsii_.Create(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent_Override(c CfnComponent, scope awscdk.Construct, id *string, props *CfnComponentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponent) SetBindingProperties(val interface{}) {
	_jsii_.Set(
		j,
		"bindingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetChildren(val interface{}) {
	_jsii_.Set(
		j,
		"children",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetCollectionProperties(val interface{}) {
	_jsii_.Set(
		j,
		"collectionProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetComponentType(val *string) {
	_jsii_.Set(
		j,
		"componentType",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetEvents(val interface{}) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetProperties(val interface{}) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetSchemaVersion(val *string) {
	_jsii_.Set(
		j,
		"schemaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetSourceId(val *string) {
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetVariants(val interface{}) {
	_jsii_.Set(
		j,
		"variants",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnComponent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnComponent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponent) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComponent) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponent) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComponent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComponent) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComponent) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponent) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnComponent) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponent) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `ActionParameters` property specifies the event action configuration for an element of a `Component` or `ComponentChild` .
//
// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components. `ActionParameters` defines the action that is performed when an event occurs on the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentPropertyProperty_ componentPropertyProperty
//   var fields interface{}
//
//   actionParametersProperty := &actionParametersProperty{
//   	anchor: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   	fields: fields,
//   	global: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   	id: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   	model: jsii.String("model"),
//   	state: &mutationActionSetStateParameterProperty{
//   		componentName: jsii.String("componentName"),
//   		property: jsii.String("property"),
//   		set: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	target: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   	type: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   	url: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnComponent_ActionParametersProperty struct {
	// The HTML anchor link to the location to open.
	//
	// Specify this value for a navigation action.
	Anchor interface{} `field:"optional" json:"anchor" yaml:"anchor"`
	// A dictionary of key-value pairs mapping Amplify Studio properties to fields in a data model.
	//
	// Use when the action performs an operation on an Amplify DataStore model.
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Specifies whether the user should be signed out globally.
	//
	// Specify this value for an auth sign out action.
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// The unique ID of the component that the `ActionParameters` apply to.
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// The name of the data model.
	//
	// Use when the action performs an operation on an Amplify DataStore model.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// A key-value pair that specifies the state property name and its initial value.
	State interface{} `field:"optional" json:"state" yaml:"state"`
	// The element within the same component to modify when the action occurs.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// The type of navigation action.
	//
	// Valid values are `url` and `anchor` . This value is required for a navigation action.
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// The URL to the location to open.
	//
	// Specify this value for a navigation action.
	Url interface{} `field:"optional" json:"url" yaml:"url"`
}

// The `ComponentBindingPropertiesValueProperties` property specifies the data binding configuration for a specific property using data stored in AWS .
//
// For AWS connected properties, you can bind a property to data stored in an Amazon S3 bucket, an Amplify DataStore model or an authenticated user attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   componentBindingPropertiesValuePropertiesProperty := &componentBindingPropertiesValuePropertiesProperty{
//   	bucket: jsii.String("bucket"),
//   	defaultValue: jsii.String("defaultValue"),
//   	field: jsii.String("field"),
//   	key: jsii.String("key"),
//   	model: jsii.String("model"),
//   	predicates: []interface{}{
//   		&predicateProperty{
//   			and: []interface{}{
//   				predicateProperty_,
//   			},
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operator: jsii.String("operator"),
//   			or: []interface{}{
//   				predicateProperty_,
//   			},
//   		},
//   	},
//   	userAttribute: jsii.String("userAttribute"),
//   }
//
type CfnComponent_ComponentBindingPropertiesValuePropertiesProperty struct {
	// An Amazon S3 bucket.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The default value to assign to the property.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The field to bind the data to.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The storage key for an Amazon S3 bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An Amplify DataStore model.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// A list of predicates for binding a component's properties to data.
	Predicates interface{} `field:"optional" json:"predicates" yaml:"predicates"`
	// An authenticated user attribute.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

// The `ComponentBindingPropertiesValue` property specifies the data binding configuration for a component at runtime.
//
// You can use `ComponentBindingPropertiesValue` to add exposed properties to a component to allow different values to be entered when a component is reused in different places in an app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   componentBindingPropertiesValueProperty := &componentBindingPropertiesValueProperty{
//   	bindingProperties: &componentBindingPropertiesValuePropertiesProperty{
//   		bucket: jsii.String("bucket"),
//   		defaultValue: jsii.String("defaultValue"),
//   		field: jsii.String("field"),
//   		key: jsii.String("key"),
//   		model: jsii.String("model"),
//   		predicates: []interface{}{
//   			&predicateProperty{
//   				and: []interface{}{
//   					predicateProperty_,
//   				},
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operator: jsii.String("operator"),
//   				or: []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   		},
//   		userAttribute: jsii.String("userAttribute"),
//   	},
//   	defaultValue: jsii.String("defaultValue"),
//   	type: jsii.String("type"),
//   }
//
type CfnComponent_ComponentBindingPropertiesValueProperty struct {
	// Describes the properties to customize with data at runtime.
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// The default value of the property.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The property type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The `ComponentChild` property specifies a nested UI configuration within a parent `Component` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentChildProperty_ componentChildProperty
//   var events interface{}
//   var properties interface{}
//
//   componentChildProperty := &componentChildProperty{
//   	componentType: jsii.String("componentType"),
//   	name: jsii.String("name"),
//   	properties: properties,
//
//   	// the properties below are optional
//   	children: []interface{}{
//   		&componentChildProperty{
//   			componentType: jsii.String("componentType"),
//   			name: jsii.String("name"),
//   			properties: properties,
//
//   			// the properties below are optional
//   			children: []interface{}{
//   				componentChildProperty_,
//   			},
//   			events: events,
//   		},
//   	},
//   	events: events,
//   }
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

// The `ComponentConditionProperty` property specifies a conditional expression for setting a component property.
//
// Use `ComponentConditionProperty` to set a property to different values conditionally, based on the value of another property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentConditionPropertyProperty := &componentConditionPropertyProperty{
//   	else: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   	field: jsii.String("field"),
//   	operand: jsii.String("operand"),
//   	operandType: jsii.String("operandType"),
//   	operator: jsii.String("operator"),
//   	property: jsii.String("property"),
//   	then: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnComponent_ComponentConditionPropertyProperty struct {
	// The value to assign to the property if the condition is not met.
	Else interface{} `field:"optional" json:"else" yaml:"else"`
	// The name of a field.
	//
	// Specify this when the property is a data model.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The value of the property to evaluate.
	Operand *string `field:"optional" json:"operand" yaml:"operand"`
	// The type of the property to evaluate.
	OperandType *string `field:"optional" json:"operandType" yaml:"operandType"`
	// The operator to use to perform the evaluation, such as `eq` to represent equals.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The name of the conditional property.
	Property *string `field:"optional" json:"property" yaml:"property"`
	// The value to assign to the property if the condition is met.
	Then interface{} `field:"optional" json:"then" yaml:"then"`
}

// The `ComponentDataConfiguration` property specifies the configuration for binding a component's properties to data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   componentDataConfigurationProperty := &componentDataConfigurationProperty{
//   	model: jsii.String("model"),
//
//   	// the properties below are optional
//   	identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	predicate: &predicateProperty{
//   		and: []interface{}{
//   			predicateProperty_,
//   		},
//   		field: jsii.String("field"),
//   		operand: jsii.String("operand"),
//   		operator: jsii.String("operator"),
//   		or: []interface{}{
//   			predicateProperty_,
//   		},
//   	},
//   	sort: []interface{}{
//   		&sortPropertyProperty{
//   			direction: jsii.String("direction"),
//   			field: jsii.String("field"),
//   		},
//   	},
//   }
//
type CfnComponent_ComponentDataConfigurationProperty struct {
	// The name of the data model to use to bind data to a component.
	Model *string `field:"required" json:"model" yaml:"model"`
	// A list of IDs to use to bind data to a component.
	//
	// Use this property to bind specifically chosen data, rather than data retrieved from a query.
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Represents the conditional logic to use when binding data to a component.
	//
	// Use this property to retrieve only a subset of the data in a collection.
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
	// Describes how to sort the component's properties.
	Sort interface{} `field:"optional" json:"sort" yaml:"sort"`
}

// The `ComponentEvent` property specifies the configuration of an event.
//
// You can bind an event and a corresponding action to a `Component` or a `ComponentChild` . A button click is an example of an event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentPropertyProperty_ componentPropertyProperty
//   var fields interface{}
//
//   componentEventProperty := &componentEventProperty{
//   	action: jsii.String("action"),
//   	parameters: &actionParametersProperty{
//   		anchor: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		fields: fields,
//   		global: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		id: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		model: jsii.String("model"),
//   		state: &mutationActionSetStateParameterProperty{
//   			componentName: jsii.String("componentName"),
//   			property: jsii.String("property"),
//   			set: &componentPropertyProperty{
//   				bindingProperties: &componentPropertyBindingPropertiesProperty{
//   					property: jsii.String("property"),
//
//   					// the properties below are optional
//   					field: jsii.String("field"),
//   				},
//   				bindings: bindings,
//   				collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   					property: jsii.String("property"),
//
//   					// the properties below are optional
//   					field: jsii.String("field"),
//   				},
//   				componentName: jsii.String("componentName"),
//   				concat: []interface{}{
//   					componentPropertyProperty_,
//   				},
//   				condition: &componentConditionPropertyProperty{
//   					else: componentPropertyProperty_,
//   					field: jsii.String("field"),
//   					operand: jsii.String("operand"),
//   					operandType: jsii.String("operandType"),
//   					operator: jsii.String("operator"),
//   					property: jsii.String("property"),
//   					then: componentPropertyProperty_,
//   				},
//   				configured: jsii.Boolean(false),
//   				defaultValue: jsii.String("defaultValue"),
//   				event: jsii.String("event"),
//   				importedValue: jsii.String("importedValue"),
//   				model: jsii.String("model"),
//   				property: jsii.String("property"),
//   				type: jsii.String("type"),
//   				userAttribute: jsii.String("userAttribute"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		target: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		type: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		url: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnComponent_ComponentEventProperty struct {
	// The action to perform when a specific event is raised.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Describes information about the action.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// The `ComponentPropertyBindingProperties` property specifies a component property to associate with a binding property.
//
// This enables exposed properties on the top level component to propagate data to the component's property values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentPropertyBindingPropertiesProperty := &componentPropertyBindingPropertiesProperty{
//   	property: jsii.String("property"),
//
//   	// the properties below are optional
//   	field: jsii.String("field"),
//   }
//
type CfnComponent_ComponentPropertyBindingPropertiesProperty struct {
	// The component property to bind to the data field.
	Property *string `field:"required" json:"property" yaml:"property"`
	// The data field to bind the property to.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

// The `ComponentProperty` property specifies the configuration for all of a component's properties.
//
// Use `ComponentProperty` to specify the values to render or bind by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentConditionPropertyProperty_ componentConditionPropertyProperty
//   var componentPropertyProperty_ componentPropertyProperty
//
//   componentPropertyProperty := &componentPropertyProperty{
//   	bindingProperties: &componentPropertyBindingPropertiesProperty{
//   		property: jsii.String("property"),
//
//   		// the properties below are optional
//   		field: jsii.String("field"),
//   	},
//   	bindings: bindings,
//   	collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   		property: jsii.String("property"),
//
//   		// the properties below are optional
//   		field: jsii.String("field"),
//   	},
//   	componentName: jsii.String("componentName"),
//   	concat: []interface{}{
//   		&componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: &componentConditionPropertyProperty{
//   				else: componentPropertyProperty_,
//   				field: jsii.String("field"),
//   				operand: jsii.String("operand"),
//   				operandType: jsii.String("operandType"),
//   				operator: jsii.String("operator"),
//   				property: jsii.String("property"),
//   				then: componentPropertyProperty_,
//   			},
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	condition: &componentConditionPropertyProperty{
//   		else: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: componentConditionPropertyProperty_,
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   		field: jsii.String("field"),
//   		operand: jsii.String("operand"),
//   		operandType: jsii.String("operandType"),
//   		operator: jsii.String("operator"),
//   		property: jsii.String("property"),
//   		then: &componentPropertyProperty{
//   			bindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			bindings: bindings,
//   			collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   				property: jsii.String("property"),
//
//   				// the properties below are optional
//   				field: jsii.String("field"),
//   			},
//   			componentName: jsii.String("componentName"),
//   			concat: []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			condition: componentConditionPropertyProperty_,
//   			configured: jsii.Boolean(false),
//   			defaultValue: jsii.String("defaultValue"),
//   			event: jsii.String("event"),
//   			importedValue: jsii.String("importedValue"),
//   			model: jsii.String("model"),
//   			property: jsii.String("property"),
//   			type: jsii.String("type"),
//   			userAttribute: jsii.String("userAttribute"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	configured: jsii.Boolean(false),
//   	defaultValue: jsii.String("defaultValue"),
//   	event: jsii.String("event"),
//   	importedValue: jsii.String("importedValue"),
//   	model: jsii.String("model"),
//   	property: jsii.String("property"),
//   	type: jsii.String("type"),
//   	userAttribute: jsii.String("userAttribute"),
//   	value: jsii.String("value"),
//   }
//
type CfnComponent_ComponentPropertyProperty struct {
	// The information to bind the component property to data at runtime.
	BindingProperties interface{} `field:"optional" json:"bindingProperties" yaml:"bindingProperties"`
	// The information to bind the component property to form data.
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
	// The information to bind the component property to data at runtime.
	//
	// Use this for collection components.
	CollectionBindingProperties interface{} `field:"optional" json:"collectionBindingProperties" yaml:"collectionBindingProperties"`
	// The name of the component that is affected by an event.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// A list of component properties to concatenate to create the value to assign to this component property.
	Concat interface{} `field:"optional" json:"concat" yaml:"concat"`
	// The conditional expression to use to assign a value to the component property.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Specifies whether the user configured the property in Amplify Studio after importing it.
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
	// The default value to assign to the component property.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// An event that occurs in your app.
	//
	// Use this for workflow data binding.
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The default value assigned to the property when the component is imported into an app.
	ImportedValue *string `field:"optional" json:"importedValue" yaml:"importedValue"`
	// The data model to use to assign a value to the component property.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The name of the component's property that is affected by an event.
	Property *string `field:"optional" json:"property" yaml:"property"`
	// The component type.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// An authenticated user attribute to use to assign a value to the component property.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
	// The value to assign to the component property.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The `ComponentVariant` property specifies the style configuration of a unique variation of a main component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var overrides interface{}
//   var variantValues interface{}
//
//   componentVariantProperty := &componentVariantProperty{
//   	overrides: overrides,
//   	variantValues: variantValues,
//   }
//
type CfnComponent_ComponentVariantProperty struct {
	// The properties of the component variant that can be overriden when customizing an instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// The combination of variants that comprise this variant.
	VariantValues interface{} `field:"optional" json:"variantValues" yaml:"variantValues"`
}

// The `MutationActionSetStateParameter` property specifies the state configuration when an action modifies a property of another element within the same component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentPropertyProperty_ componentPropertyProperty
//
//   mutationActionSetStateParameterProperty := &mutationActionSetStateParameterProperty{
//   	componentName: jsii.String("componentName"),
//   	property: jsii.String("property"),
//   	set: &componentPropertyProperty{
//   		bindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		bindings: bindings,
//   		collectionBindingProperties: &componentPropertyBindingPropertiesProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			field: jsii.String("field"),
//   		},
//   		componentName: jsii.String("componentName"),
//   		concat: []interface{}{
//   			componentPropertyProperty_,
//   		},
//   		condition: &componentConditionPropertyProperty{
//   			else: componentPropertyProperty_,
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operandType: jsii.String("operandType"),
//   			operator: jsii.String("operator"),
//   			property: jsii.String("property"),
//   			then: componentPropertyProperty_,
//   		},
//   		configured: jsii.Boolean(false),
//   		defaultValue: jsii.String("defaultValue"),
//   		event: jsii.String("event"),
//   		importedValue: jsii.String("importedValue"),
//   		model: jsii.String("model"),
//   		property: jsii.String("property"),
//   		type: jsii.String("type"),
//   		userAttribute: jsii.String("userAttribute"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnComponent_MutationActionSetStateParameterProperty struct {
	// The name of the component that is being modified.
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The name of the component property to apply the state configuration to.
	Property *string `field:"required" json:"property" yaml:"property"`
	// The state configuration to assign to the property.
	Set interface{} `field:"required" json:"set" yaml:"set"`
}

// The `Predicate` property specifies information for generating Amplify DataStore queries.
//
// Use `Predicate` to retrieve a subset of the data in a collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   predicateProperty := &predicateProperty{
//   	and: []interface{}{
//   		&predicateProperty{
//   			and: []interface{}{
//   				predicateProperty_,
//   			},
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operator: jsii.String("operator"),
//   			or: []interface{}{
//   				predicateProperty_,
//   			},
//   		},
//   	},
//   	field: jsii.String("field"),
//   	operand: jsii.String("operand"),
//   	operator: jsii.String("operator"),
//   	or: []interface{}{
//   		&predicateProperty{
//   			and: []interface{}{
//   				predicateProperty_,
//   			},
//   			field: jsii.String("field"),
//   			operand: jsii.String("operand"),
//   			operator: jsii.String("operator"),
//   			or: []interface{}{
//   				predicateProperty_,
//   			},
//   		},
//   	},
//   }
//
type CfnComponent_PredicateProperty struct {
	// A list of predicates to combine logically.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// The field to query.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The value to use when performing the evaluation.
	Operand *string `field:"optional" json:"operand" yaml:"operand"`
	// The operator to use to perform the evaluation.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// A list of predicates to combine logically.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
}

// The `SortProperty` property specifies how to sort the data that you bind to a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sortPropertyProperty := &sortPropertyProperty{
//   	direction: jsii.String("direction"),
//   	field: jsii.String("field"),
//   }
//
type CfnComponent_SortPropertyProperty struct {
	// The direction of the sort, either ascending or descending.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The field to perform the sort on.
	Field *string `field:"required" json:"field" yaml:"field"`
}

// Properties for defining a `CfnComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bindings interface{}
//   var componentChildProperty_ componentChildProperty
//   var componentPropertyProperty_ componentPropertyProperty
//   var events interface{}
//   var fields interface{}
//   var overrides interface{}
//   var predicateProperty_ predicateProperty
//   var properties interface{}
//   var variantValues interface{}
//
//   cfnComponentProps := &cfnComponentProps{
//   	bindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &ComponentBindingPropertiesValueProperty{
//   			"bindingProperties": &ComponentBindingPropertiesValuePropertiesProperty{
//   				"bucket": jsii.String("bucket"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"field": jsii.String("field"),
//   				"key": jsii.String("key"),
//   				"model": jsii.String("model"),
//   				"predicates": []interface{}{
//   					&predicateProperty{
//   						"and": []interface{}{
//   							predicateProperty_,
//   						},
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operator": jsii.String("operator"),
//   						"or": []interface{}{
//   							predicateProperty_,
//   						},
//   					},
//   				},
//   				"userAttribute": jsii.String("userAttribute"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	componentType: jsii.String("componentType"),
//   	name: jsii.String("name"),
//   	overrides: map[string]interface{}{
//   		"overridesKey": overrides,
//   	},
//   	properties: map[string]interface{}{
//   		"propertiesKey": &componentPropertyProperty{
//   			"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"bindings": bindings,
//   			"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &ComponentConditionPropertyProperty{
//   				"else": componentPropertyProperty_,
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operandType": jsii.String("operandType"),
//   				"operator": jsii.String("operator"),
//   				"property": jsii.String("property"),
//   				"then": componentPropertyProperty_,
//   			},
//   			"configured": jsii.Boolean(false),
//   			"defaultValue": jsii.String("defaultValue"),
//   			"event": jsii.String("event"),
//   			"importedValue": jsii.String("importedValue"),
//   			"model": jsii.String("model"),
//   			"property": jsii.String("property"),
//   			"type": jsii.String("type"),
//   			"userAttribute": jsii.String("userAttribute"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	variants: []interface{}{
//   		&componentVariantProperty{
//   			overrides: overrides,
//   			variantValues: variantValues,
//   		},
//   	},
//
//   	// the properties below are optional
//   	children: []interface{}{
//   		&componentChildProperty{
//   			componentType: jsii.String("componentType"),
//   			name: jsii.String("name"),
//   			properties: properties,
//
//   			// the properties below are optional
//   			children: []interface{}{
//   				componentChildProperty_,
//   			},
//   			events: events,
//   		},
//   	},
//   	collectionProperties: map[string]interface{}{
//   		"collectionPropertiesKey": &ComponentDataConfigurationProperty{
//   			"model": jsii.String("model"),
//
//   			// the properties below are optional
//   			"identifiers": []*string{
//   				jsii.String("identifiers"),
//   			},
//   			"predicate": &predicateProperty{
//   				"and": []interface{}{
//   					predicateProperty_,
//   				},
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operator": jsii.String("operator"),
//   				"or": []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   			"sort": []interface{}{
//   				&SortPropertyProperty{
//   					"direction": jsii.String("direction"),
//   					"field": jsii.String("field"),
//   				},
//   			},
//   		},
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &ComponentEventProperty{
//   			"action": jsii.String("action"),
//   			"parameters": &ActionParametersProperty{
//   				"anchor": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"fields": fields,
//   				"global": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"id": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"model": jsii.String("model"),
//   				"state": &MutationActionSetStateParameterProperty{
//   					"componentName": jsii.String("componentName"),
//   					"property": jsii.String("property"),
//   					"set": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": bindings,
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"target": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"type": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"url": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": bindings,
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	schemaVersion: jsii.String("schemaVersion"),
//   	sourceId: jsii.String("sourceId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
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
	// A list of the component's `ComponentChild` instances.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	CollectionProperties interface{} `field:"optional" json:"collectionProperties" yaml:"collectionProperties"`
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

// A CloudFormation `AWS::AmplifyUIBuilder::Theme`.
//
// The AWS::AmplifyUIBuilder::Theme resource specifies a theme within an Amplify app. A theme is a collection of style settings that apply globally to the components associated with the app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValuesProperty_ themeValuesProperty
//
//   cfnTheme := awscdk.Aws_amplifyuibuilder.NewCfnTheme(this, jsii.String("MyCfnTheme"), &cfnThemeProps{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	overrides: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnTheme interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique ID for the Amplify app associated with the theme.
	AttrAppId() *string
	// The time that the theme was created.
	AttrCreatedAt() *string
	// The name of the backend environment that is a part of the Amplify app.
	AttrEnvironmentName() *string
	// The ID for the theme.
	AttrId() *string
	// The time that the theme was modified.
	AttrModifiedAt() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the theme.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Describes the properties that can be overriden to customize a theme.
	Overrides() interface{}
	SetOverrides(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// One or more key-value pairs to use when tagging the theme.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A list of key-value pairs that defines the properties of the theme.
	Values() interface{}
	SetValues(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTheme
type jsiiProxy_CfnTheme struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTheme) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Overrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Values() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmplifyUIBuilder::Theme`.
func NewCfnTheme(scope awscdk.Construct, id *string, props *CfnThemeProps) CfnTheme {
	_init_.Initialize()

	j := jsiiProxy_CfnTheme{}

	_jsii_.Create(
		"monocdk.aws_amplifyuibuilder.CfnTheme",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmplifyUIBuilder::Theme`.
func NewCfnTheme_Override(c CfnTheme, scope awscdk.Construct, id *string, props *CfnThemeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplifyuibuilder.CfnTheme",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTheme) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetValues(val interface{}) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnTheme_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnTheme",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTheme_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnTheme",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTheme_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnTheme",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTheme_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplifyuibuilder.CfnTheme",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTheme) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTheme) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTheme) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTheme) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTheme) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTheme) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTheme) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTheme) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTheme) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTheme) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTheme) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTheme) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTheme) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTheme) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `ThemeValue` property specifies the configuration of a theme's properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValuesProperty_ themeValuesProperty
//
//   themeValueProperty := &themeValueProperty{
//   	children: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	value: jsii.String("value"),
//   }
//
type CfnTheme_ThemeValueProperty struct {
	// A list of key-value pairs that define the theme's properties.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The value of a theme property.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The `ThemeValues` property specifies key-value pair that defines a property of a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValueProperty_ themeValueProperty
//
//   themeValuesProperty := &themeValuesProperty{
//   	key: jsii.String("key"),
//   	value: &themeValueProperty{
//   		children: []interface{}{
//   			&themeValuesProperty{
//   				key: jsii.String("key"),
//   				value: themeValueProperty_,
//   			},
//   		},
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnTheme_ThemeValuesProperty struct {
	// The name of the property.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the property.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnTheme`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var themeValuesProperty_ themeValuesProperty
//
//   cfnThemeProps := &cfnThemeProps{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	overrides: []interface{}{
//   		&themeValuesProperty{
//   			key: jsii.String("key"),
//   			value: &themeValueProperty{
//   				children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnThemeProps struct {
	// The name of the theme.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of key-value pairs that defines the properties of the theme.
	Values interface{} `field:"required" json:"values" yaml:"values"`
	// Describes the properties that can be overriden to customize a theme.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// One or more key-value pairs to use when tagging the theme.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

