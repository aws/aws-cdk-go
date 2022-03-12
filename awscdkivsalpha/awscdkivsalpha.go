// The CDK Construct Library for AWS::IVS
package awscdkivsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkivsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkivsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new IVS channel.
//
// TODO: EXAMPLE
//
// Experimental.
type Channel interface {
	awscdk.Resource
	IChannel
	ChannelArn() *string
	ChannelIngestEndpoint() *string
	ChannelPlaybackUrl() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddStreamKey(id *string) StreamKey
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Channel
type jsiiProxy_Channel struct {
	internal.Type__awscdkResource
	jsiiProxy_IChannel
}

func (j *jsiiProxy_Channel) ChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Channel) ChannelIngestEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelIngestEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Channel) ChannelPlaybackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelPlaybackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Channel) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Channel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Channel) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Channel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewChannel(scope constructs.Construct, id *string, props *ChannelProps) Channel {
	_init_.Initialize()

	j := jsiiProxy_Channel{}

	_jsii_.Create(
		"@aws-cdk/aws-ivs-alpha.Channel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewChannel_Override(c Channel, scope constructs.Construct, id *string, props *ChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-ivs-alpha.Channel",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing channel.
// Experimental.
func Channel_FromChannelArn(scope constructs.Construct, id *string, channelArn *string) IChannel {
	_init_.Initialize()

	var returns IChannel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.Channel",
		"fromChannelArn",
		[]interface{}{scope, id, channelArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Channel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.Channel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Channel_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.Channel",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a stream key for this IVS Channel.
// Experimental.
func (c *jsiiProxy_Channel) AddStreamKey(id *string) StreamKey {
	var returns StreamKey

	_jsii_.Invoke(
		c,
		"addStreamKey",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_Channel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_Channel) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_Channel) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_Channel) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_Channel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a new Channel.
//
// TODO: EXAMPLE
//
// Experimental.
type ChannelProps struct {
	// Whether the channel is authorized.
	//
	// If you wish to make an authorized channel, you will need to ensure that
	// a PlaybackKeyPair has been uploaded to your account as this is used to
	// validate the signed JWT that is required for authorization
	// Experimental.
	Authorized *bool `json:"authorized" yaml:"authorized"`
	// Channel latency mode.
	// Experimental.
	LatencyMode LatencyMode `json:"latencyMode" yaml:"latencyMode"`
	// Channel name.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream will disconnect immediately
	// Experimental.
	Type ChannelType `json:"type" yaml:"type"`
}

// The channel type, which determines the allowable resolution and bitrate.
//
// If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
// Experimental.
type ChannelType string

const (
	ChannelType_STANDARD ChannelType = "STANDARD"
	ChannelType_BASIC ChannelType = "BASIC"
)

// Represents an IVS Channel.
// Experimental.
type IChannel interface {
	awscdk.IResource
	// Adds a stream key for this IVS Channel.
	// Experimental.
	AddStreamKey(id *string) StreamKey
	// The channel ARN.
	//
	// For example: arn:aws:ivs:us-west-2:123456789012:channel/abcdABCDefgh
	// Experimental.
	ChannelArn() *string
}

// The jsii proxy for IChannel
type jsiiProxy_IChannel struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannel) AddStreamKey(id *string) StreamKey {
	var returns StreamKey

	_jsii_.Invoke(
		i,
		"addStreamKey",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IChannel) ChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelArn",
		&returns,
	)
	return returns
}

// Represents an IVS Playback Key Pair.
// Experimental.
type IPlaybackKeyPair interface {
	awscdk.IResource
	// Key-pair ARN.
	//
	// For example: arn:aws:ivs:us-west-2:693991300569:playback-key/f99cde61-c2b0-4df3-8941-ca7d38acca1a
	// Experimental.
	PlaybackKeyPairArn() *string
}

// The jsii proxy for IPlaybackKeyPair
type jsiiProxy_IPlaybackKeyPair struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPlaybackKeyPair) PlaybackKeyPairArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playbackKeyPairArn",
		&returns,
	)
	return returns
}

// Represents an IVS Stream Key.
// Experimental.
type IStreamKey interface {
	awscdk.IResource
	// The stream-key ARN.
	//
	// For example: arn:aws:ivs:us-west-2:123456789012:stream-key/g1H2I3j4k5L6
	// Experimental.
	StreamKeyArn() *string
}

// The jsii proxy for IStreamKey
type jsiiProxy_IStreamKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IStreamKey) StreamKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamKeyArn",
		&returns,
	)
	return returns
}

// Channel latency mode.
// Experimental.
type LatencyMode string

const (
	LatencyMode_LOW LatencyMode = "LOW"
	LatencyMode_NORMAL LatencyMode = "NORMAL"
)

// A new IVS Playback Key Pair.
//
// TODO: EXAMPLE
//
// Experimental.
type PlaybackKeyPair interface {
	awscdk.Resource
	IPlaybackKeyPair
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	PlaybackKeyPairArn() *string
	PlaybackKeyPairFingerprint() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for PlaybackKeyPair
type jsiiProxy_PlaybackKeyPair struct {
	internal.Type__awscdkResource
	jsiiProxy_IPlaybackKeyPair
}

func (j *jsiiProxy_PlaybackKeyPair) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PlaybackKeyPair) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PlaybackKeyPair) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PlaybackKeyPair) PlaybackKeyPairArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playbackKeyPairArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PlaybackKeyPair) PlaybackKeyPairFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playbackKeyPairFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PlaybackKeyPair) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewPlaybackKeyPair(scope constructs.Construct, id *string, props *PlaybackKeyPairProps) PlaybackKeyPair {
	_init_.Initialize()

	j := jsiiProxy_PlaybackKeyPair{}

	_jsii_.Create(
		"@aws-cdk/aws-ivs-alpha.PlaybackKeyPair",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPlaybackKeyPair_Override(p PlaybackKeyPair, scope constructs.Construct, id *string, props *PlaybackKeyPairProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-ivs-alpha.PlaybackKeyPair",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PlaybackKeyPair_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.PlaybackKeyPair",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func PlaybackKeyPair_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.PlaybackKeyPair",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (p *jsiiProxy_PlaybackKeyPair) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (p *jsiiProxy_PlaybackKeyPair) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (p *jsiiProxy_PlaybackKeyPair) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (p *jsiiProxy_PlaybackKeyPair) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_PlaybackKeyPair) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a new Playback Key Pair.
//
// TODO: EXAMPLE
//
// Experimental.
type PlaybackKeyPairProps struct {
	// The public portion of a customer-generated key pair.
	// Experimental.
	PublicKeyMaterial *string `json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource.
	//
	// The value does not need to be unique.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// A new IVS Stream Key.
//
// TODO: EXAMPLE
//
// Experimental.
type StreamKey interface {
	awscdk.Resource
	IStreamKey
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	StreamKeyArn() *string
	StreamKeyValue() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for StreamKey
type jsiiProxy_StreamKey struct {
	internal.Type__awscdkResource
	jsiiProxy_IStreamKey
}

func (j *jsiiProxy_StreamKey) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamKey) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamKey) StreamKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamKey) StreamKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamKeyValue",
		&returns,
	)
	return returns
}


// Experimental.
func NewStreamKey(scope constructs.Construct, id *string, props *StreamKeyProps) StreamKey {
	_init_.Initialize()

	j := jsiiProxy_StreamKey{}

	_jsii_.Create(
		"@aws-cdk/aws-ivs-alpha.StreamKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStreamKey_Override(s StreamKey, scope constructs.Construct, id *string, props *StreamKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-ivs-alpha.StreamKey",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StreamKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.StreamKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func StreamKey_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.StreamKey",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (s *jsiiProxy_StreamKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (s *jsiiProxy_StreamKey) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (s *jsiiProxy_StreamKey) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (s *jsiiProxy_StreamKey) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_StreamKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a new Stream Key.
//
// TODO: EXAMPLE
//
// Experimental.
type StreamKeyProps struct {
	// Channel ARN for the stream.
	// Experimental.
	Channel IChannel `json:"channel" yaml:"channel"`
}

