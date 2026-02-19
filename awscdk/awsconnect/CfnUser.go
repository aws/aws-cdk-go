package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a user account for an Amazon Connect instance.
//
// For information about how to create user accounts using the Amazon Connect console, see [Add Users](https://docs.aws.amazon.com/connect/latest/adminguide/user-management.html) in the *Amazon Connect Administrator Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUser := awscdk.Aws_connect.NewCfnUser(this, jsii.String("MyCfnUser"), &CfnUserProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	RoutingProfileArn: jsii.String("routingProfileArn"),
//   	SecurityProfileArns: []*string{
//   		jsii.String("securityProfileArns"),
//   	},
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	AfterContactWorkConfigs: []interface{}{
//   		&AfterContactWorkConfigPerChannelProperty{
//   			AfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   				AfterContactWorkTimeLimit: jsii.Number(123),
//   			},
//   			Channel: jsii.String("channel"),
//
//   			// the properties below are optional
//   			AgentFirstCallbackAfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   				AfterContactWorkTimeLimit: jsii.Number(123),
//   			},
//   		},
//   	},
//   	AutoAcceptConfigs: []interface{}{
//   		&AutoAcceptConfigProperty{
//   			AutoAccept: jsii.Boolean(false),
//   			Channel: jsii.String("channel"),
//
//   			// the properties below are optional
//   			AgentFirstCallbackAutoAccept: jsii.Boolean(false),
//   		},
//   	},
//   	DirectoryUserId: jsii.String("directoryUserId"),
//   	HierarchyGroupArn: jsii.String("hierarchyGroupArn"),
//   	IdentityInfo: &UserIdentityInfoProperty{
//   		Email: jsii.String("email"),
//   		FirstName: jsii.String("firstName"),
//   		LastName: jsii.String("lastName"),
//   		Mobile: jsii.String("mobile"),
//   		SecondaryEmail: jsii.String("secondaryEmail"),
//   	},
//   	Password: jsii.String("password"),
//   	PersistentConnectionConfigs: []interface{}{
//   		&PersistentConnectionConfigProperty{
//   			Channel: jsii.String("channel"),
//   			PersistentConnection: jsii.Boolean(false),
//   		},
//   	},
//   	PhoneConfig: &UserPhoneConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   		AutoAccept: jsii.Boolean(false),
//   		DeskPhoneNumber: jsii.String("deskPhoneNumber"),
//   		PersistentConnection: jsii.Boolean(false),
//   		PhoneType: jsii.String("phoneType"),
//   	},
//   	PhoneNumberConfigs: []interface{}{
//   		&PhoneNumberConfigProperty{
//   			Channel: jsii.String("channel"),
//   			PhoneType: jsii.String("phoneType"),
//
//   			// the properties below are optional
//   			PhoneNumber: jsii.String("phoneNumber"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserProficiencies: []interface{}{
//   		&UserProficiencyProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeValue: jsii.String("attributeValue"),
//   			Level: jsii.Number(123),
//   		},
//   	},
//   	VoiceEnhancementConfigs: []interface{}{
//   		&VoiceEnhancementConfigProperty{
//   			Channel: jsii.String("channel"),
//   			VoiceEnhancementMode: jsii.String("voiceEnhancementMode"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-user.html
//
type CfnUser interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsconnect.IUserRef
	awscdk.ITaggable
	// After Contact Work configurations of a user.
	AfterContactWorkConfigs() interface{}
	SetAfterContactWorkConfigs(val interface{})
	// The Amazon Resource Name (ARN) of the user.
	AttrUserArn() *string
	// Auto-accept configurations of a user.
	AutoAcceptConfigs() interface{}
	SetAutoAcceptConfigs(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The identifier of the user account in the directory used for identity management.
	DirectoryUserId() *string
	SetDirectoryUserId(val *string)
	Env() *interfaces.ResourceEnvironment
	// The Amazon Resource Name (ARN) of the user's hierarchy group.
	HierarchyGroupArn() *string
	SetHierarchyGroupArn(val *string)
	// Information about the user identity.
	IdentityInfo() interface{}
	SetIdentityInfo(val interface{})
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn() *string
	SetInstanceArn(val *string)
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
	// The user's password.
	Password() *string
	SetPassword(val *string)
	// Persistent Connection configurations of a user.
	PersistentConnectionConfigs() interface{}
	SetPersistentConnectionConfigs(val interface{})
	// Information about the phone configuration for the user.
	PhoneConfig() interface{}
	SetPhoneConfig(val interface{})
	// Phone Number configurations of a user.
	PhoneNumberConfigs() interface{}
	SetPhoneNumberConfigs(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the user's routing profile.
	RoutingProfileArn() *string
	SetRoutingProfileArn(val *string)
	// The Amazon Resource Name (ARN) of the user's security profile.
	SecurityProfileArns() *[]*string
	SetSecurityProfileArns(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	// The user name assigned to the user account.
	Username() *string
	SetUsername(val *string)
	// One or more predefined attributes assigned to a user, with a numeric value that indicates how their level of skill in a specified area.
	UserProficiencies() interface{}
	SetUserProficiencies(val interface{})
	// A reference to a User resource.
	UserRef() *interfacesawsconnect.UserReference
	// Voice Enhancement configurations of a user.
	VoiceEnhancementConfigs() interface{}
	SetVoiceEnhancementConfigs(val interface{})
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnUser
type jsiiProxy_CfnUser struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsconnectIUserRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnUser) AfterContactWorkConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"afterContactWorkConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) AttrUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) AutoAcceptConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAcceptConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) DirectoryUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) HierarchyGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) IdentityInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) PersistentConnectionConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistentConnectionConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) PhoneConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) PhoneNumberConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phoneNumberConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) RoutingProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) SecurityProfileArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityProfileArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) UserProficiencies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProficiencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) UserRef() *interfacesawsconnect.UserReference {
	var returns *interfacesawsconnect.UserReference
	_jsii_.Get(
		j,
		"userRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) VoiceEnhancementConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"voiceEnhancementConfigs",
		&returns,
	)
	return returns
}


// Create a new `AWS::Connect::User`.
func NewCfnUser(scope constructs.Construct, id *string, props *CfnUserProps) CfnUser {
	_init_.Initialize()

	if err := validateNewCfnUserParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUser{}

	_jsii_.Create(
		"aws-cdk-lib.aws_connect.CfnUser",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Connect::User`.
func NewCfnUser_Override(c CfnUser, scope constructs.Construct, id *string, props *CfnUserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_connect.CfnUser",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUser)SetAfterContactWorkConfigs(val interface{}) {
	if err := j.validateSetAfterContactWorkConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterContactWorkConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetAutoAcceptConfigs(val interface{}) {
	if err := j.validateSetAutoAcceptConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAcceptConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetDirectoryUserId(val *string) {
	_jsii_.Set(
		j,
		"directoryUserId",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetHierarchyGroupArn(val *string) {
	_jsii_.Set(
		j,
		"hierarchyGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetIdentityInfo(val interface{}) {
	if err := j.validateSetIdentityInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityInfo",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetInstanceArn(val *string) {
	if err := j.validateSetInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceArn",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetPersistentConnectionConfigs(val interface{}) {
	if err := j.validateSetPersistentConnectionConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistentConnectionConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetPhoneConfig(val interface{}) {
	if err := j.validateSetPhoneConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneConfig",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetPhoneNumberConfigs(val interface{}) {
	if err := j.validateSetPhoneNumberConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumberConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetRoutingProfileArn(val *string) {
	if err := j.validateSetRoutingProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingProfileArn",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetSecurityProfileArns(val *[]*string) {
	if err := j.validateSetSecurityProfileArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileArns",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetUserProficiencies(val interface{}) {
	if err := j.validateSetUserProficienciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userProficiencies",
		val,
	)
}

func (j *jsiiProxy_CfnUser)SetVoiceEnhancementConfigs(val interface{}) {
	if err := j.validateSetVoiceEnhancementConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"voiceEnhancementConfigs",
		val,
	)
}

func CfnUser_ArnForUser(resource interfacesawsconnect.IUserRef) *string {
	_init_.Initialize()

	if err := validateCfnUser_ArnForUserParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnUser",
		"arnForUser",
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
func CfnUser_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUser_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnUser",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnUser_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUser_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnUser",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnUser.
func CfnUser_IsCfnUser(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUser_IsCfnUserParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnUser",
		"isCfnUser",
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
func CfnUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connect.CfnUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUser_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_connect.CfnUser",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUser) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUser) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUser) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUser) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUser) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUser) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUser) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUser) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUser) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnUser) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnUser) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUser) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUser) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUser) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUser) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUser) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnUser) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnUser) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUser) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnUser) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

