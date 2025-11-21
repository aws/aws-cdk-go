package awsworkspacesweb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb"
	"github.com/aws/constructs-go/constructs/v10"
)

// This resource specifies user settings that can be associated with a web portal.
//
// Once associated with a web portal, user settings control how users can transfer data between a streaming session and the their local devices.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserSettings := awscdk.Aws_workspacesweb.NewCfnUserSettings(this, jsii.String("MyCfnUserSettings"), &CfnUserSettingsProps{
//   	CopyAllowed: jsii.String("copyAllowed"),
//   	DownloadAllowed: jsii.String("downloadAllowed"),
//   	PasteAllowed: jsii.String("pasteAllowed"),
//   	PrintAllowed: jsii.String("printAllowed"),
//   	UploadAllowed: jsii.String("uploadAllowed"),
//
//   	// the properties below are optional
//   	AdditionalEncryptionContext: map[string]*string{
//   		"additionalEncryptionContextKey": jsii.String("additionalEncryptionContext"),
//   	},
//   	CookieSynchronizationConfiguration: &CookieSynchronizationConfigurationProperty{
//   		Allowlist: []interface{}{
//   			&CookieSpecificationProperty{
//   				Domain: jsii.String("domain"),
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Blocklist: []interface{}{
//   			&CookieSpecificationProperty{
//   				Domain: jsii.String("domain"),
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   	},
//   	CustomerManagedKey: jsii.String("customerManagedKey"),
//   	DeepLinkAllowed: jsii.String("deepLinkAllowed"),
//   	DisconnectTimeoutInMinutes: jsii.Number(123),
//   	IdleDisconnectTimeoutInMinutes: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ToolbarConfiguration: &ToolbarConfigurationProperty{
//   		HiddenToolbarItems: []*string{
//   			jsii.String("hiddenToolbarItems"),
//   		},
//   		MaxDisplayResolution: jsii.String("maxDisplayResolution"),
//   		ToolbarType: jsii.String("toolbarType"),
//   		VisualMode: jsii.String("visualMode"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-usersettings.html
//
type CfnUserSettings interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsworkspacesweb.IUserSettingsRef
	awscdk.ITaggableV2
	// The additional encryption context of the user settings.
	AdditionalEncryptionContext() interface{}
	SetAdditionalEncryptionContext(val interface{})
	// A list of web portal ARNs that this user settings resource is associated with.
	AttrAssociatedPortalArns() *[]*string
	// The ARN of the user settings.
	AttrUserSettingsArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The configuration that specifies which cookies should be synchronized from the end user's local browser to the remote browser.
	CookieSynchronizationConfiguration() interface{}
	SetCookieSynchronizationConfiguration(val interface{})
	// Specifies whether the user can copy text from the streaming session to the local device.
	CopyAllowed() *string
	SetCopyAllowed(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The customer managed key used to encrypt sensitive information in the user settings.
	CustomerManagedKey() *string
	SetCustomerManagedKey(val *string)
	// Specifies whether the user can use deep links that open automatically when connecting to a session.
	DeepLinkAllowed() *string
	SetDeepLinkAllowed(val *string)
	// The amount of time that a streaming session remains active after users disconnect.
	DisconnectTimeoutInMinutes() *float64
	SetDisconnectTimeoutInMinutes(val *float64)
	// Specifies whether the user can download files from the streaming session to the local device.
	DownloadAllowed() *string
	SetDownloadAllowed(val *string)
	Env() *interfaces.ResourceEnvironment
	// The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the disconnect timeout interval begins.
	IdleDisconnectTimeoutInMinutes() *float64
	SetIdleDisconnectTimeoutInMinutes(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Specifies whether the user can paste text from the local device to the streaming session.
	PasteAllowed() *string
	SetPasteAllowed(val *string)
	// Specifies whether the user can print to the local device.
	PrintAllowed() *string
	SetPrintAllowed(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to add to the user settings resource.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The configuration of the toolbar.
	ToolbarConfiguration() interface{}
	SetToolbarConfiguration(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Specifies whether the user can upload files from the local device to the streaming session.
	UploadAllowed() *string
	SetUploadAllowed(val *string)
	// A reference to a UserSettings resource.
	UserSettingsRef() *interfacesawsworkspacesweb.UserSettingsReference
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnUserSettings
type jsiiProxy_CfnUserSettings struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsworkspaceswebIUserSettingsRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnUserSettings) AdditionalEncryptionContext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) AttrAssociatedPortalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrAssociatedPortalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) AttrUserSettingsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserSettingsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CookieSynchronizationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieSynchronizationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CopyAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) CustomerManagedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) DeepLinkAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deepLinkAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) DisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) DownloadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"downloadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) IdleDisconnectTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) PasteAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pasteAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) PrintAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) ToolbarConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolbarConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) UploadAllowed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserSettings) UserSettingsRef() *interfacesawsworkspacesweb.UserSettingsReference {
	var returns *interfacesawsworkspacesweb.UserSettingsReference
	_jsii_.Get(
		j,
		"userSettingsRef",
		&returns,
	)
	return returns
}


// Create a new `AWS::WorkSpacesWeb::UserSettings`.
func NewCfnUserSettings(scope constructs.Construct, id *string, props *CfnUserSettingsProps) CfnUserSettings {
	_init_.Initialize()

	if err := validateNewCfnUserSettingsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserSettings{}

	_jsii_.Create(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WorkSpacesWeb::UserSettings`.
func NewCfnUserSettings_Override(c CfnUserSettings, scope constructs.Construct, id *string, props *CfnUserSettingsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetAdditionalEncryptionContext(val interface{}) {
	if err := j.validateSetAdditionalEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetCookieSynchronizationConfiguration(val interface{}) {
	if err := j.validateSetCookieSynchronizationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieSynchronizationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetCopyAllowed(val *string) {
	if err := j.validateSetCopyAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAllowed",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetCustomerManagedKey(val *string) {
	_jsii_.Set(
		j,
		"customerManagedKey",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetDeepLinkAllowed(val *string) {
	_jsii_.Set(
		j,
		"deepLinkAllowed",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetDisconnectTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"disconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetDownloadAllowed(val *string) {
	if err := j.validateSetDownloadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downloadAllowed",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetIdleDisconnectTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetPasteAllowed(val *string) {
	if err := j.validateSetPasteAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pasteAllowed",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetPrintAllowed(val *string) {
	if err := j.validateSetPrintAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printAllowed",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetToolbarConfiguration(val interface{}) {
	if err := j.validateSetToolbarConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolbarConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserSettings)SetUploadAllowed(val *string) {
	if err := j.validateSetUploadAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadAllowed",
		val,
	)
}

func CfnUserSettings_ArnForUserSettings(resource interfacesawsworkspacesweb.IUserSettingsRef) *string {
	_init_.Initialize()

	if err := validateCfnUserSettings_ArnForUserSettingsParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		"arnForUserSettings",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnUserSettings_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserSettings_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnUserSettings_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserSettings_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnUserSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserSettings_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_workspacesweb.CfnUserSettings",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserSettings) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserSettings) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserSettings) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserSettings) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserSettings) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserSettings) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserSettings) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserSettings) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserSettings) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserSettings) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserSettings) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserSettings) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnUserSettings) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserSettings) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

