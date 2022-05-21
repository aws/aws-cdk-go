package awstransfer

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Transfer::Server`.
//
// Instantiates an auto-scaling virtual server based on the selected file transfer protocol in AWS . When you make updates to your file transfer protocol-enabled server or when you work with users, use the service-generated `ServerId` property that is assigned to the newly created server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServer := awscdk.Aws_transfer.NewCfnServer(this, jsii.String("MyCfnServer"), &cfnServerProps{
//   	certificate: jsii.String("certificate"),
//   	domain: jsii.String("domain"),
//   	endpointDetails: &endpointDetailsProperty{
//   		addressAllocationIds: []*string{
//   			jsii.String("addressAllocationIds"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		vpcEndpointId: jsii.String("vpcEndpointId"),
//   		vpcId: jsii.String("vpcId"),
//   	},
//   	endpointType: jsii.String("endpointType"),
//   	identityProviderDetails: &identityProviderDetailsProperty{
//   		directoryId: jsii.String("directoryId"),
//   		function: jsii.String("function"),
//   		invocationRole: jsii.String("invocationRole"),
//   		url: jsii.String("url"),
//   	},
//   	identityProviderType: jsii.String("identityProviderType"),
//   	loggingRole: jsii.String("loggingRole"),
//   	postAuthenticationLoginBanner: jsii.String("postAuthenticationLoginBanner"),
//   	preAuthenticationLoginBanner: jsii.String("preAuthenticationLoginBanner"),
//   	protocolDetails: &protocolDetailsProperty{
//   		passiveIp: jsii.String("passiveIp"),
//   		tlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   	},
//   	protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	securityPolicyName: jsii.String("securityPolicyName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workflowDetails: &workflowDetailsProperty{
//   		onUpload: []interface{}{
//   			&workflowDetailProperty{
//   				executionRole: jsii.String("executionRole"),
//   				workflowId: jsii.String("workflowId"),
//   			},
//   		},
//   	},
//   })
//
type CfnServer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name associated with the server, in the form `arn:aws:transfer:region: *account-id* :server/ *server-id* /` .
	//
	// An example of a server ARN is: `arn:aws:transfer:us-east-1:123456789012:server/s-01234567890abcdef` .
	AttrArn() *string
	// The service-assigned ID of the server that is created.
	//
	// An example `ServerId` is `s-01234567890abcdef` .
	AttrServerId() *string
	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	//
	// Required when `Protocols` is set to `FTPS` .
	//
	// To request a new public certificate, see [Request a public certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html) in the *AWS Certificate Manager User Guide* .
	//
	// To import an existing certificate into ACM, see [Importing certificates into ACM](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *AWS Certificate Manager User Guide* .
	//
	// To request a private certificate to use FTPS through private IP addresses, see [Request a private certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html) in the *AWS Certificate Manager User Guide* .
	//
	// Certificates with the following cryptographic algorithms and key sizes are supported:
	//
	// - 2048-bit RSA (RSA_2048)
	// - 4096-bit RSA (RSA_4096)
	// - Elliptic Prime Curve 256 bit (EC_prime256v1)
	// - Elliptic Prime Curve 384 bit (EC_secp384r1)
	// - Elliptic Prime Curve 521 bit (EC_secp521r1)
	//
	// > The certificate must be a valid SSL/TLS X.509 version 3 certificate with FQDN or IP address specified and information about the issuer.
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
	//
	// When you host your endpoint within your VPC, you can make it accessible only to resources within your VPC, or you can attach Elastic IP addresses and make it accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
	EndpointDetails() interface{}
	SetEndpointDetails(val interface{})
	// The type of endpoint that you want your server to use.
	//
	// You can choose to make your server's endpoint publicly accessible (PUBLIC) or host it inside your VPC. With an endpoint that is hosted in a VPC, you can restrict access to your server and resources only within your VPC or choose to make it internet facing by attaching Elastic IP addresses directly to it.
	EndpointType() *string
	SetEndpointType(val *string)
	// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
	//
	// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
	IdentityProviderDetails() interface{}
	SetIdentityProviderDetails(val interface{})
	// Specifies the mode of authentication for a server.
	//
	// The default value is `SERVICE_MANAGED` , which allows you to store and access user credentials within the AWS Transfer Family service.
	//
	// Use `AWS_DIRECTORY_SERVICE` to provide access to Active Directory groups in AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment or in AWS using AD Connectors. This option also requires you to provide a Directory ID using the `IdentityProviderDetails` parameter.
	//
	// Use the `API_GATEWAY` value to integrate with an identity provider of your choosing. The `API_GATEWAY` setting requires you to provide an API Gateway endpoint URL to call for authentication using the `IdentityProviderDetails` parameter.
	//
	// Use the `AWS_LAMBDA` value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the `Function` parameter for the `IdentityProviderDetails` data type.
	IdentityProviderType() *string
	SetIdentityProviderType(val *string)
	// Specifies the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFS events.
	//
	// When set, user activity can be viewed in your CloudWatch logs.
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
	// Specify a string to display when users connect to a server. This string is displayed after the user authenticates.
	//
	// > The SFTP protocol does not support post-authentication display banners.
	PostAuthenticationLoginBanner() *string
	SetPostAuthenticationLoginBanner(val *string)
	// Specify a string to display when users connect to a server.
	//
	// This string is displayed before the user authenticates. For example, the following banner displays details about using the system.
	//
	// `This system is for the use of authorized users only. Individuals using this computer system without authority, or in excess of their authority, are subject to having all of their activities on this system monitored and recorded by system personnel.`
	PreAuthenticationLoginBanner() *string
	SetPreAuthenticationLoginBanner(val *string)
	// The protocol settings that are configured for your server.
	//
	// Use the `PassiveIp` parameter to indicate passive mode (for FTP and FTPS protocols). Enter a single dotted-quad IPv4 address, such as the external IP address of a firewall, router, or load balancer.
	//
	// Use the `TlsSessionResumptionMode` parameter to determine whether or not your Transfer server resumes recent, negotiated sessions through a unique session ID.
	ProtocolDetails() interface{}
	SetProtocolDetails(val interface{})
	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
	//
	// The available protocols are:
	//
	// - `SFTP` (Secure Shell (SSH) File Transfer Protocol): File transfer over SSH
	// - `FTPS` (File Transfer Protocol Secure): File transfer with TLS encryption
	// - `FTP` (File Transfer Protocol): Unencrypted file transfer
	//
	// > If you select `FTPS` , you must choose a certificate stored in AWS Certificate Manager (ACM) which is used to identify your server when clients connect to it over FTPS.
	// >
	// > If `Protocol` includes either `FTP` or `FTPS` , then the `EndpointType` must be `VPC` and the `IdentityProviderType` must be `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
	// >
	// > If `Protocol` includes `FTP` , then `AddressAllocationIds` cannot be associated.
	// >
	// > If `Protocol` is set only to `SFTP` , the `EndpointType` can be set to `PUBLIC` and the `IdentityProviderType` can be set to `SERVICE_MANAGED` .
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies the name of the security policy that is attached to the server.
	SecurityPolicyName() *string
	SetSecurityPolicyName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Key-value pairs that can be used to group and search for servers.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Specifies the workflow ID for the workflow to assign and the execution role used for executing the workflow.
	WorkflowDetails() interface{}
	SetWorkflowDetails(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

func (j *jsiiProxy_CfnServer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
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

func (j *jsiiProxy_CfnServer) WorkflowDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowDetails",
		&returns,
	)
	return returns
}


// Create a new `AWS::Transfer::Server`.
func NewCfnServer(scope constructs.Construct, id *string, props *CfnServerProps) CfnServer {
	_init_.Initialize()

	j := jsiiProxy_CfnServer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnServer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Transfer::Server`.
func NewCfnServer_Override(c CfnServer, scope constructs.Construct, id *string, props *CfnServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnServer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServer) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetEndpointDetails(val interface{}) {
	_jsii_.Set(
		j,
		"endpointDetails",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetIdentityProviderDetails(val interface{}) {
	_jsii_.Set(
		j,
		"identityProviderDetails",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetIdentityProviderType(val *string) {
	_jsii_.Set(
		j,
		"identityProviderType",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetLoggingRole(val *string) {
	_jsii_.Set(
		j,
		"loggingRole",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetPostAuthenticationLoginBanner(val *string) {
	_jsii_.Set(
		j,
		"postAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetPreAuthenticationLoginBanner(val *string) {
	_jsii_.Set(
		j,
		"preAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetProtocolDetails(val interface{}) {
	_jsii_.Set(
		j,
		"protocolDetails",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetSecurityPolicyName(val *string) {
	_jsii_.Set(
		j,
		"securityPolicyName",
		val,
	)
}

func (j *jsiiProxy_CfnServer) SetWorkflowDetails(val interface{}) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnServer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnServer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnServer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServer) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The virtual private cloud (VPC) endpoint settings that are configured for your server.
//
// When you host your endpoint within your VPC, you can make it accessible only to resources within your VPC, or you can attach Elastic IP addresses and make it accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointDetailsProperty := &endpointDetailsProperty{
//   	addressAllocationIds: []*string{
//   		jsii.String("addressAllocationIds"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnServer_EndpointDetailsProperty struct {
	// A list of address allocation IDs that are required to attach an Elastic IP address to your server's endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` and it is only valid in the `UpdateServer` API.
	AddressAllocationIds *[]*string `field:"optional" json:"addressAllocationIds" yaml:"addressAllocationIds"`
	// A list of security groups IDs that are available to attach to your server's endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	// >
	// > You can edit the `SecurityGroupIds` property in the [UpdateServer](https://docs.aws.amazon.com/transfer/latest/userguide/API_UpdateServer.html) API only if you are changing the `EndpointType` from `PUBLIC` or `VPC_ENDPOINT` to `VPC` . To change security groups associated with your server's VPC endpoint after creation, use the Amazon EC2 [ModifyVpcEndpoint](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyVpcEndpoint.html) API.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs that are required to host your server endpoint in your VPC.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC endpoint.
	//
	// > This property can only be set when `EndpointType` is set to `VPC_ENDPOINT` .
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The VPC ID of the virtual private cloud in which the server's endpoint will be hosted.
	//
	// > This property can only be set when `EndpointType` is set to `VPC` .
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
//
// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderDetailsProperty := &identityProviderDetailsProperty{
//   	directoryId: jsii.String("directoryId"),
//   	function: jsii.String("function"),
//   	invocationRole: jsii.String("invocationRole"),
//   	url: jsii.String("url"),
//   }
//
type CfnServer_IdentityProviderDetailsProperty struct {
	// The identifier of the AWS Directory Service directory that you want to stop sharing.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The ARN for a lambda function to use for the Identity provider.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Provides the type of `InvocationRole` used to authenticate the user account.
	InvocationRole *string `field:"optional" json:"invocationRole" yaml:"invocationRole"`
	// Provides the location of the service endpoint used to authenticate users.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

// Protocol settings that are configured for your server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protocolDetailsProperty := &protocolDetailsProperty{
//   	passiveIp: jsii.String("passiveIp"),
//   	tlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   }
//
type CfnServer_ProtocolDetailsProperty struct {
	// Indicates passive mode, for FTP and FTPS protocols.
	//
	// Enter a single dotted-quad IPv4 address, such as the external IP address of a firewall, router, or load balancer.
	PassiveIp *string `field:"optional" json:"passiveIp" yaml:"passiveIp"`
	// A property used with Transfer servers that use the FTPS protocol.
	//
	// TLS Session Resumption provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. `TlsSessionResumptionMode` determines whether or not the server resumes recent, negotiated sessions through a unique session ID. This property is available during `CreateServer` and `UpdateServer` calls. If a `TlsSessionResumptionMode` value is not specified during CreateServer, it is set to `ENFORCED` by default.
	//
	// - `DISABLED` : the server does not process TLS session resumption client requests and creates a new TLS session for each request.
	// - `ENABLED` : the server processes and accepts clients that are performing TLS session resumption. The server doesn't reject client data connections that do not perform the TLS session resumption client processing.
	// - `ENFORCED` : the server processes and accepts clients that are performing TLS session resumption. The server rejects client data connections that do not perform the TLS session resumption client processing. Before you set the value to `ENFORCED` , test your clients.
	//
	// > Not all FTPS clients perform TLS session resumption. So, if you choose to enforce TLS session resumption, you prevent any connections from FTPS clients that don't perform the protocol negotiation. To determine whether or not you can use the `ENFORCED` value, you need to test your clients.
	TlsSessionResumptionMode *string `field:"optional" json:"tlsSessionResumptionMode" yaml:"tlsSessionResumptionMode"`
}

// Specifies the workflow ID for the workflow to assign and the execution role used for executing the workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowDetailProperty := &workflowDetailProperty{
//   	executionRole: jsii.String("executionRole"),
//   	workflowId: jsii.String("workflowId"),
//   }
//
type CfnServer_WorkflowDetailProperty struct {
	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// A unique identifier for the workflow.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

// Container for the `WorkflowDetail` data type.
//
// It is used by actions that trigger a workflow to begin execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowDetailsProperty := &workflowDetailsProperty{
//   	onUpload: []interface{}{
//   		&workflowDetailProperty{
//   			executionRole: jsii.String("executionRole"),
//   			workflowId: jsii.String("workflowId"),
//   		},
//   	},
//   }
//
type CfnServer_WorkflowDetailsProperty struct {
	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded.
	//
	// To remove an associated workflow from a server, you can provide an empty `OnUpload` object, as in the following example.
	//
	// `aws transfer update-server --server-id s-01234567890abcdef --workflow-details '{"OnUpload":[]}'`.
	OnUpload interface{} `field:"required" json:"onUpload" yaml:"onUpload"`
}

// Properties for defining a `CfnServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerProps := &cfnServerProps{
//   	certificate: jsii.String("certificate"),
//   	domain: jsii.String("domain"),
//   	endpointDetails: &endpointDetailsProperty{
//   		addressAllocationIds: []*string{
//   			jsii.String("addressAllocationIds"),
//   		},
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		vpcEndpointId: jsii.String("vpcEndpointId"),
//   		vpcId: jsii.String("vpcId"),
//   	},
//   	endpointType: jsii.String("endpointType"),
//   	identityProviderDetails: &identityProviderDetailsProperty{
//   		directoryId: jsii.String("directoryId"),
//   		function: jsii.String("function"),
//   		invocationRole: jsii.String("invocationRole"),
//   		url: jsii.String("url"),
//   	},
//   	identityProviderType: jsii.String("identityProviderType"),
//   	loggingRole: jsii.String("loggingRole"),
//   	postAuthenticationLoginBanner: jsii.String("postAuthenticationLoginBanner"),
//   	preAuthenticationLoginBanner: jsii.String("preAuthenticationLoginBanner"),
//   	protocolDetails: &protocolDetailsProperty{
//   		passiveIp: jsii.String("passiveIp"),
//   		tlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   	},
//   	protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	securityPolicyName: jsii.String("securityPolicyName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workflowDetails: &workflowDetailsProperty{
//   		onUpload: []interface{}{
//   			&workflowDetailProperty{
//   				executionRole: jsii.String("executionRole"),
//   				workflowId: jsii.String("workflowId"),
//   			},
//   		},
//   	},
//   }
//
type CfnServerProps struct {
	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	//
	// Required when `Protocols` is set to `FTPS` .
	//
	// To request a new public certificate, see [Request a public certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html) in the *AWS Certificate Manager User Guide* .
	//
	// To import an existing certificate into ACM, see [Importing certificates into ACM](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *AWS Certificate Manager User Guide* .
	//
	// To request a private certificate to use FTPS through private IP addresses, see [Request a private certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html) in the *AWS Certificate Manager User Guide* .
	//
	// Certificates with the following cryptographic algorithms and key sizes are supported:
	//
	// - 2048-bit RSA (RSA_2048)
	// - 4096-bit RSA (RSA_4096)
	// - Elliptic Prime Curve 256 bit (EC_prime256v1)
	// - Elliptic Prime Curve 384 bit (EC_secp384r1)
	// - Elliptic Prime Curve 521 bit (EC_secp521r1)
	//
	// > The certificate must be a valid SSL/TLS X.509 version 3 certificate with FQDN or IP address specified and information about the issuer.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Specifies the domain of the storage system that is used for file transfers.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The virtual private cloud (VPC) endpoint settings that are configured for your server.
	//
	// When you host your endpoint within your VPC, you can make it accessible only to resources within your VPC, or you can attach Elastic IP addresses and make it accessible to clients over the internet. Your VPC's default security groups are automatically assigned to your endpoint.
	EndpointDetails interface{} `field:"optional" json:"endpointDetails" yaml:"endpointDetails"`
	// The type of endpoint that you want your server to use.
	//
	// You can choose to make your server's endpoint publicly accessible (PUBLIC) or host it inside your VPC. With an endpoint that is hosted in a VPC, you can restrict access to your server and resources only within your VPC or choose to make it internet facing by attaching Elastic IP addresses directly to it.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
	//
	// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
	IdentityProviderDetails interface{} `field:"optional" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// Specifies the mode of authentication for a server.
	//
	// The default value is `SERVICE_MANAGED` , which allows you to store and access user credentials within the AWS Transfer Family service.
	//
	// Use `AWS_DIRECTORY_SERVICE` to provide access to Active Directory groups in AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment or in AWS using AD Connectors. This option also requires you to provide a Directory ID using the `IdentityProviderDetails` parameter.
	//
	// Use the `API_GATEWAY` value to integrate with an identity provider of your choosing. The `API_GATEWAY` setting requires you to provide an API Gateway endpoint URL to call for authentication using the `IdentityProviderDetails` parameter.
	//
	// Use the `AWS_LAMBDA` value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the `Function` parameter for the `IdentityProviderDetails` data type.
	IdentityProviderType *string `field:"optional" json:"identityProviderType" yaml:"identityProviderType"`
	// Specifies the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a server to turn on Amazon CloudWatch logging for Amazon S3 or Amazon EFS events.
	//
	// When set, user activity can be viewed in your CloudWatch logs.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Specify a string to display when users connect to a server. This string is displayed after the user authenticates.
	//
	// > The SFTP protocol does not support post-authentication display banners.
	PostAuthenticationLoginBanner *string `field:"optional" json:"postAuthenticationLoginBanner" yaml:"postAuthenticationLoginBanner"`
	// Specify a string to display when users connect to a server.
	//
	// This string is displayed before the user authenticates. For example, the following banner displays details about using the system.
	//
	// `This system is for the use of authorized users only. Individuals using this computer system without authority, or in excess of their authority, are subject to having all of their activities on this system monitored and recorded by system personnel.`
	PreAuthenticationLoginBanner *string `field:"optional" json:"preAuthenticationLoginBanner" yaml:"preAuthenticationLoginBanner"`
	// The protocol settings that are configured for your server.
	//
	// Use the `PassiveIp` parameter to indicate passive mode (for FTP and FTPS protocols). Enter a single dotted-quad IPv4 address, such as the external IP address of a firewall, router, or load balancer.
	//
	// Use the `TlsSessionResumptionMode` parameter to determine whether or not your Transfer server resumes recent, negotiated sessions through a unique session ID.
	ProtocolDetails interface{} `field:"optional" json:"protocolDetails" yaml:"protocolDetails"`
	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
	//
	// The available protocols are:
	//
	// - `SFTP` (Secure Shell (SSH) File Transfer Protocol): File transfer over SSH
	// - `FTPS` (File Transfer Protocol Secure): File transfer with TLS encryption
	// - `FTP` (File Transfer Protocol): Unencrypted file transfer
	//
	// > If you select `FTPS` , you must choose a certificate stored in AWS Certificate Manager (ACM) which is used to identify your server when clients connect to it over FTPS.
	// >
	// > If `Protocol` includes either `FTP` or `FTPS` , then the `EndpointType` must be `VPC` and the `IdentityProviderType` must be `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
	// >
	// > If `Protocol` includes `FTP` , then `AddressAllocationIds` cannot be associated.
	// >
	// > If `Protocol` is set only to `SFTP` , the `EndpointType` can be set to `PUBLIC` and the `IdentityProviderType` can be set to `SERVICE_MANAGED` .
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Specifies the name of the security policy that is attached to the server.
	SecurityPolicyName *string `field:"optional" json:"securityPolicyName" yaml:"securityPolicyName"`
	// Key-value pairs that can be used to group and search for servers.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the workflow ID for the workflow to assign and the execution role used for executing the workflow.
	WorkflowDetails interface{} `field:"optional" json:"workflowDetails" yaml:"workflowDetails"`
}

// A CloudFormation `AWS::Transfer::User`.
//
// The `AWS::Transfer::User` resource creates a user and associates them with an existing server. You can only create and associate users with servers that have the `IdentityProviderType` set to `SERVICE_MANAGED` . Using parameters for `CreateUser` , you can specify the user name, set the home directory, store the user's public key, and assign the user's AWS Identity and Access Management (IAM) role. You can also optionally add a session policy, and assign metadata with tags that can be used to group and search for users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUser := awscdk.Aws_transfer.NewCfnUser(this, jsii.String("MyCfnUser"), &cfnUserProps{
//   	role: jsii.String("role"),
//   	serverId: jsii.String("serverId"),
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	homeDirectory: jsii.String("homeDirectory"),
//   	homeDirectoryMappings: []interface{}{
//   		&homeDirectoryMapEntryProperty{
//   			entry: jsii.String("entry"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   	homeDirectoryType: jsii.String("homeDirectoryType"),
//   	policy: jsii.String("policy"),
//   	posixProfile: &posixProfileProperty{
//   		gid: jsii.Number(123),
//   		uid: jsii.Number(123),
//
//   		// the properties below are optional
//   		secondaryGids: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	sshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnUser interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name associated with the user, in the form `arn:aws:transfer:region: *account-id* :user/ *server-id* / *username*` .
	//
	// An example of a user ARN is: `arn:aws:transfer:us-east-1:123456789012:user/user1` .
	AttrArn() *string
	// The ID of the server to which the user is attached.
	//
	// An example `ServerId` is `s-01234567890abcdef` .
	AttrServerId() *string
	// A unique string that identifies a user account associated with a server.
	//
	// An example `UserName` is `transfer-user-1` .
	AttrUserName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The landing directory (folder) for a user when they log in to the server using the client.
	//
	// A `HomeDirectory` example is `/bucket_name/home/mydirectory` .
	HomeDirectory() *string
	SetHomeDirectory(val *string)
	// Logical directory mappings that specify what Amazon S3 paths and keys should be visible to your user and how you want to make them visible.
	//
	// You will need to specify the " `Entry` " and " `Target` " pair, where `Entry` shows how the path is made visible and `Target` is the actual Amazon S3 path. If you only specify a target, it will be displayed as is. You will need to also make sure that your IAM role provides access to paths in `Target` . The following is an example.
	//
	// `'[ { "Entry": "/", "Target": "/bucket3/customized-reports/" } ]'`
	//
	// In most cases, you can use this value instead of the session policy to lock your user down to the designated home directory ("chroot"). To do this, you can set `Entry` to '/' and set `Target` to the HomeDirectory parameter value.
	//
	// > If the target of a logical directory entry does not exist in Amazon S3, the entry will be ignored. As a workaround, you can use the Amazon S3 API to create 0 byte objects as place holders for your directory. If using the CLI, use the `s3api` call instead of `s3` so you can use the put-object operation. For example, you use the following: `AWS s3api put-object --bucket bucketname --key path/to/folder/` . Make sure that the end of the key name ends in a '/' for it to be considered a folder.
	HomeDirectoryMappings() interface{}
	SetHomeDirectoryMappings(val interface{})
	// The type of landing directory (folder) you want your users' home directory to be when they log into the server.
	//
	// If you set it to `PATH` , the user will see the absolute Amazon S3 bucket or EFS paths as is in their file transfer protocol clients. If you set it `LOGICAL` , you need to provide mappings in the `HomeDirectoryMappings` for how you want to make Amazon S3 or EFS paths visible to your users.
	HomeDirectoryType() *string
	SetHomeDirectoryType(val *string)
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
	// A session policy for your user so you can use the same IAM role across multiple users.
	//
	// This policy restricts user access to portions of their Amazon S3 bucket. Variables that you can use inside this policy include `${Transfer:UserName}` , `${Transfer:HomeDirectory}` , and `${Transfer:HomeBucket}` .
	//
	// > For session policies, AWS Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the `Policy` argument.
	// >
	// > For an example of a session policy, see [Example session policy](https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html) .
	// >
	// > For more information, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the *AWS Security Token Service API Reference* .
	Policy() *string
	SetPolicy(val *string)
	// Specifies the full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon Elastic File System (Amazon EFS) file systems.
	//
	// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	PosixProfile() interface{}
	SetPosixProfile(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies the Amazon Resource Name (ARN) of the IAM role that controls your users' access to your Amazon S3 bucket or EFS file system.
	//
	// The policies attached to this role determine the level of access that you want to provide your users when transferring files into and out of your Amazon S3 bucket or EFS file system. The IAM role should also contain a trust relationship that allows the server to access your resources when servicing your users' transfer requests.
	Role() *string
	SetRole(val *string)
	// A system-assigned unique identifier for a server instance.
	//
	// This is the specific server that you added your user to.
	ServerId() *string
	SetServerId(val *string)
	// Specifies the public key portion of the Secure Shell (SSH) keys stored for the described user.
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Key-value pairs that can be used to group and search for users.
	//
	// Tags are metadata attached to users for any purpose.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// A unique string that identifies a user and is associated with a `ServerId` .
	//
	// This user name must be a minimum of 3 and a maximum of 100 characters long. The following are valid characters: a-z, A-Z, 0-9, underscore '_', hyphen '-', period '.', and at sign '@'. The user name can't start with a hyphen, period, or at sign.
	UserName() *string
	SetUserName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnUser
type jsiiProxy_CfnUser struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUser) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) AttrServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) AttrUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserName",
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

func (j *jsiiProxy_CfnUser) HomeDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) HomeDirectoryMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"homeDirectoryMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) HomeDirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryType",
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

func (j *jsiiProxy_CfnUser) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) PosixProfile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"posixProfile",
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

func (j *jsiiProxy_CfnUser) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
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

func (j *jsiiProxy_CfnUser) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}


// Create a new `AWS::Transfer::User`.
func NewCfnUser(scope constructs.Construct, id *string, props *CfnUserProps) CfnUser {
	_init_.Initialize()

	j := jsiiProxy_CfnUser{}

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnUser",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Transfer::User`.
func NewCfnUser_Override(c CfnUser, scope constructs.Construct, id *string, props *CfnUserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnUser",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUser) SetHomeDirectory(val *string) {
	_jsii_.Set(
		j,
		"homeDirectory",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetHomeDirectoryMappings(val interface{}) {
	_jsii_.Set(
		j,
		"homeDirectoryMappings",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetHomeDirectoryType(val *string) {
	_jsii_.Set(
		j,
		"homeDirectoryType",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetPosixProfile(val interface{}) {
	_jsii_.Set(
		j,
		"posixProfile",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetServerId(val *string) {
	_jsii_.Set(
		j,
		"serverId",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetSshPublicKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_CfnUser) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnUser_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnUser",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnUser_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnUser",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnUser",
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
		"aws-cdk-lib.aws_transfer.CfnUser",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUser) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUser) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUser) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUser) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUser) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUser) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUser) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUser) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUser) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUser) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUser) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents an object that contains entries and targets for `HomeDirectoryMappings` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   homeDirectoryMapEntryProperty := &homeDirectoryMapEntryProperty{
//   	entry: jsii.String("entry"),
//   	target: jsii.String("target"),
//   }
//
type CfnUser_HomeDirectoryMapEntryProperty struct {
	// Represents an entry for `HomeDirectoryMappings` .
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Represents the map target that is used in a `HomeDirectorymapEntry` .
	Target *string `field:"required" json:"target" yaml:"target"`
}

// The full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon EFS file systems.
//
// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   posixProfileProperty := &posixProfileProperty{
//   	gid: jsii.Number(123),
//   	uid: jsii.Number(123),
//
//   	// the properties below are optional
//   	secondaryGids: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnUser_PosixProfileProperty struct {
	// The POSIX group ID used for all EFS operations by this user.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all EFS operations by this user.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// The secondary POSIX group IDs used for all EFS operations by this user.
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProps := &cfnUserProps{
//   	role: jsii.String("role"),
//   	serverId: jsii.String("serverId"),
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	homeDirectory: jsii.String("homeDirectory"),
//   	homeDirectoryMappings: []interface{}{
//   		&homeDirectoryMapEntryProperty{
//   			entry: jsii.String("entry"),
//   			target: jsii.String("target"),
//   		},
//   	},
//   	homeDirectoryType: jsii.String("homeDirectoryType"),
//   	policy: jsii.String("policy"),
//   	posixProfile: &posixProfileProperty{
//   		gid: jsii.Number(123),
//   		uid: jsii.Number(123),
//
//   		// the properties below are optional
//   		secondaryGids: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	sshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserProps struct {
	// Specifies the Amazon Resource Name (ARN) of the IAM role that controls your users' access to your Amazon S3 bucket or EFS file system.
	//
	// The policies attached to this role determine the level of access that you want to provide your users when transferring files into and out of your Amazon S3 bucket or EFS file system. The IAM role should also contain a trust relationship that allows the server to access your resources when servicing your users' transfer requests.
	Role *string `field:"required" json:"role" yaml:"role"`
	// A system-assigned unique identifier for a server instance.
	//
	// This is the specific server that you added your user to.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// A unique string that identifies a user and is associated with a `ServerId` .
	//
	// This user name must be a minimum of 3 and a maximum of 100 characters long. The following are valid characters: a-z, A-Z, 0-9, underscore '_', hyphen '-', period '.', and at sign '@'. The user name can't start with a hyphen, period, or at sign.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// The landing directory (folder) for a user when they log in to the server using the client.
	//
	// A `HomeDirectory` example is `/bucket_name/home/mydirectory` .
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// Logical directory mappings that specify what Amazon S3 paths and keys should be visible to your user and how you want to make them visible.
	//
	// You will need to specify the " `Entry` " and " `Target` " pair, where `Entry` shows how the path is made visible and `Target` is the actual Amazon S3 path. If you only specify a target, it will be displayed as is. You will need to also make sure that your IAM role provides access to paths in `Target` . The following is an example.
	//
	// `'[ { "Entry": "/", "Target": "/bucket3/customized-reports/" } ]'`
	//
	// In most cases, you can use this value instead of the session policy to lock your user down to the designated home directory ("chroot"). To do this, you can set `Entry` to '/' and set `Target` to the HomeDirectory parameter value.
	//
	// > If the target of a logical directory entry does not exist in Amazon S3, the entry will be ignored. As a workaround, you can use the Amazon S3 API to create 0 byte objects as place holders for your directory. If using the CLI, use the `s3api` call instead of `s3` so you can use the put-object operation. For example, you use the following: `AWS s3api put-object --bucket bucketname --key path/to/folder/` . Make sure that the end of the key name ends in a '/' for it to be considered a folder.
	HomeDirectoryMappings interface{} `field:"optional" json:"homeDirectoryMappings" yaml:"homeDirectoryMappings"`
	// The type of landing directory (folder) you want your users' home directory to be when they log into the server.
	//
	// If you set it to `PATH` , the user will see the absolute Amazon S3 bucket or EFS paths as is in their file transfer protocol clients. If you set it `LOGICAL` , you need to provide mappings in the `HomeDirectoryMappings` for how you want to make Amazon S3 or EFS paths visible to your users.
	HomeDirectoryType *string `field:"optional" json:"homeDirectoryType" yaml:"homeDirectoryType"`
	// A session policy for your user so you can use the same IAM role across multiple users.
	//
	// This policy restricts user access to portions of their Amazon S3 bucket. Variables that you can use inside this policy include `${Transfer:UserName}` , `${Transfer:HomeDirectory}` , and `${Transfer:HomeBucket}` .
	//
	// > For session policies, AWS Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the `Policy` argument.
	// >
	// > For an example of a session policy, see [Example session policy](https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html) .
	// >
	// > For more information, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the *AWS Security Token Service API Reference* .
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Specifies the full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon Elastic File System (Amazon EFS) file systems.
	//
	// The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
	PosixProfile interface{} `field:"optional" json:"posixProfile" yaml:"posixProfile"`
	// Specifies the public key portion of the Secure Shell (SSH) keys stored for the described user.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Key-value pairs that can be used to group and search for users.
	//
	// Tags are metadata attached to users for any purpose.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Transfer::Workflow`.
//
// Allows you to create a workflow with specified steps and step details the workflow invokes after file transfer completes. After creating a workflow, you can associate the workflow created with any transfer servers by specifying the `workflow-details` field in `CreateServer` and `UpdateServer` operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   cfnWorkflow := awscdk.Aws_transfer.NewCfnWorkflow(this, jsii.String("MyCfnWorkflow"), &cfnWorkflowProps{
//   	steps: []interface{}{
//   		&workflowStepProperty{
//   			copyStepDetails: copyStepDetails,
//   			customStepDetails: customStepDetails,
//   			deleteStepDetails: deleteStepDetails,
//   			tagStepDetails: tagStepDetails,
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	onExceptionSteps: []interface{}{
//   		&workflowStepProperty{
//   			copyStepDetails: copyStepDetails,
//   			customStepDetails: customStepDetails,
//   			deleteStepDetails: deleteStepDetails,
//   			tagStepDetails: tagStepDetails,
//   			type: jsii.String("type"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnWorkflow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// A unique identifier for a workflow.
	AttrWorkflowId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies the text description for the workflow.
	Description() *string
	SetDescription(val *string)
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
	// Specifies the steps (actions) to take if errors are encountered during execution of the workflow.
	OnExceptionSteps() interface{}
	SetOnExceptionSteps(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specifies the details for the steps that are in the specified workflow.
	Steps() interface{}
	SetSteps(val interface{})
	// Key-value pairs that can be used to group and search for workflows.
	//
	// Tags are metadata attached to workflows for any purpose.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnWorkflow
type jsiiProxy_CfnWorkflow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkflow) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) AttrWorkflowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkflowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) OnExceptionSteps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onExceptionSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Steps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Transfer::Workflow`.
func NewCfnWorkflow(scope constructs.Construct, id *string, props *CfnWorkflowProps) CfnWorkflow {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkflow{}

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnWorkflow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Transfer::Workflow`.
func NewCfnWorkflow_Override(c CfnWorkflow, scope constructs.Construct, id *string, props *CfnWorkflowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_transfer.CfnWorkflow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkflow) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWorkflow) SetOnExceptionSteps(val interface{}) {
	_jsii_.Set(
		j,
		"onExceptionSteps",
		val,
	)
}

func (j *jsiiProxy_CfnWorkflow) SetSteps(val interface{}) {
	_jsii_.Set(
		j,
		"steps",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWorkflow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnWorkflow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWorkflow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnWorkflow",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_transfer.CfnWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_transfer.CfnWorkflow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkflow) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWorkflow) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkflow) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWorkflow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWorkflow) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWorkflow) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWorkflow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWorkflow) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkflow) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkflow) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWorkflow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkflow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkflow) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkflow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The basic building block of a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   workflowStepProperty := &workflowStepProperty{
//   	copyStepDetails: copyStepDetails,
//   	customStepDetails: customStepDetails,
//   	deleteStepDetails: deleteStepDetails,
//   	tagStepDetails: tagStepDetails,
//   	type: jsii.String("type"),
//   }
//
type CfnWorkflow_WorkflowStepProperty struct {
	// Details for a step that performs a file copy.
	//
	// Consists of the following values:
	//
	// - A description
	// - An S3 location for the destination of the file copy.
	// - A flag that indicates whether or not to overwrite an existing file of the same name. The default is `FALSE` .
	CopyStepDetails interface{} `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// Details for a step that invokes a lambda function.
	//
	// Consists of the lambda function name, target, and timeout (in seconds).
	CustomStepDetails interface{} `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// Details for a step that deletes the file.
	DeleteStepDetails interface{} `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// Details for a step that creates one or more tags.
	//
	// You specify one or more tags: each tag contains a key/value pair.
	TagStepDetails interface{} `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
	// Currently, the following step types are supported.
	//
	// - *COPY* : copy the file to another location
	// - *CUSTOM* : custom step with a lambda target
	// - *DELETE* : delete the file
	// - *TAG* : add a tag to the file.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// Properties for defining a `CfnWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   cfnWorkflowProps := &cfnWorkflowProps{
//   	steps: []interface{}{
//   		&workflowStepProperty{
//   			copyStepDetails: copyStepDetails,
//   			customStepDetails: customStepDetails,
//   			deleteStepDetails: deleteStepDetails,
//   			tagStepDetails: tagStepDetails,
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	onExceptionSteps: []interface{}{
//   		&workflowStepProperty{
//   			copyStepDetails: copyStepDetails,
//   			customStepDetails: customStepDetails,
//   			deleteStepDetails: deleteStepDetails,
//   			tagStepDetails: tagStepDetails,
//   			type: jsii.String("type"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWorkflowProps struct {
	// Specifies the details for the steps that are in the specified workflow.
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// Specifies the text description for the workflow.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the steps (actions) to take if errors are encountered during execution of the workflow.
	OnExceptionSteps interface{} `field:"optional" json:"onExceptionSteps" yaml:"onExceptionSteps"`
	// Key-value pairs that can be used to group and search for workflows.
	//
	// Tags are metadata attached to workflows for any purpose.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

