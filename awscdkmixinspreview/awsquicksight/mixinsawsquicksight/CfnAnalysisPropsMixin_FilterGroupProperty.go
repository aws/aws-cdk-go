package mixinsawsquicksight


// A grouping of individual filters. Filter groups are applied to the same group of visuals.
//
// For more information, see [Adding filter conditions (group filters) with AND and OR operators](https://docs.aws.amazon.com/quicksight/latest/user/add-a-compound-filter.html) in the *Amazon Quick Suite User Guide* .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html
//
type CfnAnalysisPropsMixin_FilterGroupProperty struct {
	// The filter new feature which can apply filter group to all data sets. Choose one of the following options:.
	//
	// - `ALL_DATASETS`
	// - `SINGLE_DATASET`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-crossdataset
	//
	CrossDataset *string `field:"optional" json:"crossDataset" yaml:"crossDataset"`
	// The value that uniquely identifies a `FilterGroup` within a dashboard, template, or analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-filtergroupid
	//
	FilterGroupId *string `field:"optional" json:"filterGroupId" yaml:"filterGroupId"`
	// The list of filters that are present in a `FilterGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// The configuration that specifies what scope to apply to a `FilterGroup` .
	//
	// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-scopeconfiguration
	//
	ScopeConfiguration interface{} `field:"optional" json:"scopeConfiguration" yaml:"scopeConfiguration"`
	// The status of the `FilterGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filtergroup.html#cfn-quicksight-analysis-filtergroup-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

