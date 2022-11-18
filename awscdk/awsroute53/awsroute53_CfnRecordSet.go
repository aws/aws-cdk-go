package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Route53::RecordSet`.
//
// Information about the record that you want to create.
//
// The `AWS::Route53::RecordSet` type can be used as a standalone resource or as an embedded property in the `AWS::Route53::RecordSetGroup` type. Note that some `AWS::Route53::RecordSet` properties are valid only when used within `AWS::Route53::RecordSetGroup` .
//
// For more information, see [ChangeResourceRecordSets](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) in the *Amazon Route 53 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordSet := awscdk.Aws_route53.NewCfnRecordSet(this, jsii.String("MyCfnRecordSet"), &cfnRecordSetProps{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	aliasTarget: &aliasTargetProperty{
//   		dnsName: jsii.String("dnsName"),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//
//   		// the properties below are optional
//   		evaluateTargetHealth: jsii.Boolean(false),
//   	},
//   	cidrRoutingConfig: &cidrRoutingConfigProperty{
//   		collectionId: jsii.String("collectionId"),
//   		locationName: jsii.String("locationName"),
//   	},
//   	comment: jsii.String("comment"),
//   	failover: jsii.String("failover"),
//   	geoLocation: &geoLocationProperty{
//   		continentCode: jsii.String("continentCode"),
//   		countryCode: jsii.String("countryCode"),
//   		subdivisionCode: jsii.String("subdivisionCode"),
//   	},
//   	healthCheckId: jsii.String("healthCheckId"),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	hostedZoneName: jsii.String("hostedZoneName"),
//   	multiValueAnswer: jsii.Boolean(false),
//   	region: jsii.String("region"),
//   	resourceRecords: []*string{
//   		jsii.String("resourceRecords"),
//   	},
//   	setIdentifier: jsii.String("setIdentifier"),
//   	ttl: jsii.String("ttl"),
//   	weight: jsii.Number(123),
//   })
//
type CfnRecordSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// *Alias resource record sets only:* Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.
	//
	// If you're creating resource records sets for a private hosted zone, note the following:
	//
	// - You can't create an alias resource record set in a private hosted zone to route traffic to a CloudFront distribution.
	// - Creating geolocation alias resource record sets or latency alias resource record sets in a private hosted zone is unsupported.
	// - For information about creating failover resource record sets in a private hosted zone, see [Configuring Failover in a Private Hosted Zone](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-private-hosted-zones.html) in the *Amazon Route 53 Developer Guide* .
	AliasTarget() interface{}
	SetAliasTarget(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Route53::RecordSet.CidrRoutingConfig`.
	CidrRoutingConfig() interface{}
	SetCidrRoutingConfig(val interface{})
	// *Optional:* Any comments you want to include about a change batch request.
	Comment() *string
	SetComment(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// *Failover resource record sets only:* To configure failover, you add the `Failover` element to two resource record sets.
	//
	// For one resource record set, you specify `PRIMARY` as the value for `Failover` ; for the other resource record set, you specify `SECONDARY` . In addition, you include the `HealthCheckId` element and specify the health check that you want Amazon Route 53 to perform for each resource record set.
	//
	// Except where noted, the following failover behaviors assume that you have included the `HealthCheckId` element in both resource record sets:
	//
	// - When the primary resource record set is healthy, Route 53 responds to DNS queries with the applicable value from the primary resource record set regardless of the health of the secondary resource record set.
	// - When the primary resource record set is unhealthy and the secondary resource record set is healthy, Route 53 responds to DNS queries with the applicable value from the secondary resource record set.
	// - When the secondary resource record set is unhealthy, Route 53 responds to DNS queries with the applicable value from the primary resource record set regardless of the health of the primary resource record set.
	// - If you omit the `HealthCheckId` element for the secondary resource record set, and if the primary resource record set is unhealthy, Route 53 always responds to DNS queries with the applicable value from the secondary resource record set. This is true regardless of the health of the associated endpoint.
	//
	// You can't create non-failover resource record sets that have the same values for the `Name` and `Type` elements as failover resource record sets.
	//
	// For failover alias resource record sets, you must also include the `EvaluateTargetHealth` element and set the value to true.
	//
	// For more information about configuring failover for Route 53, see the following topics in the *Amazon Route 53 Developer Guide* :
	//
	// - [Route 53 Health Checks and DNS Failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover.html)
	// - [Configuring Failover in a Private Hosted Zone](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-private-hosted-zones.html)
	Failover() *string
	SetFailover(val *string)
	// *Geolocation resource record sets only:* A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.
	//
	// For example, if you want all queries from Africa to be routed to a web server with an IP address of `192.0.2.111` , create a resource record set with a `Type` of `A` and a `ContinentCode` of `AF` .
	//
	// > Although creating geolocation and geolocation alias resource record sets in a private hosted zone is allowed, it's not supported.
	//
	// If you create separate resource record sets for overlapping geographic regions (for example, one resource record set for a continent and one for a country on the same continent), priority goes to the smallest geographic region. This allows you to route most queries for a continent to one resource and to route queries for a country on that continent to a different resource.
	//
	// You can't create two geolocation resource record sets that specify the same geographic location.
	//
	// The value `*` in the `CountryCode` element matches all geographic locations that aren't specified in other geolocation resource record sets that have the same values for the `Name` and `Type` elements.
	//
	// > Geolocation works by mapping IP addresses to locations. However, some IP addresses aren't mapped to geographic locations, so even if you create geolocation resource record sets that cover all seven continents, Route 53 will receive some DNS queries from locations that it can't identify. We recommend that you create a resource record set for which the value of `CountryCode` is `*` . Two groups of queries are routed to the resource that you specify in this record: queries that come from locations for which you haven't created geolocation resource record sets and queries from IP addresses that aren't mapped to a location. If you don't create a `*` resource record set, Route 53 returns a "no answer" response for queries from those locations.
	//
	// You can't create non-geolocation resource record sets that have the same values for the `Name` and `Type` elements as geolocation resource record sets.
	GeoLocation() interface{}
	SetGeoLocation(val interface{})
	// If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the `HealthCheckId` element and specify the ID of the applicable health check.
	//
	// Route 53 determines whether a resource record set is healthy based on one of the following:
	//
	// - By periodically sending a request to the endpoint that is specified in the health check
	// - By aggregating the status of a specified group of health checks (calculated health checks)
	// - By determining the current state of a CloudWatch alarm (CloudWatch metric health checks)
	//
	// > Route 53 doesn't check the health of the endpoint that is specified in the resource record set, for example, the endpoint specified by the IP address in the `Value` element. When you add a `HealthCheckId` element to a resource record set, Route 53 checks the health of the endpoint that you specified in the health check.
	//
	// For more information, see the following topics in the *Amazon Route 53 Developer Guide* :
	//
	// - [How Amazon Route 53 Determines Whether an Endpoint Is Healthy](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html)
	// - [Route 53 Health Checks and DNS Failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover.html)
	// - [Configuring Failover in a Private Hosted Zone](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-private-hosted-zones.html)
	//
	// *When to Specify HealthCheckId*
	//
	// Specifying a value for `HealthCheckId` is useful only when Route 53 is choosing between two or more resource record sets to respond to a DNS query, and you want Route 53 to base the choice in part on the status of a health check. Configuring health checks makes sense only in the following configurations:
	//
	// - *Non-alias resource record sets* : You're checking the health of a group of non-alias resource record sets that have the same routing policy, name, and type (such as multiple weighted records named www.example.com with a type of A) and you specify health check IDs for all the resource record sets.
	//
	// If the health check status for a resource record set is healthy, Route 53 includes the record among the records that it responds to DNS queries with.
	//
	// If the health check status for a resource record set is unhealthy, Route 53 stops responding to DNS queries using the value for that resource record set.
	//
	// If the health check status for all resource record sets in the group is unhealthy, Route 53 considers all resource record sets in the group healthy and responds to DNS queries accordingly.
	// - *Alias resource record sets* : You specify the following settings:
	//
	// - You set `EvaluateTargetHealth` to true for an alias resource record set in a group of resource record sets that have the same routing policy, name, and type (such as multiple weighted records named www.example.com with a type of A).
	// - You configure the alias resource record set to route traffic to a non-alias resource record set in the same hosted zone.
	// - You specify a health check ID for the non-alias resource record set.
	//
	// If the health check status is healthy, Route 53 considers the alias resource record set to be healthy and includes the alias record among the records that it responds to DNS queries with.
	//
	// If the health check status is unhealthy, Route 53 stops responding to DNS queries using the alias resource record set.
	//
	// > The alias resource record set can also route traffic to a *group* of non-alias resource record sets that have the same routing policy, name, and type. In that configuration, associate health checks with all of the resource record sets in the group of non-alias resource record sets.
	//
	// *Geolocation Routing*
	//
	// For geolocation resource record sets, if an endpoint is unhealthy, Route 53 looks for a resource record set for the larger, associated geographic region. For example, suppose you have resource record sets for a state in the United States, for the entire United States, for North America, and a resource record set that has `*` for `CountryCode` is `*` , which applies to all locations. If the endpoint for the state resource record set is unhealthy, Route 53 checks for healthy resource record sets in the following order until it finds a resource record set for which the endpoint is healthy:
	//
	// - The United States
	// - North America
	// - The default resource record set
	//
	// *Specifying the Health Check Endpoint by Domain Name*
	//
	// If your health checks specify the endpoint only by domain name, we recommend that you create a separate health check for each endpoint. For example, create a health check for each `HTTP` server that is serving content for `www.example.com` . For the value of `FullyQualifiedDomainName` , specify the domain name of the server (such as `us-east-2-www.example.com` ), not the name of the resource record sets ( `www.example.com` ).
	//
	// > Health check results will be unpredictable if you do the following:
	// >
	// > - Create a health check that has the same value for `FullyQualifiedDomainName` as the name of a resource record set.
	// > - Associate that health check with the resource record set.
	HealthCheckId() *string
	SetHealthCheckId(val *string)
	// The ID of the hosted zone that you want to create records in.
	//
	// Specify either `HostedZoneName` or `HostedZoneId` , but not both. If you have multiple hosted zones with the same domain name, you must specify the hosted zone using `HostedZoneId` .
	HostedZoneId() *string
	SetHostedZoneId(val *string)
	// The name of the hosted zone that you want to create records in.
	//
	// You must include a trailing dot (for example, `www.example.com.` ) as part of the `HostedZoneName` .
	//
	// When you create a stack using an AWS::Route53::RecordSet that specifies `HostedZoneName` , AWS CloudFormation attempts to find a hosted zone whose name matches the HostedZoneName. If AWS CloudFormation cannot find a hosted zone with a matching domain name, or if there is more than one hosted zone with the specified domain name, AWS CloudFormation will not create the stack.
	//
	// Specify either `HostedZoneName` or `HostedZoneId` , but not both. If you have multiple hosted zones with the same domain name, you must specify the hosted zone using `HostedZoneId` .
	HostedZoneName() *string
	SetHostedZoneName(val *string)
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
	// *Multivalue answer resource record sets only* : To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify `true` for `MultiValueAnswer` .
	//
	// Note the following:
	//
	// - If you associate a health check with a multivalue answer resource record set, Amazon Route 53 responds to DNS queries with the corresponding IP address only when the health check is healthy.
	// - If you don't associate a health check with a multivalue answer record, Route 53 always considers the record to be healthy.
	// - Route 53 responds to DNS queries with up to eight healthy records; if you have eight or fewer healthy records, Route 53 responds to all DNS queries with all the healthy records.
	// - If you have more than eight healthy records, Route 53 responds to different DNS resolvers with different combinations of healthy records.
	// - When all records are unhealthy, Route 53 responds to DNS queries with up to eight unhealthy records.
	// - If a resource becomes unavailable after a resolver caches a response, client software typically tries another of the IP addresses in the response.
	//
	// You can't create multivalue answer alias records.
	MultiValueAnswer() interface{}
	SetMultiValueAnswer(val interface{})
	// For `ChangeResourceRecordSets` requests, the name of the record that you want to create, update, or delete.
	//
	// For `ListResourceRecordSets` responses, the name of a record in the specified hosted zone.
	//
	// *ChangeResourceRecordSets Only*
	//
	// Enter a fully qualified domain name, for example, `www.example.com` . You can optionally include a trailing dot. If you omit the trailing dot, Amazon Route 53 assumes that the domain name that you specify is fully qualified. This means that Route 53 treats `www.example.com` (without a trailing dot) and `www.example.com.` (with a trailing dot) as identical.
	//
	// For information about how to specify characters other than `a-z` , `0-9` , and `-` (hyphen) and how to specify internationalized domain names, see [DNS Domain Name Format](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html) in the *Amazon Route 53 Developer Guide* .
	//
	// You can use the asterisk (*) wildcard to replace the leftmost label in a domain name, for example, `*.example.com` . Note the following:
	//
	// - The * must replace the entire label. For example, you can't specify `*prod.example.com` or `prod*.example.com` .
	// - The * can't replace any of the middle labels, for example, marketing.*.example.com.
	// - If you include * in any position other than the leftmost label in a domain name, DNS treats it as an * character (ASCII 42), not as a wildcard.
	//
	// > You can't use the * wildcard for resource records sets that have a type of NS.
	//
	// You can use the * wildcard as the leftmost label in a domain name, for example, `*.example.com` . You can't use an * for one of the middle labels, for example, `marketing.*.example.com` . In addition, the * must replace the entire label; for example, you can't specify `prod*.example.com` .
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// *Latency-based resource record sets only:* The Amazon EC2 Region where you created the resource that this resource record set refers to.
	//
	// The resource typically is an AWS resource, such as an EC2 instance or an ELB load balancer, and is referred to by an IP address or a DNS domain name, depending on the record type.
	//
	// > Although creating latency and latency alias resource record sets in a private hosted zone is allowed, it's not supported.
	//
	// When Amazon Route 53 receives a DNS query for a domain name and type for which you have created latency resource record sets, Route 53 selects the latency resource record set that has the lowest latency between the end user and the associated Amazon EC2 Region. Route 53 then returns the value that is associated with the selected resource record set.
	//
	// Note the following:
	//
	// - You can only specify one `ResourceRecord` per latency resource record set.
	// - You can only create one latency resource record set for each Amazon EC2 Region.
	// - You aren't required to create latency resource record sets for all Amazon EC2 Regions. Route 53 will choose the region with the best latency from among the regions that you create latency resource record sets for.
	// - You can't create non-latency resource record sets that have the same values for the `Name` and `Type` elements as latency resource record sets.
	Region() *string
	SetRegion(val *string)
	// One or more values that correspond with the value that you specified for the `Type` property.
	//
	// For example, if you specified `A` for `Type` , you specify one or more IP addresses in IPv4 format for `ResourceRecords` . For information about the format of values for each record type, see [Supported DNS Resource Record Types](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html) in the *Amazon Route 53 Developer Guide* .
	//
	// Note the following:
	//
	// - You can specify more than one value for all record types except CNAME and SOA.
	// - The maximum length of a value is 4000 characters.
	// - If you're creating an alias record, omit `ResourceRecords` .
	ResourceRecords() *[]*string
	SetResourceRecords(val *[]*string)
	// *Resource record sets that have a routing policy other than simple:* An identifier that differentiates among multiple resource record sets that have the same combination of name and type, such as multiple weighted resource record sets named acme.example.com that have a type of A. In a group of resource record sets that have the same name and type, the value of `SetIdentifier` must be unique for each resource record set.
	//
	// For information about routing policies, see [Choosing a Routing Policy](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html) in the *Amazon Route 53 Developer Guide* .
	SetIdentifier() *string
	SetSetIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The resource record cache time to live (TTL), in seconds. Note the following:.
	//
	// - If you're creating or updating an alias resource record set, omit `TTL` . Amazon Route 53 uses the value of `TTL` for the alias target.
	// - If you're associating this resource record set with a health check (if you're adding a `HealthCheckId` element), we recommend that you specify a `TTL` of 60 seconds or less so clients respond quickly to changes in health status.
	// - All of the resource record sets in a group of weighted resource record sets must have the same value for `TTL` .
	// - If a group of weighted resource record sets includes one or more weighted alias resource record sets for which the alias target is an ELB load balancer, we recommend that you specify a `TTL` of 60 seconds for all of the non-alias weighted resource record sets that have the same name and type. Values other than 60 seconds (the TTL for load balancers) will change the effect of the values that you specify for `Weight` .
	Ttl() *string
	SetTtl(val *string)
	// The DNS record type.
	//
	// For information about different record types and how data is encoded for them, see [Supported DNS Resource Record Types](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html) in the *Amazon Route 53 Developer Guide* .
	//
	// Valid values for basic resource record sets: `A` | `AAAA` | `CAA` | `CNAME` | `DS` | `MX` | `NAPTR` | `NS` | `PTR` | `SOA` | `SPF` | `SRV` | `TXT`
	//
	// Values for weighted, latency, geolocation, and failover resource record sets: `A` | `AAAA` | `CAA` | `CNAME` | `MX` | `NAPTR` | `PTR` | `SPF` | `SRV` | `TXT` . When creating a group of weighted, latency, geolocation, or failover resource record sets, specify the same value for all of the resource record sets in the group.
	//
	// Valid values for multivalue answer resource record sets: `A` | `AAAA` | `MX` | `NAPTR` | `PTR` | `SPF` | `SRV` | `TXT`
	//
	// > SPF records were formerly used to verify the identity of the sender of email messages. However, we no longer recommend that you create resource record sets for which the value of `Type` is `SPF` . RFC 7208, *Sender Policy Framework (SPF) for Authorizing Use of Domains in Email, Version 1* , has been updated to say, "...[I]ts existence and mechanism defined in [RFC4408] have led to some interoperability issues. Accordingly, its use is no longer appropriate for SPF version 1; implementations are not to use it." In RFC 7208, see section 14.1, [The SPF DNS Record Type](https://docs.aws.amazon.com/http://tools.ietf.org/html/rfc7208#section-14.1) .
	//
	// Values for alias resource record sets:
	//
	// - *Amazon API Gateway custom regional APIs and edge-optimized APIs:* `A`
	// - *CloudFront distributions:* `A`
	//
	// If IPv6 is enabled for the distribution, create two resource record sets to route traffic to your distribution, one with a value of `A` and one with a value of `AAAA` .
	// - *Amazon API Gateway environment that has a regionalized subdomain* : `A`
	// - *ELB load balancers:* `A` | `AAAA`
	// - *Amazon S3 buckets:* `A`
	// - *Amazon Virtual Private Cloud interface VPC endpoints* `A`
	// - *Another resource record set in this hosted zone:* Specify the type of the resource record set that you're creating the alias for. All values are supported except `NS` and `SOA` .
	//
	// > If you're creating an alias record that has the same name as the hosted zone (known as the zone apex), you can't route traffic to a record for which the value of `Type` is `CNAME` . This is because the alias record must have the same type as the record you're routing traffic to, and creating a CNAME record for the zone apex isn't supported even for an alias record.
	Type() *string
	SetType(val *string)
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
	// *Weighted resource record sets only:* Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set.
	//
	// Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total. Note the following:
	//
	// - You must specify a value for the `Weight` element for every weighted resource record set.
	// - You can only specify one `ResourceRecord` per weighted resource record set.
	// - You can't create latency, failover, or geolocation resource record sets that have the same values for the `Name` and `Type` elements as weighted resource record sets.
	// - You can create a maximum of 100 weighted resource record sets that have the same values for the `Name` and `Type` elements.
	// - For weighted (but not weighted alias) resource record sets, if you set `Weight` to `0` for a resource record set, Route 53 never responds to queries with the applicable value for that resource record set. However, if you set `Weight` to `0` for all resource record sets that have the same combination of DNS name and type, traffic is routed to all resources with equal probability.
	//
	// The effect of setting `Weight` to `0` is different when you associate health checks with weighted resource record sets. For more information, see [Options for Configuring Route 53 Active-Active and Active-Passive Failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring-options.html) in the *Amazon Route 53 Developer Guide* .
	Weight() *float64
	SetWeight(val *float64)
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

// The jsii proxy struct for CfnRecordSet
type jsiiProxy_CfnRecordSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRecordSet) AliasTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CidrRoutingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) GeoLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) HostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) MultiValueAnswer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) ResourceRecords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSet) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53::RecordSet`.
func NewCfnRecordSet(scope constructs.Construct, id *string, props *CfnRecordSetProps) CfnRecordSet {
	_init_.Initialize()

	if err := validateNewCfnRecordSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecordSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53::RecordSet`.
func NewCfnRecordSet_Override(c CfnRecordSet, scope constructs.Construct, id *string, props *CfnRecordSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetAliasTarget(val interface{}) {
	if err := j.validateSetAliasTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasTarget",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetCidrRoutingConfig(val interface{}) {
	if err := j.validateSetCidrRoutingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrRoutingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetFailover(val *string) {
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetGeoLocation(val interface{}) {
	if err := j.validateSetGeoLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoLocation",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetHealthCheckId(val *string) {
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetHostedZoneId(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetHostedZoneName(val *string) {
	_jsii_.Set(
		j,
		"hostedZoneName",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetMultiValueAnswer(val interface{}) {
	if err := j.validateSetMultiValueAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetResourceRecords(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceRecords",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetSetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetTtl(val *string) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnRecordSet)SetWeight(val *float64) {
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRecordSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRecordSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSet_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnRecordSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53.CfnRecordSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecordSet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRecordSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRecordSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRecordSet) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnRecordSet) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnRecordSet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRecordSet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRecordSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnRecordSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecordSet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

