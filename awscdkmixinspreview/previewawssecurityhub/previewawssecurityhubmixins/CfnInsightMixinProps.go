package previewawssecurityhubmixins


// Properties for CfnInsightPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInsightMixinProps := &CfnInsightMixinProps{
//   	Filters: &AwsSecurityFindingFiltersProperty{
//   		AwsAccountId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		AwsAccountName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		CompanyName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceAssociatedStandardsId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceSecurityControlId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceSecurityControlParametersName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceSecurityControlParametersValue: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceStatus: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Confidence: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		CreatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		Criticality: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		Description: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingProviderFieldsConfidence: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		FindingProviderFieldsCriticality: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		FindingProviderFieldsRelatedFindingsId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingProviderFieldsRelatedFindingsProductArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingProviderFieldsSeverityLabel: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingProviderFieldsSeverityOriginal: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FindingProviderFieldsTypes: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FirstObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		GeneratorId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Id: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Keyword: []interface{}{
//   			&KeywordFilterProperty{
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		LastObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		MalwareName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MalwarePath: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MalwareState: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MalwareType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NetworkDestinationDomain: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NetworkDestinationIpV4: []interface{}{
//   			&IpFilterProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   		NetworkDestinationIpV6: []interface{}{
//   			&IpFilterProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   		NetworkDestinationPort: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		NetworkDirection: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NetworkProtocol: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NetworkSourceDomain: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NetworkSourceIpV4: []interface{}{
//   			&IpFilterProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   		NetworkSourceIpV6: []interface{}{
//   			&IpFilterProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   		NetworkSourceMac: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NetworkSourcePort: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		NoteText: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NoteUpdatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		NoteUpdatedBy: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProcessLaunchedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		ProcessName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProcessParentPid: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		ProcessPath: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProcessPid: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		ProcessTerminatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		ProductArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProductFields: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProductName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RecommendationText: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RecordState: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Region: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RelatedFindingsId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RelatedFindingsProductArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceApplicationArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceApplicationName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceIamInstanceProfileArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceImageId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceIpV4Addresses: []interface{}{
//   			&IpFilterProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceIpV6Addresses: []interface{}{
//   			&IpFilterProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceKeyName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceLaunchedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceSubnetId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsEc2InstanceVpcId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsIamAccessKeyCreatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		ResourceAwsIamAccessKeyPrincipalName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsIamAccessKeyStatus: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsIamAccessKeyUserName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsIamUserUserName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsS3BucketOwnerId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceAwsS3BucketOwnerName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceContainerImageId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceContainerImageName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceContainerLaunchedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		ResourceContainerName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceDetailsOther: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourcePartition: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceRegion: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceTags: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Sample: []interface{}{
//   			&BooleanFilterProperty{
//   				Value: jsii.Boolean(false),
//   			},
//   		},
//   		SeverityLabel: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		SeverityNormalized: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		SeverityProduct: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		SourceUrl: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThreatIntelIndicatorCategory: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThreatIntelIndicatorLastObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		ThreatIntelIndicatorSource: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThreatIntelIndicatorSourceUrl: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThreatIntelIndicatorType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThreatIntelIndicatorValue: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Title: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Type: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		UpdatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		UserDefinedFields: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VerificationState: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VulnerabilitiesExploitAvailable: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VulnerabilitiesFixAvailable: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		WorkflowState: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		WorkflowStatus: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	GroupByAttribute: jsii.String("groupByAttribute"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-insight.html
//
type CfnInsightMixinProps struct {
	// One or more attributes used to filter the findings included in the insight.
	//
	// The insight only includes findings that match the criteria defined in the filters. You can filter by up to ten finding attributes. For each attribute, you can provide up to 20 filter values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-insight.html#cfn-securityhub-insight-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The grouping attribute for the insight's findings.
	//
	// Indicates how to group the matching findings, and identifies the type of item that the insight applies to. For example, if an insight is grouped by resource identifier, then the insight produces a list of resource identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-insight.html#cfn-securityhub-insight-groupbyattribute
	//
	GroupByAttribute *string `field:"optional" json:"groupByAttribute" yaml:"groupByAttribute"`
	// The name of a Security Hub CSPM insight.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-insight.html#cfn-securityhub-insight-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

