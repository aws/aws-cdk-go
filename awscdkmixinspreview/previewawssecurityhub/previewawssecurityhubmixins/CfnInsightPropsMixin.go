package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecurityhub/previewawssecurityhubmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SecurityHub::Insight` resource creates a custom insight in Security Hub .
//
// An insight is a collection of findings that relate to a security issue that requires attention or remediation. For more information, see [Insights in Security Hub](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-insights.html) in the *Security Hub User Guide* .
//
// Tags aren't supported for this resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInsightPropsMixin := awscdkmixinspreview.Mixins.NewCfnInsightPropsMixin(&CfnInsightMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-insight.html
//
type CfnInsightPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInsightMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInsightPropsMixin
type jsiiProxy_CfnInsightPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInsightPropsMixin) Props() *CfnInsightMixinProps {
	var returns *CfnInsightMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInsightPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::Insight`.
func NewCfnInsightPropsMixin(props *CfnInsightMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInsightPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInsightPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInsightPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnInsightPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::Insight`.
func NewCfnInsightPropsMixin_Override(c CfnInsightPropsMixin, props *CfnInsightMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnInsightPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInsightPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInsightPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnInsightPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInsightPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnInsightPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInsightPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInsightPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

