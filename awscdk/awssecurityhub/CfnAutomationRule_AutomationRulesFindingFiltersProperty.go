package awssecurityhub


// The criteria that determine which findings a rule applies to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automationRulesFindingFiltersProperty := &AutomationRulesFindingFiltersProperty{
//   	AwsAccountId: []interface{}{
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
//   	ProductArn: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProductName: []interface{}{
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
//   	SeverityLabel: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	SourceUrl: []interface{}{
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
//   	WorkflowStatus: []interface{}{
//   		&StringFilterProperty{
//   			Comparison: jsii.String("comparison"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html
//
type CfnAutomationRule_AutomationRulesFindingFiltersProperty struct {
	// The AWS account ID in which a finding was generated.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 100 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-awsaccountid
	//
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The name of the company for the product that generated the finding.
	//
	// For control-based findings, the company is AWS .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-companyname
	//
	CompanyName interface{} `field:"optional" json:"companyName" yaml:"companyName"`
	// The unique identifier of a standard in which a control is enabled.
	//
	// This field consists of the resource portion of the Amazon Resource Name (ARN) returned for a standard in the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API response.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-complianceassociatedstandardsid
	//
	ComplianceAssociatedStandardsId interface{} `field:"optional" json:"complianceAssociatedStandardsId" yaml:"complianceAssociatedStandardsId"`
	// The security control ID for which a finding was generated. Security control IDs are the same across standards.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-compliancesecuritycontrolid
	//
	ComplianceSecurityControlId interface{} `field:"optional" json:"complianceSecurityControlId" yaml:"complianceSecurityControlId"`
	// The result of a security check. This field is only used for findings generated from controls.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-compliancestatus
	//
	ComplianceStatus interface{} `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// The likelihood that a finding accurately identifies the behavior or issue that it was intended to identify.
	//
	// `Confidence` is scored on a 0–100 basis using a ratio scale. A value of `0` means 0 percent confidence, and a value of `100` means 100 percent confidence. For example, a data exfiltration detection based on a statistical deviation of network traffic has low confidence because an actual exfiltration hasn't been verified. For more information, see [Confidence](https://docs.aws.amazon.com/securityhub/latest/userguide/asff-top-level-attributes.html#asff-confidence) in the *Security Hub User Guide* .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-confidence
	//
	Confidence interface{} `field:"optional" json:"confidence" yaml:"confidence"`
	// A timestamp that indicates when this finding record was created.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-createdat
	//
	CreatedAt interface{} `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The level of importance that is assigned to the resources that are associated with a finding.
	//
	// `Criticality` is scored on a 0–100 basis, using a ratio scale that supports only full integers. A score of `0` means that the underlying resources have no criticality, and a score of `100` is reserved for the most critical resources. For more information, see [Criticality](https://docs.aws.amazon.com/securityhub/latest/userguide/asff-top-level-attributes.html#asff-criticality) in the *Security Hub User Guide* .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-criticality
	//
	Criticality interface{} `field:"optional" json:"criticality" yaml:"criticality"`
	// A finding's description.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-description
	//
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// A timestamp that indicates when the potential security issue captured by a finding was first observed by the security findings product.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-firstobservedat
	//
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// The identifier for the solution-specific component that generated a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 100 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-generatorid
	//
	GeneratorId interface{} `field:"optional" json:"generatorId" yaml:"generatorId"`
	// The product-specific identifier for a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-id
	//
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// A timestamp that indicates when the security findings provider most recently observed a change in the resource that is involved in the finding.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-lastobservedat
	//
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// The text of a user-defined note that's added to a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-notetext
	//
	NoteText interface{} `field:"optional" json:"noteText" yaml:"noteText"`
	// The timestamp of when the note was updated.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-noteupdatedat
	//
	NoteUpdatedAt interface{} `field:"optional" json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// The principal that created a note.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-noteupdatedby
	//
	NoteUpdatedBy interface{} `field:"optional" json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// The Amazon Resource Name (ARN) for a third-party product that generated a finding in Security Hub.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-productarn
	//
	ProductArn interface{} `field:"optional" json:"productArn" yaml:"productArn"`
	// Provides the name of the product that generated the finding. For control-based findings, the product name is Security Hub.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-productname
	//
	ProductName interface{} `field:"optional" json:"productName" yaml:"productName"`
	// Provides the current state of a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-recordstate
	//
	RecordState interface{} `field:"optional" json:"recordState" yaml:"recordState"`
	// The product-generated identifier for a related finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-relatedfindingsid
	//
	RelatedFindingsId interface{} `field:"optional" json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// The ARN for the product that generated a related finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-relatedfindingsproductarn
	//
	RelatedFindingsProductArn interface{} `field:"optional" json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// Custom fields and values about the resource that a finding pertains to.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-resourcedetailsother
	//
	ResourceDetailsOther interface{} `field:"optional" json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// The identifier for the given resource type.
	//
	// For AWS resources that are identified by Amazon Resource Names (ARNs), this is the ARN. For AWS resources that lack ARNs, this is the identifier as defined by the AWS service that created the resource. For non- AWS resources, this is a unique identifier that is associated with the resource.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 100 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-resourceid
	//
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The partition in which the resource that the finding pertains to is located.
	//
	// A partition is a group of AWS Regions . Each AWS account is scoped to one partition.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-resourcepartition
	//
	ResourcePartition interface{} `field:"optional" json:"resourcePartition" yaml:"resourcePartition"`
	// The AWS Region where the resource that a finding pertains to is located.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-resourceregion
	//
	ResourceRegion interface{} `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// A list of AWS tags associated with a resource at the time the finding was processed.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// A finding's title.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 100 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-resourcetype
	//
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The severity value of the finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-severitylabel
	//
	SeverityLabel interface{} `field:"optional" json:"severityLabel" yaml:"severityLabel"`
	// Provides a URL that links to a page about the current finding in the finding product.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-sourceurl
	//
	SourceUrl interface{} `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// A finding's title.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 100 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// One or more finding types in the format of namespace/category/classifier that classify a finding.
	//
	// For a list of namespaces, classifiers, and categories, see [Types taxonomy for ASFF](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format-type-taxonomy.html) in the *Security Hub User Guide* .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// A timestamp that indicates when the finding record was most recently updated.
	//
	// For more information about the validation and formatting of timestamp fields in Security Hub , see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps) .
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-updatedat
	//
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// A list of user-defined name and value string pairs added to a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-userdefinedfields
	//
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// Provides the veracity of a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-verificationstate
	//
	VerificationState interface{} `field:"optional" json:"verificationState" yaml:"verificationState"`
	// Provides information about the status of the investigation into a finding.
	//
	// Array Members: Minimum number of 1 item. Maximum number of 20 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfilters.html#cfn-securityhub-automationrule-automationrulesfindingfilters-workflowstatus
	//
	WorkflowStatus interface{} `field:"optional" json:"workflowStatus" yaml:"workflowStatus"`
}

