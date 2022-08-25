package awsfms


// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &cfnPolicyProps{
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
//   }
//
type CfnPolicyProps struct {
	// Used only when tags are specified in the `ResourceTags` property.
	//
	// If this property is `True` , resources with the specified tags are not in scope of the policy. If it's `False` , only resources with the specified tags are in scope of the policy.
	ExcludeResourceTags interface{} `field:"required" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// The name of the AWS Firewall Manager policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Indicates if the policy should be automatically applied to new resources.
	RemediationEnabled interface{} `field:"required" json:"remediationEnabled" yaml:"remediationEnabled"`
	// The type of resource protected by or in scope of the policy.
	//
	// This is in the format shown in the [AWS Resource Types Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) . To apply this policy to multiple resource types, specify a resource type of `ResourceTypeList` and then specify the resource types in a `ResourceTypeList` .
	//
	// For AWS WAF and Shield Advanced, example resource types include `AWS::ElasticLoadBalancingV2::LoadBalancer` and `AWS::CloudFront::Distribution` . For a security group common policy, valid values are `AWS::EC2::NetworkInterface` and `AWS::EC2::Instance` . For a security group content audit policy, valid values are `AWS::EC2::SecurityGroup` , `AWS::EC2::NetworkInterface` , and `AWS::EC2::Instance` . For a security group usage audit policy, the value is `AWS::EC2::SecurityGroup` . For an AWS Network Firewall policy or DNS Firewall policy, the value is `AWS::EC2::VPC` .
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
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
	SecurityServicePolicyData interface{} `field:"required" json:"securityServicePolicyData" yaml:"securityServicePolicyData"`
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
	DeleteAllPolicyResources interface{} `field:"optional" json:"deleteAllPolicyResources" yaml:"deleteAllPolicyResources"`
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
	ExcludeMap interface{} `field:"optional" json:"excludeMap" yaml:"excludeMap"`
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
	IncludeMap interface{} `field:"optional" json:"includeMap" yaml:"includeMap"`
	// Indicates whether AWS Firewall Manager should automatically remove protections from resources that leave the policy scope and clean up resources that Firewall Manager is managing for accounts when those accounts leave policy scope.
	//
	// For example, Firewall Manager will disassociate a Firewall Manager managed web ACL from a protected customer resource when the customer resource leaves policy scope.
	//
	// By default, Firewall Manager doesn't remove protections or delete Firewall Manager managed resources.
	//
	// This option is not available for Shield Advanced or AWS WAF Classic policies.
	ResourcesCleanUp interface{} `field:"optional" json:"resourcesCleanUp" yaml:"resourcesCleanUp"`
	// An array of `ResourceTag` objects, used to explicitly include resources in the policy scope or explicitly exclude them.
	//
	// If this isn't set, then tags aren't used to modify policy scope. See also `ExcludeResourceTags` .
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// An array of `ResourceType` objects.
	//
	// Use this only to specify multiple resource types. To specify a single resource type, use `ResourceType` .
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// A collection of key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	Tags *[]*CfnPolicy_PolicyTagProperty `field:"optional" json:"tags" yaml:"tags"`
}

