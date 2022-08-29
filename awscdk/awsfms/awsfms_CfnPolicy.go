package awsfms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::FMS::Policy`.
//
// An AWS Firewall Manager policy.
//
// Firewall Manager provides the following types of policies:
//
// - An AWS Shield Advanced policy, which applies Shield Advanced protection to specified accounts and resources.
// - An AWS WAF policy (type WAFV2), which defines rule groups to run first in the corresponding AWS WAF web ACL and rule groups to run last in the web ACL.
// - An AWS WAF Classic policy, which defines a rule group. AWS WAF Classic doesn't support rule groups in Amazon CloudFront , so, to create AWS WAF Classic policies through CloudFront , you first need to create your rule groups outside of CloudFront .
// - A security group policy, which manages VPC security groups across your AWS organization.
// - An AWS Network Firewall policy, which provides firewall rules to filter network traffic in specified Amazon VPCs.
// - A DNS Firewall policy, which provides Amazon Route 53 Resolver DNS Firewall rules to filter DNS queries for specified Amazon VPCs.
//
// Each policy is specific to one of the types. If you want to enforce more than one policy type across accounts, create multiple policies. You can create multiple policies for each type.
//
// These policies require some setup to use. For more information, see the sections on prerequisites and getting started under [AWS Firewall Manager](https://docs.aws.amazon.com/waf/latest/developerguide/fms-prereq.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicy := awscdk.Aws_fms.NewCfnPolicy(this, jsii.String("MyCfnPolicy"), &cfnPolicyProps{
//   	excludeResourceTags: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	remediationEnabled: jsii.Boolean(false),
//   	resourceType: jsii.String("resourceType"),
//   	securityServicePolicyData: &securityServicePolicyDataProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		managedServiceData: jsii.String("managedServiceData"),
//   		policyOption: &policyOptionProperty{
//   			networkFirewallPolicy: &networkFirewallPolicyProperty{
//   				firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   			},
//   			thirdPartyFirewallPolicy: &thirdPartyFirewallPolicyProperty{
//   				firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	deleteAllPolicyResources: jsii.Boolean(false),
//   	excludeMap: map[string][]*string{
//   		"account": []*string{
//   			jsii.String("account"),
//   		},
//   		"orgunit": []*string{
//   			jsii.String("orgunit"),
//   		},
//   	},
//   	includeMap: map[string][]*string{
//   		"account": []*string{
//   			jsii.String("account"),
//   		},
//   		"orgunit": []*string{
//   			jsii.String("orgunit"),
//   		},
//   	},
//   	resourcesCleanUp: jsii.Boolean(false),
//   	resourceTags: []interface{}{
//   		&resourceTagProperty{
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceTypeList: []*string{
//   		jsii.String("resourceTypeList"),
//   	},
//   	tags: []policyTagProperty{
//   		&policyTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the policy.
	AttrArn() *string
	// The ID of the policy.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Used when deleting a policy. If `true` , Firewall Manager performs cleanup according to the policy type.
	//
	// For AWS WAF and Shield Advanced policies, Firewall Manager does the following:
	//
	// - Deletes rule groups created by Firewall Manager
	// - Removes web ACLs from in-scope resources
	// - Deletes web ACLs that contain no rules or rule groups
	//
	// For security group policies, Firewall Manager does the following for each security group in the policy:
	//
	// - Disassociates the security group from in-scope resources
	// - Deletes the security group if it was created through Firewall Manager and if it's no longer associated with any resources through another policy
	//
	// After the cleanup, in-scope resources are no longer protected by web ACLs in this policy. Protection of out-of-scope resources remains unchanged. Scope is determined by tags that you create and accounts that you associate with the policy. When creating the policy, if you specify that only resources in specific accounts or with specific tags are in scope of the policy, those accounts and resources are handled by the policy. All others are out of scope. If you don't specify tags or accounts, all resources are in scope.
	DeleteAllPolicyResources() interface{}
	SetDeleteAllPolicyResources(val interface{})
	// Specifies the AWS account IDs and AWS Organizations organizational units (OUs) to exclude from the policy.
	//
	// Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.
	//
	// You can specify inclusions or exclusions, but not both. If you specify an `IncludeMap` , AWS Firewall Manager applies the policy to all accounts specified by the `IncludeMap` , and does not evaluate any `ExcludeMap` specifications. If you do not specify an `IncludeMap` , then Firewall Manager applies the policy to all accounts except for those specified by the `ExcludeMap` .
	//
	// You can specify account IDs, OUs, or a combination:
	//
	// - Specify account IDs by setting the key to `ACCOUNT` . For example, the following is a valid map: `{“ACCOUNT” : [“accountID1”, “accountID2”]}` .
	// - Specify OUs by setting the key to `ORGUNIT` . For example, the following is a valid map: `{“ORGUNIT” : [“ouid111”, “ouid112”]}` .
	// - Specify accounts and OUs together in a single map, separated with a comma. For example, the following is a valid map: `{“ACCOUNT” : [“accountID1”, “accountID2”], “ORGUNIT” : [“ouid111”, “ouid112”]}` .
	ExcludeMap() interface{}
	SetExcludeMap(val interface{})
	// Used only when tags are specified in the `ResourceTags` property.
	//
	// If this property is `True` , resources with the specified tags are not in scope of the policy. If it's `False` , only resources with the specified tags are in scope of the policy.
	ExcludeResourceTags() interface{}
	SetExcludeResourceTags(val interface{})
	// Specifies the AWS account IDs and AWS Organizations organizational units (OUs) to include in the policy.
	//
	// Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.
	//
	// You can specify inclusions or exclusions, but not both. If you specify an `IncludeMap` , AWS Firewall Manager applies the policy to all accounts specified by the `IncludeMap` , and does not evaluate any `ExcludeMap` specifications. If you do not specify an `IncludeMap` , then Firewall Manager applies the policy to all accounts except for those specified by the `ExcludeMap` .
	//
	// You can specify account IDs, OUs, or a combination:
	//
	// - Specify account IDs by setting the key to `ACCOUNT` . For example, the following is a valid map: `{“ACCOUNT” : [“accountID1”, “accountID2”]}` .
	// - Specify OUs by setting the key to `ORGUNIT` . For example, the following is a valid map: `{“ORGUNIT” : [“ouid111”, “ouid112”]}` .
	// - Specify accounts and OUs together in a single map, separated with a comma. For example, the following is a valid map: `{“ACCOUNT” : [“accountID1”, “accountID2”], “ORGUNIT” : [“ouid111”, “ouid112”]}` .
	IncludeMap() interface{}
	SetIncludeMap(val interface{})
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
	// The name of the AWS Firewall Manager policy.
	PolicyName() *string
	SetPolicyName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Indicates if the policy should be automatically applied to new resources.
	RemediationEnabled() interface{}
	SetRemediationEnabled(val interface{})
	// Indicates whether AWS Firewall Manager should automatically remove protections from resources that leave the policy scope and clean up resources that Firewall Manager is managing for accounts when those accounts leave policy scope.
	//
	// For example, Firewall Manager will disassociate a Firewall Manager managed web ACL from a protected customer resource when the customer resource leaves policy scope.
	//
	// By default, Firewall Manager doesn't remove protections or delete Firewall Manager managed resources.
	//
	// This option is not available for Shield Advanced or AWS WAF Classic policies.
	ResourcesCleanUp() interface{}
	SetResourcesCleanUp(val interface{})
	// An array of `ResourceTag` objects, used to explicitly include resources in the policy scope or explicitly exclude them.
	//
	// If this isn't set, then tags aren't used to modify policy scope. See also `ExcludeResourceTags` .
	ResourceTags() interface{}
	SetResourceTags(val interface{})
	// The type of resource protected by or in scope of the policy.
	//
	// This is in the format shown in the [AWS Resource Types Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) . To apply this policy to multiple resource types, specify a resource type of `ResourceTypeList` and then specify the resource types in a `ResourceTypeList` .
	//
	// For AWS WAF and Shield Advanced, example resource types include `AWS::ElasticLoadBalancingV2::LoadBalancer` and `AWS::CloudFront::Distribution` . For a security group common policy, valid values are `AWS::EC2::NetworkInterface` and `AWS::EC2::Instance` . For a security group content audit policy, valid values are `AWS::EC2::SecurityGroup` , `AWS::EC2::NetworkInterface` , and `AWS::EC2::Instance` . For a security group usage audit policy, the value is `AWS::EC2::SecurityGroup` . For an AWS Network Firewall policy or DNS Firewall policy, the value is `AWS::EC2::VPC` .
	ResourceType() *string
	SetResourceType(val *string)
	// An array of `ResourceType` objects.
	//
	// Use this only to specify multiple resource types. To specify a single resource type, use `ResourceType` .
	ResourceTypeList() *[]*string
	SetResourceTypeList(val *[]*string)
	// Details about the security service that is being used to protect the resources.
	//
	// This contains the following settings:
	//
	// - Type - Indicates the service type that the policy uses to protect the resource. For security group policies, Firewall Manager supports one security group for each common policy and for each content audit policy. This is an adjustable limit that you can increase by contacting AWS Support .
	//
	// Valid values: `DNS_FIREWALL` | `NETWORK_FIREWALL` | `SECURITY_GROUPS_COMMON` | `SECURITY_GROUPS_CONTENT_AUDIT` | `SECURITY_GROUPS_USAGE_AUDIT` | `SHIELD_ADVANCED` | `WAFV2` | `WAF`
	// - ManagedServiceData - Details about the service that are specific to the service type, in JSON format.
	//
	// - Example: `DNS_FIREWALL`
	//
	// `"ManagedServiceData": "{ \"type\": \"DNS_FIREWALL\", \"preProcessRuleGroups\": [{\"ruleGroupId\": \"rslvr-frg-123456\", \"priority\": 11}], \"postProcessRuleGroups\": [{\"ruleGroupId\": \"rslvr-frg-123456\", \"priority\": 9902}]}"`
	// - Example: `NETWORK_FIREWALL`
	//
	// `"ManagedServiceData":"{\"type\":\"NETWORK_FIREWALL\",\"networkFirewallStatelessRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:000000000000:stateless-rulegroup\/example\",\"priority\":1}],\"networkFirewallStatelessDefaultActions\":[\"aws:drop\",\"example\"],\"networkFirewallStatelessFragmentDefaultActions\":[\"aws:drop\",\"example\"],\"networkFirewallStatelessCustomActions\":[{\"actionName\":\"example\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"example\"}]}}}],\"networkFirewallStatefulRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:000000000000:stateful-rulegroup\/example\"}],\"networkFirewallOrchestrationConfig\":{\"singleFirewallEndpointPerVPC\":false,\"allowedIPV4CidrList\":[]}}"`
	// - Example: `SECURITY_GROUPS_COMMON`
	//
	// `"SecurityServicePolicyData":{"Type":"SECURITY_GROUPS_COMMON","ManagedServiceData":"{\"type\":\"SECURITY_GROUPS_COMMON\",\"revertManualSecurityGroupChanges\":false,\"exclusiveResourceSecurityGroupManagement\":false,\"securityGroups\":[{\"id\":\" sg-000e55995d61a06bd\"}]}"},"RemediationEnabled":false,"ResourceType":"AWS::EC2::NetworkInterface"}`
	// - Example: `SECURITY_GROUPS_CONTENT_AUDIT`
	//
	// `"SecurityServicePolicyData":{"Type":"SECURITY_GROUPS_CONTENT_AUDIT","ManagedServiceData":"{\"type\":\"SECURITY_GROUPS_CONTENT_AUDIT\",\"securityGroups\":[{\"id\":\" sg-000e55995d61a06bd \"}],\"securityGroupAction\":{\"type\":\"ALLOW\"}}"},"RemediationEnabled":false,"ResourceType":"AWS::EC2::NetworkInterface"}`
	//
	// The security group action for content audit can be `ALLOW` or `DENY` . For `ALLOW` , all in-scope security group rules must be within the allowed range of the policy's security group rules. For `DENY` , all in-scope security group rules must not contain a value or a range that matches a rule value or range in the policy security group.
	// - Example: `SECURITY_GROUPS_USAGE_AUDIT`
	//
	// `"SecurityServicePolicyData":{"Type":"SECURITY_GROUPS_USAGE_AUDIT","ManagedServiceData":"{\"type\":\"SECURITY_GROUPS_USAGE_AUDIT\",\"deleteUnusedSecurityGroups\":true,\"coalesceRedundantSecurityGroups\":true}"},"RemediationEnabled":false,"Resou rceType":"AWS::EC2::SecurityGroup"}`
	// - Specification for `SHIELD_ADVANCED` for Amazon CloudFront distributions
	//
	// `"ManagedServiceData": "{\"type\": \"SHIELD_ADVANCED\", \"automaticResponseConfiguration\": {\"automaticResponseStatus\":\"ENABLED|IGNORED|DISABLED\", \"automaticResponseAction\":\"BLOCK|COUNT\"}, \"overrideCustomerWebaclClassic\":true|false}"`
	//
	// For example: `"ManagedServiceData": "{\"type\":\"SHIELD_ADVANCED\",\"automaticResponseConfiguration\": {\"automaticResponseStatus\":\"ENABLED\", \"automaticResponseAction\":\"COUNT\"}}"`
	//
	// The default value for `automaticResponseStatus` is `IGNORED` . The value for `automaticResponseAction` is only required when `automaticResponseStatus` is set to `ENABLED` . The default value for `overrideCustomerWebaclClassic` is `false` .
	//
	// For other resource types that you can protect with a Shield Advanced policy, this `ManagedServiceData` configuration is an empty string.
	// - Example: `WAFV2`
	//
	// `"ManagedServiceData": "{\"type\":\"WAFV2\",\"preProcessRuleGroups\":[{\"ruleGroupArn\":null,\"overrideAction\":{\"type\":\"NONE\"},\"managedRuleGroupIdentifier\":{\"version\":null,\"vendorName\":\"AWS\",\"managedRuleGroupName\":\"AWSManagedRulesAmazonIpReputationList\"},\"ruleGroupType\":\"ManagedRuleGroup\",\"excludeRules\":[]}],\"postProcessRuleGroups\":[],\"defaultAction\":{\"type\":\"ALLOW\"},\"overrideCustomerWebACLAssociation\":false,\"loggingConfiguration\":{\"logDestinationConfigs\":[\"arn:aws:firehose:us-west-2:12345678912:deliverystream/aws-waf-logs-fms-admin-destination\"],\"redactedFields\":[{\"redactedFieldType\":\"SingleHeader\",\"redactedFieldValue\":\"Cookies\"},{\"redactedFieldType\":\"Method\"}]}}"`
	//
	// In the `loggingConfiguration` , you can specify one `logDestinationConfigs` , you can optionally provide up to 20 `redactedFields` , and the `RedactedFieldType` must be one of `URI` , `QUERY_STRING` , `HEADER` , or `METHOD` .
	// - Example: `WAF Classic`
	//
	// `"ManagedServiceData": "{\"type\": \"WAF\", \"ruleGroups\": [{\"id\":\"12345678-1bcd-9012-efga-0987654321ab\", \"overrideAction\" : {\"type\": \"COUNT\"}}],\"defaultAction\": {\"type\": \"BLOCK\"}}`
	//
	// AWS WAF Classic doesn't support rule groups in CloudFront . To create a WAF Classic policy through CloudFormation, create your rule groups outside of CloudFront , then provide the rule group IDs in the WAF managed service data specification.
	SecurityServicePolicyData() interface{}
	SetSecurityServicePolicyData(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A collection of key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	Tags() *[]*CfnPolicy_PolicyTagProperty
	SetTags(val *[]*CfnPolicy_PolicyTagProperty)
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

// The jsii proxy struct for CfnPolicy
type jsiiProxy_CfnPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPolicy) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) DeleteAllPolicyResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) ExcludeMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) ExcludeResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) IncludeMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) RemediationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) ResourcesCleanUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesCleanUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) ResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) ResourceTypeList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) SecurityServicePolicyData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityServicePolicyData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) Tags() *[]*CfnPolicy_PolicyTagProperty {
	var returns *[]*CfnPolicy_PolicyTagProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicy) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::FMS::Policy`.
func NewCfnPolicy(scope constructs.Construct, id *string, props *CfnPolicyProps) CfnPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fms.CfnPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FMS::Policy`.
func NewCfnPolicy_Override(c CfnPolicy, scope constructs.Construct, id *string, props *CfnPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fms.CfnPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPolicy) SetDeleteAllPolicyResources(val interface{}) {
	_jsii_.Set(
		j,
		"deleteAllPolicyResources",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetExcludeMap(val interface{}) {
	_jsii_.Set(
		j,
		"excludeMap",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetExcludeResourceTags(val interface{}) {
	_jsii_.Set(
		j,
		"excludeResourceTags",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetIncludeMap(val interface{}) {
	_jsii_.Set(
		j,
		"includeMap",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetRemediationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"remediationEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetResourcesCleanUp(val interface{}) {
	_jsii_.Set(
		j,
		"resourcesCleanUp",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetResourceTags(val interface{}) {
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetResourceTypeList(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypeList",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetSecurityServicePolicyData(val interface{}) {
	_jsii_.Set(
		j,
		"securityServicePolicyData",
		val,
	)
}

func (j *jsiiProxy_CfnPolicy) SetTags(val *[]*CfnPolicy_PolicyTagProperty) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fms.CfnPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fms.CfnPolicy",
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
func CfnPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fms.CfnPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_fms.CfnPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

