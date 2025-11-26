package previewawssecurityhubmixins


// A collection of filters that are applied to all active findings aggregated by Security Hub .
//
// You can filter by up to ten finding attributes. For each attribute, you can provide up to 20 filter values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsSecurityFindingFiltersProperty := &AwsSecurityFindingFiltersProperty{
//   	AwsAccountId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	AwsAccountName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	CompanyName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ComplianceAssociatedStandardsId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ComplianceSecurityControlId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ComplianceSecurityControlParametersName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ComplianceSecurityControlParametersValue: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ComplianceStatus: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Confidence: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	CreatedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	Criticality: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	Description: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FindingProviderFieldsConfidence: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	FindingProviderFieldsCriticality: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	FindingProviderFieldsRelatedFindingsId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FindingProviderFieldsRelatedFindingsProductArn: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FindingProviderFieldsSeverityLabel: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FindingProviderFieldsSeverityOriginal: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FindingProviderFieldsTypes: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FirstObservedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	GeneratorId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Id: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Keyword: []interface{}{
//   		&KeywordFilterProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	LastObservedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	MalwareName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MalwarePath: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MalwareState: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MalwareType: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NetworkDestinationDomain: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NetworkDestinationIpV4: []interface{}{
//   		&IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	NetworkDestinationIpV6: []interface{}{
//   		&IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	NetworkDestinationPort: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	NetworkDirection: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NetworkProtocol: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NetworkSourceDomain: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NetworkSourceIpV4: []interface{}{
//   		&IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	NetworkSourceIpV6: []interface{}{
//   		&IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	NetworkSourceMac: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NetworkSourcePort: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	NoteText: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NoteUpdatedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	NoteUpdatedBy: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProcessLaunchedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	ProcessName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProcessParentPid: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	ProcessPath: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProcessPid: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	ProcessTerminatedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	ProductArn: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProductFields: []interface{}{
//   		&MapFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProductName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RecommendationText: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RecordState: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Region: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RelatedFindingsId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RelatedFindingsProductArn: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceApplicationArn: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceApplicationName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceIamInstanceProfileArn: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceImageId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceIpV4Addresses: []interface{}{
//   		&IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceIpV6Addresses: []interface{}{
//   		&IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceKeyName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceLaunchedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceSubnetId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceType: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsEc2InstanceVpcId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsIamAccessKeyCreatedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	ResourceAwsIamAccessKeyPrincipalName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsIamAccessKeyStatus: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsIamAccessKeyUserName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsIamUserUserName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsS3BucketOwnerId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceAwsS3BucketOwnerName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceContainerImageId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceContainerImageName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceContainerLaunchedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	ResourceContainerName: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceDetailsOther: []interface{}{
//   		&MapFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceId: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourcePartition: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceRegion: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceTags: []interface{}{
//   		&MapFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ResourceType: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Sample: []interface{}{
//   		&BooleanFilterProperty{
//   			Value: jsii.Boolean(false),
//   		},
//   	},
//   	SeverityLabel: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	SeverityNormalized: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	SeverityProduct: []interface{}{
//   		&NumberFilterProperty{
//   			Eq: jsii.Number(123),
//   			Gte: jsii.Number(123),
//   			Lte: jsii.Number(123),
//   		},
//   	},
//   	SourceUrl: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThreatIntelIndicatorCategory: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThreatIntelIndicatorLastObservedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	ThreatIntelIndicatorSource: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThreatIntelIndicatorSourceUrl: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThreatIntelIndicatorType: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThreatIntelIndicatorValue: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UpdatedAt: []interface{}{
//   		&DateFilterProperty{
//   			DateRange: &DateRangeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//   	UserDefinedFields: []interface{}{
//   		&MapFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VerificationState: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VulnerabilitiesExploitAvailable: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VulnerabilitiesFixAvailable: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowState: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowStatus: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html
//
type CfnInsightPropsMixin_AwsSecurityFindingFiltersProperty struct {
	// The AWS account ID in which a finding is generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-awsaccountid
	//
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The name of the AWS account in which a finding is generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-awsaccountname
	//
	AwsAccountName interface{} `field:"optional" json:"awsAccountName" yaml:"awsAccountName"`
	// The name of the findings provider (company) that owns the solution (product) that generates findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-companyname
	//
	CompanyName interface{} `field:"optional" json:"companyName" yaml:"companyName"`
	// The unique identifier of a standard in which a control is enabled.
	//
	// This field consists of the resource portion of the Amazon Resource Name (ARN) returned for a standard in the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-complianceassociatedstandardsid
	//
	ComplianceAssociatedStandardsId interface{} `field:"optional" json:"complianceAssociatedStandardsId" yaml:"complianceAssociatedStandardsId"`
	// The unique identifier of a control across standards.
	//
	// Values for this field typically consist of an AWS service and a number, such as APIGateway.5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-compliancesecuritycontrolid
	//
	ComplianceSecurityControlId interface{} `field:"optional" json:"complianceSecurityControlId" yaml:"complianceSecurityControlId"`
	// The name of a security control parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-compliancesecuritycontrolparametersname
	//
	ComplianceSecurityControlParametersName interface{} `field:"optional" json:"complianceSecurityControlParametersName" yaml:"complianceSecurityControlParametersName"`
	// The current value of a security control parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-compliancesecuritycontrolparametersvalue
	//
	ComplianceSecurityControlParametersValue interface{} `field:"optional" json:"complianceSecurityControlParametersValue" yaml:"complianceSecurityControlParametersValue"`
	// Exclusive to findings that are generated as the result of a check run against a specific rule in a supported standard, such as CIS AWS Foundations.
	//
	// Contains security standard-related finding details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-compliancestatus
	//
	ComplianceStatus interface{} `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// A finding's confidence.
	//
	// Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify.
	//
	// Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent confidence and 100 means 100 percent confidence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-confidence
	//
	Confidence interface{} `field:"optional" json:"confidence" yaml:"confidence"`
	// A timestamp that indicates when the security findings provider created the potential security issue that a finding reflects.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-createdat
	//
	CreatedAt interface{} `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The level of importance assigned to the resources associated with the finding.
	//
	// A score of 0 means that the underlying resources have no criticality, and a score of 100 is reserved for the most critical resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-criticality
	//
	Criticality interface{} `field:"optional" json:"criticality" yaml:"criticality"`
	// A finding's description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-description
	//
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// The finding provider value for the finding confidence.
	//
	// Confidence is defined as the likelihood that a finding accurately identifies the behavior or issue that it was intended to identify.
	//
	// Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent confidence and 100 means 100 percent confidence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldsconfidence
	//
	FindingProviderFieldsConfidence interface{} `field:"optional" json:"findingProviderFieldsConfidence" yaml:"findingProviderFieldsConfidence"`
	// The finding provider value for the level of importance assigned to the resources associated with the findings.
	//
	// A score of 0 means that the underlying resources have no criticality, and a score of 100 is reserved for the most critical resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldscriticality
	//
	FindingProviderFieldsCriticality interface{} `field:"optional" json:"findingProviderFieldsCriticality" yaml:"findingProviderFieldsCriticality"`
	// The finding identifier of a related finding that is identified by the finding provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldsrelatedfindingsid
	//
	FindingProviderFieldsRelatedFindingsId interface{} `field:"optional" json:"findingProviderFieldsRelatedFindingsId" yaml:"findingProviderFieldsRelatedFindingsId"`
	// The ARN of the solution that generated a related finding that is identified by the finding provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldsrelatedfindingsproductarn
	//
	FindingProviderFieldsRelatedFindingsProductArn interface{} `field:"optional" json:"findingProviderFieldsRelatedFindingsProductArn" yaml:"findingProviderFieldsRelatedFindingsProductArn"`
	// The finding provider value for the severity label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldsseveritylabel
	//
	FindingProviderFieldsSeverityLabel interface{} `field:"optional" json:"findingProviderFieldsSeverityLabel" yaml:"findingProviderFieldsSeverityLabel"`
	// The finding provider's original value for the severity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldsseverityoriginal
	//
	FindingProviderFieldsSeverityOriginal interface{} `field:"optional" json:"findingProviderFieldsSeverityOriginal" yaml:"findingProviderFieldsSeverityOriginal"`
	// One or more finding types that the finding provider assigned to the finding.
	//
	// Uses the format of `namespace/category/classifier` that classify a finding.
	//
	// Valid namespace values are: Software and Configuration Checks | TTPs | Effects | Unusual Behaviors | Sensitive Data Identifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-findingproviderfieldstypes
	//
	FindingProviderFieldsTypes interface{} `field:"optional" json:"findingProviderFieldsTypes" yaml:"findingProviderFieldsTypes"`
	// A timestamp that indicates when the security findings provider first observed the potential security issue that a finding captured.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-firstobservedat
	//
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// The identifier for the solution-specific component (a discrete unit of logic) that generated a finding.
	//
	// In various security findings providers' solutions, this generator can be called a rule, a check, a detector, a plugin, etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-generatorid
	//
	GeneratorId interface{} `field:"optional" json:"generatorId" yaml:"generatorId"`
	// The security findings provider-specific identifier for a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-id
	//
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// This field is deprecated.
	//
	// A keyword for a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-keyword
	//
	Keyword interface{} `field:"optional" json:"keyword" yaml:"keyword"`
	// A timestamp that indicates when the security findings provider most recently observed a change in the resource that is involved in the finding.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-lastobservedat
	//
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// The name of the malware that was observed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-malwarename
	//
	MalwareName interface{} `field:"optional" json:"malwareName" yaml:"malwareName"`
	// The filesystem path of the malware that was observed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-malwarepath
	//
	MalwarePath interface{} `field:"optional" json:"malwarePath" yaml:"malwarePath"`
	// The state of the malware that was observed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-malwarestate
	//
	MalwareState interface{} `field:"optional" json:"malwareState" yaml:"malwareState"`
	// The type of the malware that was observed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-malwaretype
	//
	MalwareType interface{} `field:"optional" json:"malwareType" yaml:"malwareType"`
	// The destination domain of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networkdestinationdomain
	//
	NetworkDestinationDomain interface{} `field:"optional" json:"networkDestinationDomain" yaml:"networkDestinationDomain"`
	// The destination IPv4 address of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networkdestinationipv4
	//
	NetworkDestinationIpV4 interface{} `field:"optional" json:"networkDestinationIpV4" yaml:"networkDestinationIpV4"`
	// The destination IPv6 address of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networkdestinationipv6
	//
	NetworkDestinationIpV6 interface{} `field:"optional" json:"networkDestinationIpV6" yaml:"networkDestinationIpV6"`
	// The destination port of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networkdestinationport
	//
	NetworkDestinationPort interface{} `field:"optional" json:"networkDestinationPort" yaml:"networkDestinationPort"`
	// Indicates the direction of network traffic associated with a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networkdirection
	//
	NetworkDirection interface{} `field:"optional" json:"networkDirection" yaml:"networkDirection"`
	// The protocol of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networkprotocol
	//
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// The source domain of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networksourcedomain
	//
	NetworkSourceDomain interface{} `field:"optional" json:"networkSourceDomain" yaml:"networkSourceDomain"`
	// The source IPv4 address of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networksourceipv4
	//
	NetworkSourceIpV4 interface{} `field:"optional" json:"networkSourceIpV4" yaml:"networkSourceIpV4"`
	// The source IPv6 address of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networksourceipv6
	//
	NetworkSourceIpV6 interface{} `field:"optional" json:"networkSourceIpV6" yaml:"networkSourceIpV6"`
	// The source media access control (MAC) address of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networksourcemac
	//
	NetworkSourceMac interface{} `field:"optional" json:"networkSourceMac" yaml:"networkSourceMac"`
	// The source port of network-related information about a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-networksourceport
	//
	NetworkSourcePort interface{} `field:"optional" json:"networkSourcePort" yaml:"networkSourcePort"`
	// The text of a note.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-notetext
	//
	NoteText interface{} `field:"optional" json:"noteText" yaml:"noteText"`
	// The timestamp of when the note was updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-noteupdatedat
	//
	NoteUpdatedAt interface{} `field:"optional" json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// The principal that created a note.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-noteupdatedby
	//
	NoteUpdatedBy interface{} `field:"optional" json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// A timestamp that identifies when the process was launched.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-processlaunchedat
	//
	ProcessLaunchedAt interface{} `field:"optional" json:"processLaunchedAt" yaml:"processLaunchedAt"`
	// The name of the process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-processname
	//
	ProcessName interface{} `field:"optional" json:"processName" yaml:"processName"`
	// The parent process ID.
	//
	// This field accepts positive integers between `O` and `2147483647` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-processparentpid
	//
	ProcessParentPid interface{} `field:"optional" json:"processParentPid" yaml:"processParentPid"`
	// The path to the process executable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-processpath
	//
	ProcessPath interface{} `field:"optional" json:"processPath" yaml:"processPath"`
	// The process ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-processpid
	//
	ProcessPid interface{} `field:"optional" json:"processPid" yaml:"processPid"`
	// A timestamp that identifies when the process was terminated.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-processterminatedat
	//
	ProcessTerminatedAt interface{} `field:"optional" json:"processTerminatedAt" yaml:"processTerminatedAt"`
	// The ARN generated by Security Hub that uniquely identifies a third-party company (security findings provider) after this provider's product (solution that generates findings) is registered with Security Hub.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-productarn
	//
	ProductArn interface{} `field:"optional" json:"productArn" yaml:"productArn"`
	// A data type where security findings providers can include additional solution-specific details that aren't part of the defined `AwsSecurityFinding` format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-productfields
	//
	ProductFields interface{} `field:"optional" json:"productFields" yaml:"productFields"`
	// The name of the solution (product) that generates findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-productname
	//
	ProductName interface{} `field:"optional" json:"productName" yaml:"productName"`
	// The recommendation of what to do about the issue described in a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-recommendationtext
	//
	RecommendationText interface{} `field:"optional" json:"recommendationText" yaml:"recommendationText"`
	// The updated record state for the finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-recordstate
	//
	RecordState interface{} `field:"optional" json:"recordState" yaml:"recordState"`
	// The Region from which the finding was generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-region
	//
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// The solution-generated identifier for a related finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-relatedfindingsid
	//
	RelatedFindingsId interface{} `field:"optional" json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// The ARN of the solution that generated a related finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-relatedfindingsproductarn
	//
	RelatedFindingsProductArn interface{} `field:"optional" json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// The ARN of the application that is related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceapplicationarn
	//
	ResourceApplicationArn interface{} `field:"optional" json:"resourceApplicationArn" yaml:"resourceApplicationArn"`
	// The name of the application that is related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceapplicationname
	//
	ResourceApplicationName interface{} `field:"optional" json:"resourceApplicationName" yaml:"resourceApplicationName"`
	// The IAM profile ARN of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instanceiaminstanceprofilearn
	//
	ResourceAwsEc2InstanceIamInstanceProfileArn interface{} `field:"optional" json:"resourceAwsEc2InstanceIamInstanceProfileArn" yaml:"resourceAwsEc2InstanceIamInstanceProfileArn"`
	// The Amazon Machine Image (AMI) ID of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instanceimageid
	//
	ResourceAwsEc2InstanceImageId interface{} `field:"optional" json:"resourceAwsEc2InstanceImageId" yaml:"resourceAwsEc2InstanceImageId"`
	// The IPv4 addresses associated with the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instanceipv4addresses
	//
	ResourceAwsEc2InstanceIpV4Addresses interface{} `field:"optional" json:"resourceAwsEc2InstanceIpV4Addresses" yaml:"resourceAwsEc2InstanceIpV4Addresses"`
	// The IPv6 addresses associated with the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instanceipv6addresses
	//
	ResourceAwsEc2InstanceIpV6Addresses interface{} `field:"optional" json:"resourceAwsEc2InstanceIpV6Addresses" yaml:"resourceAwsEc2InstanceIpV6Addresses"`
	// The key name associated with the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instancekeyname
	//
	ResourceAwsEc2InstanceKeyName interface{} `field:"optional" json:"resourceAwsEc2InstanceKeyName" yaml:"resourceAwsEc2InstanceKeyName"`
	// The date and time the instance was launched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instancelaunchedat
	//
	ResourceAwsEc2InstanceLaunchedAt interface{} `field:"optional" json:"resourceAwsEc2InstanceLaunchedAt" yaml:"resourceAwsEc2InstanceLaunchedAt"`
	// The identifier of the subnet that the instance was launched in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instancesubnetid
	//
	ResourceAwsEc2InstanceSubnetId interface{} `field:"optional" json:"resourceAwsEc2InstanceSubnetId" yaml:"resourceAwsEc2InstanceSubnetId"`
	// The instance type of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instancetype
	//
	ResourceAwsEc2InstanceType interface{} `field:"optional" json:"resourceAwsEc2InstanceType" yaml:"resourceAwsEc2InstanceType"`
	// The identifier of the VPC that the instance was launched in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsec2instancevpcid
	//
	ResourceAwsEc2InstanceVpcId interface{} `field:"optional" json:"resourceAwsEc2InstanceVpcId" yaml:"resourceAwsEc2InstanceVpcId"`
	// The creation date/time of the IAM access key related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsiamaccesskeycreatedat
	//
	ResourceAwsIamAccessKeyCreatedAt interface{} `field:"optional" json:"resourceAwsIamAccessKeyCreatedAt" yaml:"resourceAwsIamAccessKeyCreatedAt"`
	// The name of the principal that is associated with an IAM access key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsiamaccesskeyprincipalname
	//
	ResourceAwsIamAccessKeyPrincipalName interface{} `field:"optional" json:"resourceAwsIamAccessKeyPrincipalName" yaml:"resourceAwsIamAccessKeyPrincipalName"`
	// The status of the IAM access key related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsiamaccesskeystatus
	//
	ResourceAwsIamAccessKeyStatus interface{} `field:"optional" json:"resourceAwsIamAccessKeyStatus" yaml:"resourceAwsIamAccessKeyStatus"`
	// This field is deprecated.
	//
	// The username associated with the IAM access key related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsiamaccesskeyusername
	//
	ResourceAwsIamAccessKeyUserName interface{} `field:"optional" json:"resourceAwsIamAccessKeyUserName" yaml:"resourceAwsIamAccessKeyUserName"`
	// The name of an IAM user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawsiamuserusername
	//
	ResourceAwsIamUserUserName interface{} `field:"optional" json:"resourceAwsIamUserUserName" yaml:"resourceAwsIamUserUserName"`
	// The canonical user ID of the owner of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawss3bucketownerid
	//
	ResourceAwsS3BucketOwnerId interface{} `field:"optional" json:"resourceAwsS3BucketOwnerId" yaml:"resourceAwsS3BucketOwnerId"`
	// The display name of the owner of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceawss3bucketownername
	//
	ResourceAwsS3BucketOwnerName interface{} `field:"optional" json:"resourceAwsS3BucketOwnerName" yaml:"resourceAwsS3BucketOwnerName"`
	// The identifier of the image related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcecontainerimageid
	//
	ResourceContainerImageId interface{} `field:"optional" json:"resourceContainerImageId" yaml:"resourceContainerImageId"`
	// The name of the image related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcecontainerimagename
	//
	ResourceContainerImageName interface{} `field:"optional" json:"resourceContainerImageName" yaml:"resourceContainerImageName"`
	// A timestamp that identifies when the container was started.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcecontainerlaunchedat
	//
	ResourceContainerLaunchedAt interface{} `field:"optional" json:"resourceContainerLaunchedAt" yaml:"resourceContainerLaunchedAt"`
	// The name of the container related to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcecontainername
	//
	ResourceContainerName interface{} `field:"optional" json:"resourceContainerName" yaml:"resourceContainerName"`
	// The details of a resource that doesn't have a specific subfield for the resource type defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcedetailsother
	//
	ResourceDetailsOther interface{} `field:"optional" json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// The canonical identifier for the given resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceid
	//
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The canonical AWS partition name that the Region is assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcepartition
	//
	ResourcePartition interface{} `field:"optional" json:"resourcePartition" yaml:"resourcePartition"`
	// The canonical AWS external Region name where this resource is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourceregion
	//
	ResourceRegion interface{} `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// A list of AWS tags associated with a resource at the time the finding was processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Specifies the type of the resource that details are provided for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-resourcetype
	//
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Indicates whether or not sample findings are included in the filter results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-sample
	//
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// The label of a finding's severity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-severitylabel
	//
	SeverityLabel interface{} `field:"optional" json:"severityLabel" yaml:"severityLabel"`
	// Deprecated. The normalized severity of a finding. Instead of providing `Normalized` , provide `Label` .
	//
	// The value of `Normalized` can be an integer between `0` and `100` .
	//
	// If you provide `Label` and don't provide `Normalized` , then `Normalized` is set automatically as follows.
	//
	// - `INFORMATIONAL` - 0
	// - `LOW` - 1
	// - `MEDIUM` - 40
	// - `HIGH` - 70
	// - `CRITICAL` - 90.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-severitynormalized
	//
	SeverityNormalized interface{} `field:"optional" json:"severityNormalized" yaml:"severityNormalized"`
	// Deprecated. This attribute isn't included in findings. Instead of providing `Product` , provide `Original` .
	//
	// The native severity as defined by the AWS service or integrated partner product that generated the finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-severityproduct
	//
	SeverityProduct interface{} `field:"optional" json:"severityProduct" yaml:"severityProduct"`
	// A URL that links to a page about the current finding in the security findings provider's solution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-sourceurl
	//
	SourceUrl interface{} `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// The category of a threat intelligence indicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-threatintelindicatorcategory
	//
	ThreatIntelIndicatorCategory interface{} `field:"optional" json:"threatIntelIndicatorCategory" yaml:"threatIntelIndicatorCategory"`
	// A timestamp that identifies the last observation of a threat intelligence indicator.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-threatintelindicatorlastobservedat
	//
	ThreatIntelIndicatorLastObservedAt interface{} `field:"optional" json:"threatIntelIndicatorLastObservedAt" yaml:"threatIntelIndicatorLastObservedAt"`
	// The source of the threat intelligence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-threatintelindicatorsource
	//
	ThreatIntelIndicatorSource interface{} `field:"optional" json:"threatIntelIndicatorSource" yaml:"threatIntelIndicatorSource"`
	// The URL for more details from the source of the threat intelligence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-threatintelindicatorsourceurl
	//
	ThreatIntelIndicatorSourceUrl interface{} `field:"optional" json:"threatIntelIndicatorSourceUrl" yaml:"threatIntelIndicatorSourceUrl"`
	// The type of a threat intelligence indicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-threatintelindicatortype
	//
	ThreatIntelIndicatorType interface{} `field:"optional" json:"threatIntelIndicatorType" yaml:"threatIntelIndicatorType"`
	// The value of a threat intelligence indicator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-threatintelindicatorvalue
	//
	ThreatIntelIndicatorValue interface{} `field:"optional" json:"threatIntelIndicatorValue" yaml:"threatIntelIndicatorValue"`
	// A finding's title.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// A finding type in the format of `namespace/category/classifier` that classifies a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// A timestamp that indicates when the security findings provider last updated the finding record.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-updatedat
	//
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// A list of name/value string pairs associated with the finding.
	//
	// These are custom, user-defined fields added to a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-userdefinedfields
	//
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// The veracity of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-verificationstate
	//
	VerificationState interface{} `field:"optional" json:"verificationState" yaml:"verificationState"`
	// Indicates whether a software vulnerability in your environment has a known exploit.
	//
	// You can filter findings by this field only if you use Security Hub and Amazon Inspector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-vulnerabilitiesexploitavailable
	//
	VulnerabilitiesExploitAvailable interface{} `field:"optional" json:"vulnerabilitiesExploitAvailable" yaml:"vulnerabilitiesExploitAvailable"`
	// Indicates whether a vulnerability is fixed in a newer version of the affected software packages.
	//
	// You can filter findings by this field only if you use Security Hub and Amazon Inspector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-vulnerabilitiesfixavailable
	//
	VulnerabilitiesFixAvailable interface{} `field:"optional" json:"vulnerabilitiesFixAvailable" yaml:"vulnerabilitiesFixAvailable"`
	// The workflow state of a finding.
	//
	// Note that this field is deprecated. To search for a finding based on its workflow status, use `WorkflowStatus` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-workflowstate
	//
	WorkflowState interface{} `field:"optional" json:"workflowState" yaml:"workflowState"`
	// The status of the investigation into a finding. Allowed values are the following.
	//
	// - `NEW` - The initial state of a finding, before it is reviewed.
	//
	// Security Hub also resets the workflow status from `NOTIFIED` or `RESOLVED` to `NEW` in the following cases:
	//
	// - `RecordState` changes from `ARCHIVED` to `ACTIVE` .
	// - `Compliance.Status` changes from `PASSED` to either `WARNING` , `FAILED` , or `NOT_AVAILABLE` .
	// - `NOTIFIED` - Indicates that the resource owner has been notified about the security issue. Used when the initial reviewer is not the resource owner, and needs intervention from the resource owner.
	//
	// If one of the following occurs, the workflow status is changed automatically from `NOTIFIED` to `NEW` :
	//
	// - `RecordState` changes from `ARCHIVED` to `ACTIVE` .
	// - `Compliance.Status` changes from `PASSED` to `FAILED` , `WARNING` , or `NOT_AVAILABLE` .
	// - `SUPPRESSED` - Indicates that you reviewed the finding and don't believe that any action is needed.
	//
	// The workflow status of a `SUPPRESSED` finding does not change if `RecordState` changes from `ARCHIVED` to `ACTIVE` .
	// - `RESOLVED` - The finding was reviewed and remediated and is now considered resolved.
	//
	// The finding remains `RESOLVED` unless one of the following occurs:
	//
	// - `RecordState` changes from `ARCHIVED` to `ACTIVE` .
	// - `Compliance.Status` changes from `PASSED` to `FAILED` , `WARNING` , or `NOT_AVAILABLE` .
	//
	// In those cases, the workflow status is automatically reset to `NEW` .
	//
	// For findings from controls, if `Compliance.Status` is `PASSED` , then Security Hub automatically sets the workflow status to `RESOLVED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-awssecurityfindingfilters.html#cfn-securityhub-insight-awssecurityfindingfilters-workflowstatus
	//
	WorkflowStatus interface{} `field:"optional" json:"workflowStatus" yaml:"workflowStatus"`
}

