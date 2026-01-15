package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSet`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html
//
type CfnDataSetProps struct {
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Groupings of columns that work together in certain Amazon Quick Sight features.
	//
	// Currently, only geospatial hierarchy is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-columngroups
	//
	ColumnGroups interface{} `field:"optional" json:"columnGroups" yaml:"columnGroups"`
	// A set of one or more definitions of a `ColumnLevelPermissionRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-columnlevelpermissionrules
	//
	ColumnLevelPermissionRules interface{} `field:"optional" json:"columnLevelPermissionRules" yaml:"columnLevelPermissionRules"`
	// The data preparation configuration associated with this dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-dataprepconfiguration
	//
	DataPrepConfiguration interface{} `field:"optional" json:"dataPrepConfiguration" yaml:"dataPrepConfiguration"`
	// An ID for the dataset that you want to create.
	//
	// This ID is unique per AWS Region for each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetid
	//
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
	// The parameters that are declared in a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetparameters
	//
	DatasetParameters interface{} `field:"optional" json:"datasetParameters" yaml:"datasetParameters"`
	// The refresh properties of a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetrefreshproperties
	//
	DataSetRefreshProperties interface{} `field:"optional" json:"dataSetRefreshProperties" yaml:"dataSetRefreshProperties"`
	// The usage configuration to apply to child datasets that reference this dataset as a source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-datasetusageconfiguration
	//
	DataSetUsageConfiguration interface{} `field:"optional" json:"dataSetUsageConfiguration" yaml:"dataSetUsageConfiguration"`
	// The folder that contains fields and nested subfolders for your dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-fieldfolders
	//
	FieldFolders interface{} `field:"optional" json:"fieldFolders" yaml:"fieldFolders"`
	// <p>When you create the dataset, Quick Suite adds the dataset to these folders.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-folderarns
	//
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// Indicates whether you want to import the data into SPICE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-importmode
	//
	ImportMode *string `field:"optional" json:"importMode" yaml:"importMode"`
	// The wait policy to use when creating or updating a Dataset.
	//
	// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-ingestionwaitpolicy
	//
	IngestionWaitPolicy interface{} `field:"optional" json:"ingestionWaitPolicy" yaml:"ingestionWaitPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-logicaltablemap
	//
	// Deprecated: this property has been deprecated.
	LogicalTableMap interface{} `field:"optional" json:"logicalTableMap" yaml:"logicalTableMap"`
	// The display name for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The performance optimization configuration of a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-performanceconfiguration
	//
	PerformanceConfiguration interface{} `field:"optional" json:"performanceConfiguration" yaml:"performanceConfiguration"`
	// A list of resource permissions on the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Declares the physical tables that are available in the underlying data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-physicaltablemap
	//
	PhysicalTableMap interface{} `field:"optional" json:"physicalTableMap" yaml:"physicalTableMap"`
	// <p>Information about a dataset that contains permissions for row-level security (RLS).
	//
	// The permissions dataset maps fields to users or groups. For more information, see
	//             <a href="https://docs.aws.amazon.com/quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.html">Using Row-Level Security (RLS) to Restrict Access to a Dataset</a> in the <i>Amazon Quick Suite User
	//                 Guide</i>.</p>
	//          <p>The option to deny permissions by setting <code>PermissionPolicy</code> to <code>DENY_ACCESS</code> is
	//             not supported for new RLS datasets.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-rowlevelpermissiondataset
	//
	// Deprecated: this property has been deprecated.
	RowLevelPermissionDataSet interface{} `field:"optional" json:"rowLevelPermissionDataSet" yaml:"rowLevelPermissionDataSet"`
	// <p>The configuration of tags on a dataset to set row-level security.
	//
	// </p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-rowlevelpermissiontagconfiguration
	//
	// Deprecated: this property has been deprecated.
	RowLevelPermissionTagConfiguration interface{} `field:"optional" json:"rowLevelPermissionTagConfiguration" yaml:"rowLevelPermissionTagConfiguration"`
	// The semantic model configuration associated with this dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-semanticmodelconfiguration
	//
	SemanticModelConfiguration interface{} `field:"optional" json:"semanticModelConfiguration" yaml:"semanticModelConfiguration"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The usage of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html#cfn-quicksight-dataset-useas
	//
	UseAs *string `field:"optional" json:"useAs" yaml:"useAs"`
}

