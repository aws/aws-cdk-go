package mixinsawsguardduty


// Represents a map of finding properties that match specified conditions and values when querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var criterion interface{}
//
//   findingCriteriaProperty := &FindingCriteriaProperty{
//   	Criterion: criterion,
//   	ItemType: &ConditionProperty{
//   		Eq: []*string{
//   			jsii.String("eq"),
//   		},
//   		EqualTo: []*string{
//   			jsii.String("equalTo"),
//   		},
//   		GreaterThan: jsii.Number(123),
//   		GreaterThanOrEqual: jsii.Number(123),
//   		Gt: jsii.Number(123),
//   		Gte: jsii.Number(123),
//   		LessThan: jsii.Number(123),
//   		LessThanOrEqual: jsii.Number(123),
//   		Lt: jsii.Number(123),
//   		Lte: jsii.Number(123),
//   		Neq: []*string{
//   			jsii.String("neq"),
//   		},
//   		NotEquals: []*string{
//   			jsii.String("notEquals"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-findingcriteria.html
//
type CfnFilterPropsMixin_FindingCriteriaProperty struct {
	// Represents a map of finding properties that match specified conditions and values when querying findings.
	//
	// For information about JSON criterion mapping to their console equivalent, see [Finding criteria](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_filter-findings.html#filter_criteria) . The following are the available criterion:
	//
	// - accountId
	// - id
	// - region
	// - severity
	//
	// To filter on the basis of severity, the API and AWS CLI use the following input list for the `FindingCriteria` condition:
	//
	// - *Low* : `["1", "2", "3"]`
	// - *Medium* : `["4", "5", "6"]`
	// - *High* : `["7", "8", "9"]`
	//
	// For more information, see [Severity levels for GuardDuty findings](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings.html#guardduty_findings-severity) in the *Amazon GuardDuty User Guide* .
	// - type
	// - updatedAt
	//
	// Type: ISO 8601 string format: `YYYY-MM-DDTHH:MM:SS.SSSZ` or `YYYY-MM-DDTHH:MM:SSZ` depending on whether the value contains milliseconds.
	// - resource.accessKeyDetails.accessKeyId
	// - resource.accessKeyDetails.principalId
	// - resource.accessKeyDetails.userName
	// - resource.accessKeyDetails.userType
	// - resource.instanceDetails.iamInstanceProfile.id
	// - resource.instanceDetails.imageId
	// - resource.instanceDetails.instanceId
	// - resource.instanceDetails.tags.key
	// - resource.instanceDetails.tags.value
	// - resource.instanceDetails.networkInterfaces.ipv6Addresses
	// - resource.instanceDetails.networkInterfaces.privateIpAddresses.privateIpAddress
	// - resource.instanceDetails.networkInterfaces.publicDnsName
	// - resource.instanceDetails.networkInterfaces.publicIp
	// - resource.instanceDetails.networkInterfaces.securityGroups.groupId
	// - resource.instanceDetails.networkInterfaces.securityGroups.groupName
	// - resource.instanceDetails.networkInterfaces.subnetId
	// - resource.instanceDetails.networkInterfaces.vpcId
	// - resource.instanceDetails.outpostArn
	// - resource.resourceType
	// - resource.s3BucketDetails.publicAccess.effectivePermissions
	// - resource.s3BucketDetails.name
	// - resource.s3BucketDetails.tags.key
	// - resource.s3BucketDetails.tags.value
	// - resource.s3BucketDetails.type
	// - service.action.actionType
	// - service.action.awsApiCallAction.api
	// - service.action.awsApiCallAction.callerType
	// - service.action.awsApiCallAction.errorCode
	// - service.action.awsApiCallAction.remoteIpDetails.city.cityName
	// - service.action.awsApiCallAction.remoteIpDetails.country.countryName
	// - service.action.awsApiCallAction.remoteIpDetails.ipAddressV4
	// - service.action.awsApiCallAction.remoteIpDetails.ipAddressV6
	// - service.action.awsApiCallAction.remoteIpDetails.organization.asn
	// - service.action.awsApiCallAction.remoteIpDetails.organization.asnOrg
	// - service.action.awsApiCallAction.serviceName
	// - service.action.dnsRequestAction.domain
	// - service.action.dnsRequestAction.domainWithSuffix
	// - service.action.networkConnectionAction.blocked
	// - service.action.networkConnectionAction.connectionDirection
	// - service.action.networkConnectionAction.localPortDetails.port
	// - service.action.networkConnectionAction.protocol
	// - service.action.networkConnectionAction.remoteIpDetails.city.cityName
	// - service.action.networkConnectionAction.remoteIpDetails.country.countryName
	// - service.action.networkConnectionAction.remoteIpDetails.ipAddressV4
	// - service.action.networkConnectionAction.remoteIpDetails.ipAddressV6
	// - service.action.networkConnectionAction.remoteIpDetails.organization.asn
	// - service.action.networkConnectionAction.remoteIpDetails.organization.asnOrg
	// - service.action.networkConnectionAction.remotePortDetails.port
	// - service.action.awsApiCallAction.remoteAccountDetails.affiliated
	// - service.action.kubernetesApiCallAction.remoteIpDetails.ipAddressV4
	// - service.action.kubernetesApiCallAction.remoteIpDetails.ipAddressV6
	// - service.action.kubernetesApiCallAction.namespace
	// - service.action.kubernetesApiCallAction.remoteIpDetails.organization.asn
	// - service.action.kubernetesApiCallAction.requestUri
	// - service.action.kubernetesApiCallAction.statusCode
	// - service.action.networkConnectionAction.localIpDetails.ipAddressV4
	// - service.action.networkConnectionAction.localIpDetails.ipAddressV6
	// - service.action.networkConnectionAction.protocol
	// - service.action.awsApiCallAction.serviceName
	// - service.action.awsApiCallAction.remoteAccountDetails.accountId
	// - service.additionalInfo.threatListName
	// - service.resourceRole
	// - resource.eksClusterDetails.name
	// - resource.kubernetesDetails.kubernetesWorkloadDetails.name
	// - resource.kubernetesDetails.kubernetesWorkloadDetails.namespace
	// - resource.kubernetesDetails.kubernetesUserDetails.username
	// - resource.kubernetesDetails.kubernetesWorkloadDetails.containers.image
	// - resource.kubernetesDetails.kubernetesWorkloadDetails.containers.imagePrefix
	// - service.ebsVolumeScanDetails.scanId
	// - service.ebsVolumeScanDetails.scanDetections.threatDetectedByName.threatNames.name
	// - service.ebsVolumeScanDetails.scanDetections.threatDetectedByName.threatNames.severity
	// - service.ebsVolumeScanDetails.scanDetections.threatDetectedByName.threatNames.filePaths.hash
	// - service.malwareScanDetails.threats.name
	// - resource.ecsClusterDetails.name
	// - resource.ecsClusterDetails.taskDetails.containers.image
	// - resource.ecsClusterDetails.taskDetails.definitionArn
	// - resource.containerDetails.image
	// - resource.rdsDbInstanceDetails.dbInstanceIdentifier
	// - resource.rdsDbInstanceDetails.dbClusterIdentifier
	// - resource.rdsDbInstanceDetails.engine
	// - resource.rdsDbUserDetails.user
	// - resource.rdsDbInstanceDetails.tags.key
	// - resource.rdsDbInstanceDetails.tags.value
	// - service.runtimeDetails.process.executableSha256
	// - service.runtimeDetails.process.name
	// - service.runtimeDetails.process.name
	// - resource.lambdaDetails.functionName
	// - resource.lambdaDetails.functionArn
	// - resource.lambdaDetails.tags.key
	// - resource.lambdaDetails.tags.value
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-findingcriteria.html#cfn-guardduty-filter-findingcriteria-criterion
	//
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-findingcriteria.html#cfn-guardduty-filter-findingcriteria-itemtype
	//
	ItemType interface{} `field:"optional" json:"itemType" yaml:"itemType"`
}

