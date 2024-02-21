package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a Client VPN endpoint.
//
// A Client VPN endpoint is the resource you create and configure to enable and manage client VPN sessions. It is the destination endpoint at which all client VPN sessions are terminated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnEndpoint := awscdk.Aws_ec2.NewCfnClientVpnEndpoint(this, jsii.String("MyCfnClientVpnEndpoint"), &CfnClientVpnEndpointProps{
//   	AuthenticationOptions: []interface{}{
//   		&ClientAuthenticationRequestProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			ActiveDirectory: &DirectoryServiceAuthenticationRequestProperty{
//   				DirectoryId: jsii.String("directoryId"),
//   			},
//   			FederatedAuthentication: &FederatedAuthenticationRequestProperty{
//   				SamlProviderArn: jsii.String("samlProviderArn"),
//
//   				// the properties below are optional
//   				SelfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   			},
//   			MutualAuthentication: &CertificateAuthenticationRequestProperty{
//   				ClientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   			},
//   		},
//   	},
//   	ClientCidrBlock: jsii.String("clientCidrBlock"),
//   	ConnectionLogOptions: &ConnectionLogOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   		CloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   	},
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//
//   	// the properties below are optional
//   	ClientConnectOptions: &ClientConnectOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	ClientLoginBannerOptions: &ClientLoginBannerOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		BannerText: jsii.String("bannerText"),
//   	},
//   	Description: jsii.String("description"),
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SelfServicePortal: jsii.String("selfServicePortal"),
//   	SessionTimeoutHours: jsii.Number(123),
//   	SplitTunnel: jsii.Boolean(false),
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TransportProtocol: jsii.String("transportProtocol"),
//   	VpcId: jsii.String("vpcId"),
//   	VpnPort: jsii.Number(123),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html
//
type CfnClientVpnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions() interface{}
	SetAuthenticationOptions(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	ClientCidrBlock() *string
	SetClientCidrBlock(val *string)
	// The options for managing connection authorization for new client connections.
	ClientConnectOptions() interface{}
	SetClientConnectOptions(val interface{})
	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	ClientLoginBannerOptions() interface{}
	SetClientLoginBannerOptions(val interface{})
	// Information about the client connection logging options.
	ConnectionLogOptions() interface{}
	SetConnectionLogOptions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A brief description of the Client VPN endpoint.
	Description() *string
	SetDescription(val *string)
	// Information about the DNS servers to be used for DNS resolution.
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IDs of one or more security groups to apply to the target network.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	SelfServicePortal() *string
	SetSelfServicePortal(val *string)
	// The ARN of the server certificate.
	ServerCertificateArn() *string
	SetServerCertificateArn(val *string)
	// The maximum VPN session duration time in hours.
	SessionTimeoutHours() *float64
	SetSessionTimeoutHours(val *float64)
	// Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.
	SplitTunnel() interface{}
	SetSplitTunnel(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to apply to the Client VPN endpoint during creation.
	TagSpecifications() interface{}
	SetTagSpecifications(val interface{})
	// The transport protocol to be used by the VPN session.
	TransportProtocol() *string
	SetTransportProtocol(val *string)
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
	// The ID of the VPC to associate with the Client VPN endpoint.
	VpcId() *string
	SetVpcId(val *string)
	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	VpnPort() *float64
	SetVpnPort(val *float64)
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

// The jsii proxy struct for CfnClientVpnEndpoint
type jsiiProxy_CfnClientVpnEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClientVpnEndpoint) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) AuthenticationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ClientConnectOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ClientLoginBannerOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ConnectionLogOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SplitTunnel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) TagSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}


func NewCfnClientVpnEndpoint(scope constructs.Construct, id *string, props *CfnClientVpnEndpointProps) CfnClientVpnEndpoint {
	_init_.Initialize()

	if err := validateNewCfnClientVpnEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClientVpnEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnClientVpnEndpoint_Override(c CfnClientVpnEndpoint, scope constructs.Construct, id *string, props *CfnClientVpnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetAuthenticationOptions(val interface{}) {
	if err := j.validateSetAuthenticationOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetClientCidrBlock(val *string) {
	if err := j.validateSetClientCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCidrBlock",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetClientConnectOptions(val interface{}) {
	if err := j.validateSetClientConnectOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientConnectOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetClientLoginBannerOptions(val interface{}) {
	if err := j.validateSetClientLoginBannerOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientLoginBannerOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetConnectionLogOptions(val interface{}) {
	if err := j.validateSetConnectionLogOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionLogOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetDnsServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSelfServicePortal(val *string) {
	_jsii_.Set(
		j,
		"selfServicePortal",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetServerCertificateArn(val *string) {
	if err := j.validateSetServerCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSessionTimeoutHours(val *float64) {
	_jsii_.Set(
		j,
		"sessionTimeoutHours",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSplitTunnel(val interface{}) {
	if err := j.validateSetSplitTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitTunnel",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetTagSpecifications(val interface{}) {
	if err := j.validateSetTagSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetTransportProtocol(val *string) {
	_jsii_.Set(
		j,
		"transportProtocol",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetVpnPort(val *float64) {
	_jsii_.Set(
		j,
		"vpnPort",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnClientVpnEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpoint_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnClientVpnEndpoint_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpoint_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
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
func CfnClientVpnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClientVpnEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnClientVpnEndpoint) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnClientVpnEndpoint) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnClientVpnEndpoint) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

