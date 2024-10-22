package awstransfer

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Instantiates an auto-scaling virtual server based on the selected file transfer protocol in AWS .
//
// When you make updates to your file transfer protocol-enabled server or when you work with users, use the service-generated `ServerId` property that is assigned to the newly created server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServer := awscdk.Aws_transfer.NewCfnServer(this, jsii.String("MyCfnServer"), &CfnServerProps{
//   	Certificate: jsii.String("certificate"),
//   	Domain: jsii.String("domain"),
//   	EndpointDetails: &EndpointDetailsProperty{
//   		AddressAllocationIds: []*string{
//   			jsii.String("addressAllocationIds"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcEndpointId: jsii.String("vpcEndpointId"),
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	EndpointType: jsii.String("endpointType"),
//   	IdentityProviderDetails: &IdentityProviderDetailsProperty{
//   		DirectoryId: jsii.String("directoryId"),
//   		Function: jsii.String("function"),
//   		InvocationRole: jsii.String("invocationRole"),
//   		SftpAuthenticationMethods: jsii.String("sftpAuthenticationMethods"),
//   		Url: jsii.String("url"),
//   	},
//   	IdentityProviderType: jsii.String("identityProviderType"),
//   	LoggingRole: jsii.String("loggingRole"),
//   	PostAuthenticationLoginBanner: jsii.String("postAuthenticationLoginBanner"),
//   	PreAuthenticationLoginBanner: jsii.String("preAuthenticationLoginBanner"),
//   	ProtocolDetails: &ProtocolDetailsProperty{
//   		As2Transports: []*string{
//   			jsii.String("as2Transports"),
//   		},
//   		PassiveIp: jsii.String("passiveIp"),
//   		SetStatOption: jsii.String("setStatOption"),
//   		TlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   	},
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	S3StorageOptions: &S3StorageOptionsProperty{
//   		DirectoryListingOptimization: jsii.String("directoryListingOptimization"),
//   	},
//   	SecurityPolicyName: jsii.String("securityPolicyName"),
//   	StructuredLogDestinations: []*string{
//   		jsii.String("structuredLogDestinations"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowDetails: &WorkflowDetailsProperty{
//   		OnPartialUpload: []interface{}{
//   			&WorkflowDetailProperty{
//   				ExecutionRole: jsii.String("executionRole"),
//   				WorkflowId: jsii.String("workflowId"),
//   			},
//   		},
//   		OnUpload: []interface{}{
//   			&WorkflowDetailProperty{
//   				ExecutionRole: jsii.String("executionRole"),
//   				WorkflowId: jsii.String("workflowId"),
//   			},
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html
//
type CfnServer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The Amazon Resource Name associated with the server, in the form `arn:aws:transfer:region: *account-id* :server/ *server-id* /` .
	//
	// An example of a server ARN is: `arn:aws:transfer:us-east-1:123456789012:server/s-01234567890abcdef` .
	AttrArn() *string
	// The list of egress IP addresses of this server.
	//
	// These IP addresses are only relevant for servers that use the AS2 protocol. They are used for sending asynchronous MDNs. These IP addresses are assigned automatically when you create an AS2 server. Additionally, if you update an existing server and add the AS2 protocol, static IP addresses are assigned as well.
	AttrAs2ServiceManagedEgressIpAddresses() *[]*string
	// The service-assigned ID of the server that is created.
	//
	// An example `ServerId` is `s-01234567890abcdef` .
	AttrServerId() *string
	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	Certificate() *string
	SetCertificate(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies the domain of the storage system that is used for file transfers.
	Domain() *string
	SetDomain(val *string)
	// The virtual private cloud (VPC) endpoint settings that are configured for your server.
	EndpointDetails() interface{}
	SetEndpointDetails(val interface{})
	// The type of endpoint that you want your server to use.
	EndpointType() *string
	SetEndpointType(val *string)
	// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` , `AWS _LAMBDA` or `API_GATEWAY` .
	IdentityProviderDetails() interface{}
	SetIdentityProviderDetails(val interface{})
	// The mode of authentication for a server.
	IdentityProviderType() *string
	SetIdentityProviderType(val *string)
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFSevents.
	LoggingRole() *string
	SetLoggingRole(val *string)
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
	// Specifies a string to display when users connect to a server.
	//
	// This string is displayed after the user authenticates.
	PostAuthenticationLoginBanner() *string
	SetPostAuthenticationLoginBanner(val *string)
	// Specifies a string to display when users connect to a server.
	PreAuthenticationLoginBanner() *string
	SetPreAuthenticationLoginBanner(val *string)
	// The protocol settings that are configured for your server.
	ProtocolDetails() interface{}
	SetProtocolDetails(val interface{})
	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies whether or not performance for your Amazon S3 directories is optimized.
	//
	// This is disabled by default.
	S3StorageOptions() interface{}
	SetS3StorageOptions(val interface{})
	// Specifies the name of the security policy for the server.
	SecurityPolicyName() *string
	SetSecurityPolicyName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specifies the log groups to which your server logs are sent.
	StructuredLogDestinations() *[]*string
	SetStructuredLogDestinations(val *[]*string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Key-value pairs that can be used to group and search for servers.
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
	// Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow.
	WorkflowDetails() interface{}
	SetWorkflowDetails(val interface{})
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

// The jsii proxy struct for CfnServer
type jsiiProxy_CfnServer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnServer) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) AttrAs2ServiceManagedEgressIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrAs2ServiceManagedEgressIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) AttrServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) EndpointDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) IdentityProviderDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) IdentityProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) LoggingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) PostAuthenticationLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) PreAuthenticationLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) ProtocolDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) S3StorageOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3StorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) SecurityPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) StructuredLogDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuredLogDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServer) WorkflowDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowDetails",
		&returns,
	)
	return returns
}


func NewCfnServer(scope constructs.Construct, id *string, props *CfnServerProps) CfnServer {
	_init_.Initialize()

	if err := validateNewCfnServerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnServer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnServer_Override(c CfnServer, scope constructs.Construct, id *string, props *CfnServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnServer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServer)SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetEndpointDetails(val interface{}) {
	if err := j.validateSetEndpointDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointDetails",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetIdentityProviderDetails(val interface{}) {
	if err := j.validateSetIdentityProviderDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderDetails",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetIdentityProviderType(val *string) {
	_jsii_.Set(
		j,
		"identityProviderType",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetLoggingRole(val *string) {
	_jsii_.Set(
		j,
		"loggingRole",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetPostAuthenticationLoginBanner(val *string) {
	_jsii_.Set(
		j,
		"postAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetPreAuthenticationLoginBanner(val *string) {
	_jsii_.Set(
		j,
		"preAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetProtocolDetails(val interface{}) {
	if err := j.validateSetProtocolDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolDetails",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetS3StorageOptions(val interface{}) {
	if err := j.validateSetS3StorageOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3StorageOptions",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetSecurityPolicyName(val *string) {
	_jsii_.Set(
		j,
		"securityPolicyName",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetStructuredLogDestinations(val *[]*string) {
	_jsii_.Set(
		j,
		"structuredLogDestinations",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnServer)SetWorkflowDetails(val interface{}) {
	if err := j.validateSetWorkflowDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowDetails",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnServer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServer_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnServer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnServer_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServer_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnServer",
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
func CfnServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_transfer.CfnServer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServer) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServer) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServer) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServer) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServer) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServer) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServer) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServer) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnServer) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnServer) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServer) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServer) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServer) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnServer) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnServer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServer) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

