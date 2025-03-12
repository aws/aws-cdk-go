package awsfms


// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &CfnPolicyProps{
//   	ExcludeResourceTags: jsii.Boolean(false),
//   	PolicyName: jsii.String("policyName"),
//   	RemediationEnabled: jsii.Boolean(false),
//   	SecurityServicePolicyData: &SecurityServicePolicyDataProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		ManagedServiceData: jsii.String("managedServiceData"),
//   		PolicyOption: &PolicyOptionProperty{
//   			NetworkAclCommonPolicy: &NetworkAclCommonPolicyProperty{
//   				NetworkAclEntrySet: &NetworkAclEntrySetProperty{
//   					ForceRemediateForFirstEntries: jsii.Boolean(false),
//   					ForceRemediateForLastEntries: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					FirstEntries: []interface{}{
//   						&NetworkAclEntryProperty{
//   							Egress: jsii.Boolean(false),
//   							Protocol: jsii.String("protocol"),
//   							RuleAction: jsii.String("ruleAction"),
//
//   							// the properties below are optional
//   							CidrBlock: jsii.String("cidrBlock"),
//   							IcmpTypeCode: &IcmpTypeCodeProperty{
//   								Code: jsii.Number(123),
//   								Type: jsii.Number(123),
//   							},
//   							Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   							PortRange: &PortRangeProperty{
//   								From: jsii.Number(123),
//   								To: jsii.Number(123),
//   							},
//   						},
//   					},
//   					LastEntries: []interface{}{
//   						&NetworkAclEntryProperty{
//   							Egress: jsii.Boolean(false),
//   							Protocol: jsii.String("protocol"),
//   							RuleAction: jsii.String("ruleAction"),
//
//   							// the properties below are optional
//   							CidrBlock: jsii.String("cidrBlock"),
//   							IcmpTypeCode: &IcmpTypeCodeProperty{
//   								Code: jsii.Number(123),
//   								Type: jsii.Number(123),
//   							},
//   							Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   							PortRange: &PortRangeProperty{
//   								From: jsii.Number(123),
//   								To: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			NetworkFirewallPolicy: &NetworkFirewallPolicyProperty{
//   				FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   			},
//   			ThirdPartyFirewallPolicy: &ThirdPartyFirewallPolicyProperty{
//   				FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	DeleteAllPolicyResources: jsii.Boolean(false),
//   	ExcludeMap: map[string][]*string{
//   		"account": []*string{
//   			jsii.String("account"),
//   		},
//   		"orgunit": []*string{
//   			jsii.String("orgunit"),
//   		},
//   	},
//   	IncludeMap: map[string][]*string{
//   		"account": []*string{
//   			jsii.String("account"),
//   		},
//   		"orgunit": []*string{
//   			jsii.String("orgunit"),
//   		},
//   	},
//   	PolicyDescription: jsii.String("policyDescription"),
//   	ResourcesCleanUp: jsii.Boolean(false),
//   	ResourceSetIds: []*string{
//   		jsii.String("resourceSetIds"),
//   	},
//   	ResourceTagLogicalOperator: jsii.String("resourceTagLogicalOperator"),
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	ResourceTypeList: []*string{
//   		jsii.String("resourceTypeList"),
//   	},
//   	Tags: []policyTagProperty{
//   		&policyTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html
//
type CfnPolicyProps struct {
	// Used only when tags are specified in the `ResourceTags` property.
	//
	// If this property is `True` , resources with the specified tags are not in scope of the policy. If it's `False` , only resources with the specified tags are in scope of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-excluderesourcetags
	//
	ExcludeResourceTags interface{} `field:"required" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// The name of the AWS Firewall Manager policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Indicates if the policy should be automatically applied to new resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-remediationenabled
	//
	RemediationEnabled interface{} `field:"required" json:"remediationEnabled" yaml:"remediationEnabled"`
	// Details about the security service that is being used to protect the resources.
	//
	// This contains the following settings:
	//
	// - Type - Indicates the service type that the policy uses to protect the resource. For security group policies, Firewall Manager supports one security group for each common policy and for each content audit policy. This is an adjustable limit that you can increase by contacting  .
	//
	// Valid values: `DNS_FIREWALL` | `NETWORK_FIREWALL` | `SECURITY_GROUPS_COMMON` | `SECURITY_GROUPS_CONTENT_AUDIT` | `SECURITY_GROUPS_USAGE_AUDIT` | `SHIELD_ADVANCED` | `THIRD_PARTY_FIREWALL` | `WAFV2` | `WAF`
	// - ManagedServiceData - Details about the service that are specific to the service type, in JSON format.
	//
	// - Example: `DNS_FIREWALL`
	//
	// `"{\"type\":\"DNS_FIREWALL\",\"preProcessRuleGroups\":[{\"ruleGroupId\":\"rslvr-frg-1\",\"priority\":10}],\"postProcessRuleGroups\":[{\"ruleGroupId\":\"rslvr-frg-2\",\"priority\":9911}]}"`
	//
	// > Valid values for `preProcessRuleGroups` are between 1 and 99. Valid values for `postProcessRuleGroups` are between 9901 and 10000.
	// - Example: `NETWORK_FIREWALL` - Centralized deployment model
	//
	// `"{\"type\":\"NETWORK_FIREWALL\",\"awsNetworkFirewallConfig\":{\"networkFirewallStatelessRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateless-rulegroup/test\",\"priority\":1}],\"networkFirewallStatelessDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessFragmentDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessCustomActions\":[{\"actionName\":\"customActionName\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"metricdimensionvalue\"}]}}}],\"networkFirewallStatefulRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateful-rulegroup/test\"}],\"networkFirewallLoggingConfiguration\":{\"logDestinationConfigs\":[{\"logDestinationType\":\"S3\",\"logType\":\"ALERT\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}},{\"logDestinationType\":\"S3\",\"logType\":\"FLOW\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}}],\"overrideExistingConfig\":true}},\"firewallDeploymentModel\":{\"centralizedFirewallDeploymentModel\":{\"centralizedFirewallOrchestrationConfig\":{\"inspectionVpcIds\":[{\"resourceId\":\"vpc-1234\",\"accountId\":\"123456789011\"}],\"firewallCreationConfig\":{\"endpointLocation\":{\"availabilityZoneConfigList\":[{\"availabilityZoneId\":null,\"availabilityZoneName\":\"us-east-1a\",\"allowedIPV4CidrList\":[\"10.0.0.0/28\"]}]}},\"allowedIPV4CidrList\":[]}}}}"`
	//
	// To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html) to `DISTRIBUTED` .
	// - Example: `NETWORK_FIREWALL` - Distributed deployment model with automatic Availability Zone configuration
	//
	// `"{\"type\":\"NETWORK_FIREWALL\",\"networkFirewallStatelessRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateless-rulegroup/test\",\"priority\":1}],\"networkFirewallStatelessDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessFragmentDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessCustomActions\":[{\"actionName\":\"customActionName\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"metricdimensionvalue\"}]}}}],\"networkFirewallStatefulRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateful-rulegroup/test\"}],\"networkFirewallOrchestrationConfig\":{\"singleFirewallEndpointPerVPC\":false,\"allowedIPV4CidrList\":[\"10.0.0.0/28\",\"192.168.0.0/28\"],\"routeManagementAction\":\"OFF\"},\"networkFirewallLoggingConfiguration\":{\"logDestinationConfigs\":[{\"logDestinationType\":\"S3\",\"logType\":\"ALERT\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}},{\"logDestinationType\":\"S3\",\"logType\":\"FLOW\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}}],\"overrideExistingConfig\":true}}"`
	//
	// With automatic Availbility Zone configuration, Firewall Manager chooses which Availability Zones to create the endpoints in. To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html) to `DISTRIBUTED` .
	// - Example: `NETWORK_FIREWALL` - Distributed deployment model with automatic Availability Zone configuration and route management
	//
	// `"{\"type\":\"NETWORK_FIREWALL\",\"networkFirewallStatelessRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateless-rulegroup/test\",\"priority\":1}],\"networkFirewallStatelessDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessFragmentDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessCustomActions\":[{\"actionName\":\"customActionName\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"metricdimensionvalue\"}]}}}],\"networkFirewallStatefulRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateful-rulegroup/test\"}],\"networkFirewallOrchestrationConfig\":{\"singleFirewallEndpointPerVPC\":false,\"allowedIPV4CidrList\":[\"10.0.0.0/28\",\"192.168.0.0/28\"],\"routeManagementAction\":\"MONITOR\",\"routeManagementTargetTypes\":[\"InternetGateway\"]},\"networkFirewallLoggingConfiguration\":{\"logDestinationConfigs\":[{\"logDestinationType\":\"S3\",\"logType\":\"ALERT\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}},{\"logDestinationType\":\"S3\",\"logType\": \"FLOW\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}}],\"overrideExistingConfig\":true}}"`
	//
	// To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html) to `DISTRIBUTED` .
	// - Example: `NETWORK_FIREWALL` - Distributed deployment model with custom Availability Zone configuration
	//
	// `"{\"type\":\"NETWORK_FIREWALL\",\"networkFirewallStatelessRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateless-rulegroup/test\",\"priority\":1}],\"networkFirewallStatelessDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessFragmentDefaultActions\":[\"aws:forward_to_sfe\",\"fragmentcustomactionname\"],\"networkFirewallStatelessCustomActions\":[{\"actionName\":\"customActionName\", \"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"metricdimensionvalue\"}]}}},{\"actionName\":\"fragmentcustomactionname\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"fragmentmetricdimensionvalue\"}]}}}],\"networkFirewallStatefulRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateful-rulegroup/test\"}],\"networkFirewallOrchestrationConfig\":{\"firewallCreationConfig\":{ \"endpointLocation\":{\"availabilityZoneConfigList\":[{\"availabilityZoneName\":\"us-east-1a\",\"allowedIPV4CidrList\":[\"10.0.0.0/28\"]},{\"availabilityZoneName\":\"us-east-1b\",\"allowedIPV4CidrList\":[ \"10.0.0.0/28\"]}]} },\"singleFirewallEndpointPerVPC\":false,\"allowedIPV4CidrList\":null,\"routeManagementAction\":\"OFF\",\"networkFirewallLoggingConfiguration\":{\"logDestinationConfigs\":[{\"logDestinationType\":\"S3\",\"logType\":\"ALERT\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}},{\"logDestinationType\":\"S3\",\"logType\":\"FLOW\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}}],\"overrideExistingConfig\":boolean}}"`
	//
	// With custom Availability Zone configuration, you define which specific Availability Zones to create endpoints in by configuring `firewallCreationConfig` . To configure the Availability Zones in `firewallCreationConfig` , specify either the `availabilityZoneName` or `availabilityZoneId` parameter, not both parameters.
	//
	// To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html) to `DISTRIBUTED` .
	// - Example: `NETWORK_FIREWALL` - Distributed deployment model with custom Availability Zone configuration and route management
	//
	// `"{\"type\":\"NETWORK_FIREWALL\",\"networkFirewallStatelessRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateless-rulegroup/test\",\"priority\":1}],\"networkFirewallStatelessDefaultActions\":[\"aws:forward_to_sfe\",\"customActionName\"],\"networkFirewallStatelessFragmentDefaultActions\":[\"aws:forward_to_sfe\",\"fragmentcustomactionname\"],\"networkFirewallStatelessCustomActions\":[{\"actionName\":\"customActionName\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"metricdimensionvalue\"}]}}},{\"actionName\":\"fragmentcustomactionname\",\"actionDefinition\":{\"publishMetricAction\":{\"dimensions\":[{\"value\":\"fragmentmetricdimensionvalue\"}]}}}],\"networkFirewallStatefulRuleGroupReferences\":[{\"resourceARN\":\"arn:aws:network-firewall:us-east-1:123456789011:stateful-rulegroup/test\"}],\"networkFirewallOrchestrationConfig\":{\"firewallCreationConfig\":{\"endpointLocation\":{\"availabilityZoneConfigList\":[{\"availabilityZoneName\":\"us-east-1a\",\"allowedIPV4CidrList\":[\"10.0.0.0/28\"]},{\"availabilityZoneName\":\"us-east-1b\",\"allowedIPV4CidrList\":[\"10.0.0.0/28\"]}]}},\"singleFirewallEndpointPerVPC\":false,\"allowedIPV4CidrList\":null,\"routeManagementAction\":\"MONITOR\",\"routeManagementTargetTypes\":[\"InternetGateway\"],\"routeManagementConfig\":{\"allowCrossAZTrafficIfNoEndpoint\":true}},\"networkFirewallLoggingConfiguration\":{\"logDestinationConfigs\":[{\"logDestinationType\":\"S3\",\"logType\":\"ALERT\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}},{\"logDestinationType\":\"S3\",\"logType\":\"FLOW\",\"logDestination\":{\"bucketName\":\"s3-bucket-name\"}}],\"overrideExistingConfig\":boolean}}"`
	//
	// To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html) to `DISTRIBUTED` .
	// - Example: `THIRD_PARTY_FIREWALL` - Palo Alto Networks Cloud Next-Generation Firewall centralized deployment model
	//
	// `"{ \"type\":\"THIRD_PARTY_FIREWALL\", \"thirdPartyFirewall\":\"PALO_ALTO_NETWORKS_CLOUD_NGFW\", \"thirdPartyFirewallConfig\":{ \"thirdPartyFirewallPolicyList\":[\"global-1\"] },\"firewallDeploymentModel\":{\"centralizedFirewallDeploymentModel\":{\"centralizedFirewallOrchestrationConfig\":{\"inspectionVpcIds\":[{\"resourceId\":\"vpc-1234\",\"accountId\":\"123456789011\"}],\"firewallCreationConfig\":{\"endpointLocation\":{\"availabilityZoneConfigList\":[{\"availabilityZoneId\":null,\"availabilityZoneName\":\"us-east-1a\",\"allowedIPV4CidrList\":[\"10.0.0.0/28\"]}]}},\"allowedIPV4CidrList\":[]}}}}"`
	//
	// To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-thirdpartyfirewallpolicy.html) to `CENTRALIZED` .
	// - Example: `THIRD_PARTY_FIREWALL` - Palo Alto Networks Cloud Next-Generation Firewall distributed deployment model
	//
	// `"{\"type\":\"THIRD_PARTY_FIREWALL\",\"thirdPartyFirewall\":\"PALO_ALTO_NETWORKS_CLOUD_NGFW\",\"thirdPartyFirewallConfig\":{\"thirdPartyFirewallPolicyList\":[\"global-1\"] },\"firewallDeploymentModel\":{ \"distributedFirewallDeploymentModel\":{ \"distributedFirewallOrchestrationConfig\":{\"firewallCreationConfig\":{\"endpointLocation\":{ \"availabilityZoneConfigList\":[ {\"availabilityZoneName\":\"${AvailabilityZone}\" } ] } }, \"allowedIPV4CidrList\":[ ] } } } }"`
	//
	// To use the distributed deployment model, you must set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-thirdpartyfirewallpolicy.html) to `DISTRIBUTED` .
	// - Specification for `SHIELD_ADVANCED` for Amazon CloudFront distributions
	//
	// `"{\"type\":\"SHIELD_ADVANCED\",\"automaticResponseConfiguration\": {\"automaticResponseStatus\":\"ENABLED|IGNORED|DISABLED\", \"automaticResponseAction\":\"BLOCK|COUNT\"}, \"overrideCustomerWebaclClassic\":true|false}"`
	//
	// For example: `"{\"type\":\"SHIELD_ADVANCED\",\"automaticResponseConfiguration\": {\"automaticResponseStatus\":\"ENABLED\", \"automaticResponseAction\":\"COUNT\"}}"`
	//
	// The default value for `automaticResponseStatus` is `IGNORED` . The value for `automaticResponseAction` is only required when `automaticResponseStatus` is set to `ENABLED` . The default value for `overrideCustomerWebaclClassic` is `false` .
	//
	// For other resource types that you can protect with a Shield Advanced policy, this `ManagedServiceData` configuration is an empty string.
	// - Example: `WAFV2`
	//
	// `"{\"type\":\"WAFV2\",\"preProcessRuleGroups\":[{\"ruleGroupArn\":null,\"overrideAction\":{\"type\":\"NONE\"},\"managedRuleGroupIdentifier\":{\"version\":null,\"vendorName\":\"AWS\",\"managedRuleGroupName\":\"AWSManagedRulesAmazonIpReputationList\"},\"ruleGroupType\":\"ManagedRuleGroup\",\"excludeRules\":[{\"name\":\"NoUserAgent_HEADER\"}]}],\"postProcessRuleGroups\":[],\"defaultAction\":{\"type\":\"ALLOW\"},\"overrideCustomerWebACLAssociation\":false,\"loggingConfiguration\":{\"logDestinationConfigs\":[\"arn:aws:firehose:us-west-2:12345678912:deliverystream/aws-waf-logs-fms-admin-destination\"],\"redactedFields\":[{\"redactedFieldType\":\"SingleHeader\",\"redactedFieldValue\":\"Cookies\"},{\"redactedFieldType\":\"Method\"}]}}"`
	//
	// In the `loggingConfiguration` , you can specify one `logDestinationConfigs` , you can optionally provide up to 20 `redactedFields` , and the `RedactedFieldType` must be one of `URI` , `QUERY_STRING` , `HEADER` , or `METHOD` .
	// - Example: `AWS WAF Classic`
	//
	// `"{\"type\": \"WAF\", \"ruleGroups\": [{\"id\":\"12345678-1bcd-9012-efga-0987654321ab\", \"overrideAction\" : {\"type\": \"COUNT\"}}], \"defaultAction\": {\"type\": \"BLOCK\"}}"`
	// - Example: `WAFV2` - AWS Firewall Manager support for AWS WAF managed rule group versioning
	//
	// `"{\"type\":\"WAFV2\",\"preProcessRuleGroups\":[{\"ruleGroupArn\":null,\"overrideAction\":{\"type\":\"NONE\"},\"managedRuleGroupIdentifier\":{\"versionEnabled\":true,\"version\":\"Version_2.0\",\"vendorName\":\"AWS\",\"managedRuleGroupName\":\"AWSManagedRulesCommonRuleSet\"},\"ruleGroupType\":\"ManagedRuleGroup\",\"excludeRules\":[{\"name\":\"NoUserAgent_HEADER\"}]}],\"postProcessRuleGroups\":[],\"defaultAction\":{\"type\":\"ALLOW\"},\"overrideCustomerWebACLAssociation\":false,\"loggingConfiguration\":{\"logDestinationConfigs\":[\"arn:aws:firehose:us-west-2:12345678912:deliverystream/aws-waf-logs-fms-admin-destination\"],\"redactedFields\":[{\"redactedFieldType\":\"SingleHeader\",\"redactedFieldValue\":\"Cookies\"},{\"redactedFieldType\":\"Method\"}]}}"`
	//
	// To use a specific version of a AWS WAF managed rule group in your Firewall Manager policy, you must set `versionEnabled` to `true` , and set `version` to the version you'd like to use. If you don't set `versionEnabled` to `true` , or if you omit `versionEnabled` , then Firewall Manager uses the default version of the AWS WAF managed rule group.
	// - Example: `SECURITY_GROUPS_COMMON`
	//
	// `"{\"type\":\"SECURITY_GROUPS_COMMON\",\"revertManualSecurityGroupChanges\":false,\"exclusiveResourceSecurityGroupManagement\":false, \"applyToAllEC2InstanceENIs\":false,\"securityGroups\":[{\"id\":\" sg-000e55995d61a06bd\"}]}"`
	// - Example: Shared VPCs. Apply the preceding policy to resources in shared VPCs as well as to those in VPCs that the account owns
	//
	// `"{\"type\":\"SECURITY_GROUPS_COMMON\",\"revertManualSecurityGroupChanges\":false,\"exclusiveResourceSecurityGroupManagement\":false, \"applyToAllEC2InstanceENIs\":false,\"includeSharedVPC\":true,\"securityGroups\":[{\"id\":\" sg-000e55995d61a06bd\"}]}"`
	// - Example: `SECURITY_GROUPS_CONTENT_AUDIT`
	//
	// `"{\"type\":\"SECURITY_GROUPS_CONTENT_AUDIT\",\"securityGroups\":[{\"id\":\"sg-000e55995d61a06bd\"}],\"securityGroupAction\":{\"type\":\"ALLOW\"}}"`
	//
	// The security group action for content audit can be `ALLOW` or `DENY` . For `ALLOW` , all in-scope security group rules must be within the allowed range of the policy's security group rules. For `DENY` , all in-scope security group rules must not contain a value or a range that matches a rule value or range in the policy security group.
	// - Example: `SECURITY_GROUPS_USAGE_AUDIT`
	//
	// `"{\"type\":\"SECURITY_GROUPS_USAGE_AUDIT\",\"deleteUnusedSecurityGroups\":true,\"coalesceRedundantSecurityGroups\":true}"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-securityservicepolicydata
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-deleteallpolicyresources
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-excludemap
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-includemap
	//
	IncludeMap interface{} `field:"optional" json:"includeMap" yaml:"includeMap"`
	// Your description of the AWS Firewall Manager policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-policydescription
	//
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// Indicates whether AWS Firewall Manager should automatically remove protections from resources that leave the policy scope and clean up resources that Firewall Manager is managing for accounts when those accounts leave policy scope.
	//
	// For example, Firewall Manager will disassociate a Firewall Manager managed web ACL from a protected customer resource when the customer resource leaves policy scope.
	//
	// By default, Firewall Manager doesn't remove protections or delete Firewall Manager managed resources.
	//
	// This option is not available for Shield Advanced or AWS WAF Classic policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-resourcescleanup
	//
	ResourcesCleanUp interface{} `field:"optional" json:"resourcesCleanUp" yaml:"resourcesCleanUp"`
	// The unique identifiers of the resource sets used by the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-resourcesetids
	//
	ResourceSetIds *[]*string `field:"optional" json:"resourceSetIds" yaml:"resourceSetIds"`
	// Specifies whether to combine multiple resource tags with AND, so that a resource must have all tags to be included or excluded, or OR, so that a resource must have at least one tag.
	//
	// Default: `AND`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-resourcetaglogicaloperator
	//
	ResourceTagLogicalOperator *string `field:"optional" json:"resourceTagLogicalOperator" yaml:"resourceTagLogicalOperator"`
	// An array of `ResourceTag` objects, used to explicitly include resources in the policy scope or explicitly exclude them.
	//
	// If this isn't set, then tags aren't used to modify policy scope. See also `ExcludeResourceTags` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The type of resource protected by or in scope of the policy.
	//
	// This is in the format shown in the [AWS Resource Types Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html) . To apply this policy to multiple resource types, specify a resource type of `ResourceTypeList` and then specify the resource types in a `ResourceTypeList` .
	//
	// The following are valid resource types for each Firewall Manager policy type:
	//
	// - AWS WAF Classic - `AWS::ApiGateway::Stage` , `AWS::CloudFront::Distribution` , and `AWS::ElasticLoadBalancingV2::LoadBalancer` .
	// - AWS WAF - `AWS::ApiGateway::Stage` , `AWS::ElasticLoadBalancingV2::LoadBalancer` , and `AWS::CloudFront::Distribution` .
	// - Shield Advanced - `AWS::ElasticLoadBalancingV2::LoadBalancer` , `AWS::ElasticLoadBalancing::LoadBalancer` , `AWS::EC2::EIP` , and `AWS::CloudFront::Distribution` .
	// - Network ACL - `AWS::EC2::Subnet` .
	// - Security group usage audit - `AWS::EC2::SecurityGroup` .
	// - Security group content audit - `AWS::EC2::SecurityGroup` , `AWS::EC2::NetworkInterface` , and `AWS::EC2::Instance` .
	// - DNS Firewall, AWS Network Firewall , and third-party firewall - `AWS::EC2::VPC` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// An array of `ResourceType` objects.
	//
	// Use this only to specify multiple resource types. To specify a single resource type, use `ResourceType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-resourcetypelist
	//
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// A collection of key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-policy.html#cfn-fms-policy-tags
	//
	Tags *[]*CfnPolicy_PolicyTagProperty `field:"optional" json:"tags" yaml:"tags"`
}

