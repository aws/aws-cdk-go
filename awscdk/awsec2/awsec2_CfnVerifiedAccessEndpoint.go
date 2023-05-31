package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::EC2::VerifiedAccessEndpoint`.
//
// An AWS Verified Access endpoint specifies the application that AWS Verified Access provides access to. It must be attached to an AWS Verified Access group. An AWS Verified Access endpoint must also have an attached access policy before you attached it to a group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVerifiedAccessEndpoint := awscdk.Aws_ec2.NewCfnVerifiedAccessEndpoint(this, jsii.String("MyCfnVerifiedAccessEndpoint"), &CfnVerifiedAccessEndpointProps{
//   	ApplicationDomain: jsii.String("applicationDomain"),
//   	AttachmentType: jsii.String("attachmentType"),
//   	DomainCertificateArn: jsii.String("domainCertificateArn"),
//   	EndpointDomainPrefix: jsii.String("endpointDomainPrefix"),
//   	EndpointType: jsii.String("endpointType"),
//   	VerifiedAccessGroupId: jsii.String("verifiedAccessGroupId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LoadBalancerOptions: &LoadBalancerOptionsProperty{
//   		LoadBalancerArn: jsii.String("loadBalancerArn"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	NetworkInterfaceOptions: &NetworkInterfaceOptionsProperty{
//   		NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyEnabled: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnVerifiedAccessEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The DNS name for users to reach your application.
	ApplicationDomain() *string
	SetApplicationDomain(val *string)
	// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
	AttachmentType() *string
	SetAttachmentType(val *string)
	// The creation time.
	AttrCreationTime() *string
	// Use this to construct the redirect URI to add to your OIDC provider's allow list.
	AttrDeviceValidationDomain() *string
	// The DNS name generated for the endpoint.
	AttrEndpointDomain() *string
	// The last updated time.
	AttrLastUpdatedTime() *string
	// The endpoint status.
	AttrStatus() *string
	// The ID of the Verified Access endpoint.
	AttrVerifiedAccessEndpointId() *string
	// The instance identifier.
	AttrVerifiedAccessInstanceId() *string
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
	// A description for the AWS Verified Access endpoint.
	Description() *string
	SetDescription(val *string)
	// The ARN of a public TLS/SSL certificate imported into or created with ACM.
	DomainCertificateArn() *string
	SetDomainCertificateArn(val *string)
	// A custom identifier that is prepended to the DNS name that is generated for the endpoint.
	EndpointDomainPrefix() *string
	SetEndpointDomainPrefix(val *string)
	// The type of AWS Verified Access endpoint.
	//
	// Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
	EndpointType() *string
	SetEndpointType(val *string)
	// The load balancer details if creating the AWS Verified Access endpoint as `load-balancer` type.
	LoadBalancerOptions() interface{}
	SetLoadBalancerOptions(val interface{})
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
	// The options for network-interface type endpoint.
	NetworkInterfaceOptions() interface{}
	SetNetworkInterfaceOptions(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The Verified Access policy document.
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	// The status of the Verified Access policy.
	PolicyEnabled() interface{}
	SetPolicyEnabled(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The IDs of the security groups for the endpoint.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId() *string
	SetVerifiedAccessGroupId(val *string)
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

// The jsii proxy struct for CfnVerifiedAccessEndpoint
type jsiiProxy_CfnVerifiedAccessEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) ApplicationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrDeviceValidationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDeviceValidationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrEndpointDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrVerifiedAccessEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVerifiedAccessEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) AttrVerifiedAccessInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVerifiedAccessInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) DomainCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) EndpointDomainPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) LoadBalancerOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) NetworkInterfaceOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) PolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint) VerifiedAccessGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupId",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::VerifiedAccessEndpoint`.
func NewCfnVerifiedAccessEndpoint(scope awscdk.Construct, id *string, props *CfnVerifiedAccessEndpointProps) CfnVerifiedAccessEndpoint {
	_init_.Initialize()

	if err := validateNewCfnVerifiedAccessEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVerifiedAccessEndpoint{}

	_jsii_.Create(
		"monocdk.aws_ec2.CfnVerifiedAccessEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::VerifiedAccessEndpoint`.
func NewCfnVerifiedAccessEndpoint_Override(c CfnVerifiedAccessEndpoint, scope awscdk.Construct, id *string, props *CfnVerifiedAccessEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.CfnVerifiedAccessEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetApplicationDomain(val *string) {
	if err := j.validateSetApplicationDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationDomain",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetAttachmentType(val *string) {
	if err := j.validateSetAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentType",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetDomainCertificateArn(val *string) {
	if err := j.validateSetDomainCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainCertificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetEndpointDomainPrefix(val *string) {
	if err := j.validateSetEndpointDomainPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointDomainPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetLoadBalancerOptions(val interface{}) {
	if err := j.validateSetLoadBalancerOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerOptions",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetNetworkInterfaceOptions(val interface{}) {
	if err := j.validateSetNetworkInterfaceOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceOptions",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetPolicyDocument(val *string) {
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetPolicyEnabled(val interface{}) {
	if err := j.validateSetPolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnVerifiedAccessEndpoint)SetVerifiedAccessGroupId(val *string) {
	if err := j.validateSetVerifiedAccessGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAccessGroupId",
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
func CfnVerifiedAccessEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedAccessEndpoint_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnVerifiedAccessEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVerifiedAccessEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedAccessEndpoint_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnVerifiedAccessEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnVerifiedAccessEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedAccessEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnVerifiedAccessEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVerifiedAccessEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ec2.CfnVerifiedAccessEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessEndpoint) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

