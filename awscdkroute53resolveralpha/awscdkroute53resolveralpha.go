// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The way that you want DNS Firewall to block the request.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(10),
//   	firewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
//   })
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(20),
//   	firewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
//   })
//
// Experimental.
type DnsBlockResponse interface {
	// The DNS record's type.
	// Experimental.
	BlockOverrideDnsType() *string
	// The custom DNS record to send back in response to the query.
	// Experimental.
	BlockOverrideDomain() *string
	// The recommended amount of time for the DNS resolver or web browser to cache the provided override record.
	// Experimental.
	BlockOverrideTtl() awscdk.Duration
	// The way that you want DNS Firewall to block the request.
	// Experimental.
	BlockResponse() *string
}

// The jsii proxy struct for DnsBlockResponse
type jsiiProxy_DnsBlockResponse struct {
	_ byte // padding
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideDnsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideTtl() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"blockOverrideTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewDnsBlockResponse_Override(d DnsBlockResponse) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
		nil, // no parameters
		d,
	)
}

// Respond indicating that the query was successful, but no response is available for it.
// Experimental.
func DnsBlockResponse_NoData() DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
		"noData",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Respond indicating that the domain name that's in the query doesn't exist.
// Experimental.
func DnsBlockResponse_NxDomain() DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
		"nxDomain",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Provides a custom override response to the query.
// Experimental.
func DnsBlockResponse_Override(domain *string, ttl awscdk.Duration) DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
		"override",
		[]interface{}{domain, ttl},
		&returns,
	)

	return returns
}

// Domains configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import route53resolver_alpha "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha"
//
//   domainsConfig := &domainsConfig{
//   	domainFileUrl: jsii.String("domainFileUrl"),
//   	domains: []*string{
//   		jsii.String("domains"),
//   	},
//   }
//
// Experimental.
type DomainsConfig struct {
	// The fully qualified URL or URI of the file stored in Amazon S3 that contains the list of domains to import.
	//
	// The file must be a text file and must contain
	// a single domain per line. The content type of the S3 object must be `plain/text`.
	// Experimental.
	DomainFileUrl *string `field:"optional" json:"domainFileUrl" yaml:"domainFileUrl"`
	// A list of domains.
	// Experimental.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
}

// A Firewall Domain List.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomainList interface {
	awscdk.Resource
	IFirewallDomainList
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN (Amazon Resource Name) of the domain list.
	// Experimental.
	FirewallDomainListArn() *string
	// The date and time that the domain list was created.
	// Experimental.
	FirewallDomainListCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallDomainListCreatorRequestId() *string
	// The number of domains in the list.
	// Experimental.
	FirewallDomainListDomainCount() *float64
	// The ID of the domain list.
	// Experimental.
	FirewallDomainListId() *string
	// The owner of the list, used only for lists that are not managed by you.
	//
	// For example, the managed domain list `AWSManagedDomainsMalwareDomainList`
	// has the managed owner name `Route 53 Resolver DNS Firewall`.
	// Experimental.
	FirewallDomainListManagedOwnerName() *string
	// The date and time that the domain list was last modified.
	// Experimental.
	FirewallDomainListModificationTime() *string
	// The status of the domain list.
	// Experimental.
	FirewallDomainListStatus() *string
	// Additional information about the status of the rule group.
	// Experimental.
	FirewallDomainListStatusMessage() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallDomainList
type jsiiProxy_FirewallDomainList struct {
	internal.Type__awscdkResource
	jsiiProxy_IFirewallDomainList
}

func (j *jsiiProxy_FirewallDomainList) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallDomainListDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallDomainList(scope constructs.Construct, id *string, props *FirewallDomainListProps) FirewallDomainList {
	_init_.Initialize()

	j := jsiiProxy_FirewallDomainList{}

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallDomainList_Override(f FirewallDomainList, scope constructs.Construct, id *string, props *FirewallDomainListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing Firewall Rule Group.
// Experimental.
func FirewallDomainList_FromFirewallDomainListId(scope constructs.Construct, id *string, firewallDomainListId *string) IFirewallDomainList {
	_init_.Initialize()

	var returns IFirewallDomainList

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"fromFirewallDomainListId",
		[]interface{}{scope, id, firewallDomainListId},
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
// Experimental.
func FirewallDomainList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func FirewallDomainList_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallDomainList_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallDomainList) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Firewall Domain List.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomainListProps struct {
	// A list of domains.
	// Experimental.
	Domains FirewallDomains `field:"required" json:"domains" yaml:"domains"`
	// A name for the domain list.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A list of domains.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomains interface {
	// Binds the domains to a domain list.
	// Experimental.
	Bind(scope constructs.Construct) *DomainsConfig
}

// The jsii proxy struct for FirewallDomains
type jsiiProxy_FirewallDomains struct {
	_ byte // padding
}

// Experimental.
func NewFirewallDomains_Override(f FirewallDomains) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		nil, // no parameters
		f,
	)
}

// Firewall domains created from a local disk path to a text file.
//
// The file must be a text file (`.txt` extension) and must contain a single
// domain per line. It will be uploaded to S3.
// Experimental.
func FirewallDomains_FromAsset(assetPath *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		"fromAsset",
		[]interface{}{assetPath},
		&returns,
	)

	return returns
}

// Firewall domains created from a list of domains.
// Experimental.
func FirewallDomains_FromList(list *[]*string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		"fromList",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// Firewall domains created from a file stored in Amazon S3.
//
// The file must be a text file and must contain a single domain per line.
// The content type of the S3 object must be `plain/text`.
// Experimental.
func FirewallDomains_FromS3(bucket awss3.IBucket, key *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Firewall domains created from the URL of a file stored in Amazon S3.
//
// The file must be a text file and must contain a single domain per line.
// The content type of the S3 object must be `plain/text`.
// Experimental.
func FirewallDomains_FromS3Url(url *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		"fromS3Url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomains) Bind(scope constructs.Construct) *DomainsConfig {
	var returns *DomainsConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(10),
//   	firewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
//   })
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(20),
//   	firewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
//   })
//
// Experimental.
type FirewallRule struct {
	// The action for this rule.
	// Experimental.
	Action FirewallRuleAction `field:"required" json:"action" yaml:"action"`
	// The domain list for this rule.
	// Experimental.
	FirewallDomainList IFirewallDomainList `field:"required" json:"firewallDomainList" yaml:"firewallDomainList"`
	// The priority of the rule in the rule group.
	//
	// This value must be unique within
	// the rule group.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleAction interface {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list.
	// Experimental.
	Action() *string
	// The way that you want DNS Firewall to block the request.
	// Experimental.
	BlockResponse() DnsBlockResponse
}

// The jsii proxy struct for FirewallRuleAction
type jsiiProxy_FirewallRuleAction struct {
	_ byte // padding
}

func (j *jsiiProxy_FirewallRuleAction) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleAction) BlockResponse() DnsBlockResponse {
	var returns DnsBlockResponse
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleAction_Override(f FirewallRuleAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		nil, // no parameters
		f,
	)
}

// Permit the request to go through but send an alert to the logs.
// Experimental.
func FirewallRuleAction_Alert() FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		"alert",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Permit the request to go through.
// Experimental.
func FirewallRuleAction_Allow() FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		"allow",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Disallow the request.
// Experimental.
func FirewallRuleAction_Block(response DnsBlockResponse) FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		"block",
		[]interface{}{response},
		&returns,
	)

	return returns
}

// A Firewall Rule Group.
//
// Example:
//   var myBlockList firewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleGroup interface {
	awscdk.Resource
	IFirewallRuleGroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN (Amazon Resource Name) of the rule group.
	// Experimental.
	FirewallRuleGroupArn() *string
	// The date and time that the rule group was created.
	// Experimental.
	FirewallRuleGroupCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallRuleGroupCreatorRequestId() *string
	// The ID of the rule group.
	// Experimental.
	FirewallRuleGroupId() *string
	// The date and time that the rule group was last modified.
	// Experimental.
	FirewallRuleGroupModificationTime() *string
	// The AWS account ID for the account that created the rule group.
	// Experimental.
	FirewallRuleGroupOwnerId() *string
	// The number of rules in the rule group.
	// Experimental.
	FirewallRuleGroupRuleCount() *float64
	// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account.
	// Experimental.
	FirewallRuleGroupShareStatus() *string
	// The status of the rule group.
	// Experimental.
	FirewallRuleGroupStatus() *string
	// Additional information about the status of the rule group.
	// Experimental.
	FirewallRuleGroupStatusMessage() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a rule to this group.
	// Experimental.
	AddRule(rule *FirewallRule) FirewallRuleGroup
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Associates this Firewall Rule Group with a VPC.
	// Experimental.
	Associate(id *string, props *FirewallRuleGroupAssociationOptions) FirewallRuleGroupAssociation
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallRuleGroup
type jsiiProxy_FirewallRuleGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IFirewallRuleGroup
}

func (j *jsiiProxy_FirewallRuleGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupRuleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallRuleGroupRuleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleGroup(scope constructs.Construct, id *string, props *FirewallRuleGroupProps) FirewallRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_FirewallRuleGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallRuleGroup_Override(f FirewallRuleGroup, scope constructs.Construct, id *string, props *FirewallRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing Firewall Rule Group.
// Experimental.
func FirewallRuleGroup_FromFirewallRuleGroupId(scope constructs.Construct, id *string, firewallRuleGroupId *string) IFirewallRuleGroup {
	_init_.Initialize()

	var returns IFirewallRuleGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"fromFirewallRuleGroupId",
		[]interface{}{scope, id, firewallRuleGroupId},
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
// Experimental.
func FirewallRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func FirewallRuleGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallRuleGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) AddRule(rule *FirewallRule) FirewallRuleGroup {
	var returns FirewallRuleGroup

	_jsii_.Invoke(
		f,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallRuleGroup) Associate(id *string, props *FirewallRuleGroupAssociationOptions) FirewallRuleGroupAssociation {
	var returns FirewallRuleGroupAssociation

	_jsii_.Invoke(
		f,
		"associate",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A Firewall Rule Group Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import route53resolver_alpha "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var firewallRuleGroup firewallRuleGroup
//   var vpc vpc
//
//   firewallRuleGroupAssociation := route53resolver_alpha.NewFirewallRuleGroupAssociation(this, jsii.String("MyFirewallRuleGroupAssociation"), &firewallRuleGroupAssociationProps{
//   	firewallRuleGroup: firewallRuleGroup,
//   	priority: jsii.Number(123),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	mutationProtection: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   })
//
// Experimental.
type FirewallRuleGroupAssociation interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN (Amazon Resource Name) of the association.
	// Experimental.
	FirewallRuleGroupAssociationArn() *string
	// The date and time that the association was created.
	// Experimental.
	FirewallRuleGroupAssociationCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallRuleGroupAssociationCreatorRequestId() *string
	// The ID of the association.
	// Experimental.
	FirewallRuleGroupAssociationId() *string
	// The owner of the association, used only for lists that are not managed by you.
	//
	// If you use AWS Firewall Manager to manage your firewallls from DNS Firewall,
	// then this reports Firewall Manager as the managed owner.
	// Experimental.
	FirewallRuleGroupAssociationManagedOwnerName() *string
	// The date and time that the association was last modified.
	// Experimental.
	FirewallRuleGroupAssociationModificationTime() *string
	// The status of the association.
	// Experimental.
	FirewallRuleGroupAssociationStatus() *string
	// Additional information about the status of the association.
	// Experimental.
	FirewallRuleGroupAssociationStatusMessage() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallRuleGroupAssociation
type jsiiProxy_FirewallRuleGroupAssociation struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleGroupAssociation(scope constructs.Construct, id *string, props *FirewallRuleGroupAssociationProps) FirewallRuleGroupAssociation {
	_init_.Initialize()

	j := jsiiProxy_FirewallRuleGroupAssociation{}

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallRuleGroupAssociation_Override(f FirewallRuleGroupAssociation, scope constructs.Construct, id *string, props *FirewallRuleGroupAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		f,
	)
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
// Experimental.
func FirewallRuleGroupAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func FirewallRuleGroupAssociation_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociation",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallRuleGroupAssociation_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroupAssociation",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for a Firewall Rule Group Association.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleGroup firewallRuleGroup
//   var myVpc vpc
//
//
//   ruleGroup.associate(jsii.String("Association"), &firewallRuleGroupAssociationOptions{
//   	priority: jsii.Number(101),
//   	vpc: myVpc,
//   })
//
// Experimental.
type FirewallRuleGroupAssociationOptions struct {
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC
	// traffic starting from rule group with the lowest numeric priority setting.
	//
	// This value must be greater than 100 and less than 9,000.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The VPC that to associate with the rule group.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	// Experimental.
	MutationProtection *bool `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for a Firewall Rule Group Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import route53resolver_alpha "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var firewallRuleGroup firewallRuleGroup
//   var vpc vpc
//
//   firewallRuleGroupAssociationProps := &firewallRuleGroupAssociationProps{
//   	firewallRuleGroup: firewallRuleGroup,
//   	priority: jsii.Number(123),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	mutationProtection: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   }
//
// Experimental.
type FirewallRuleGroupAssociationProps struct {
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC
	// traffic starting from rule group with the lowest numeric priority setting.
	//
	// This value must be greater than 100 and less than 9,000.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The VPC that to associate with the rule group.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	// Experimental.
	MutationProtection *bool `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The firewall rule group which must be associated.
	// Experimental.
	FirewallRuleGroup IFirewallRuleGroup `field:"required" json:"firewallRuleGroup" yaml:"firewallRuleGroup"`
}

// Properties for a Firewall Rule Group.
//
// Example:
//   var myBlockList firewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleGroupProps struct {
	// The name of the rule group.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of rules for this group.
	// Experimental.
	Rules *[]*FirewallRule `field:"optional" json:"rules" yaml:"rules"`
}

// A Firewall Domain List.
// Experimental.
type IFirewallDomainList interface {
	awscdk.IResource
	// The ID of the domain list.
	// Experimental.
	FirewallDomainListId() *string
}

// The jsii proxy for IFirewallDomainList
type jsiiProxy_IFirewallDomainList struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFirewallDomainList) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

// A Firewall Rule Group.
// Experimental.
type IFirewallRuleGroup interface {
	awscdk.IResource
	// The ID of the rule group.
	// Experimental.
	FirewallRuleGroupId() *string
}

// The jsii proxy for IFirewallRuleGroup
type jsiiProxy_IFirewallRuleGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFirewallRuleGroup) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

