package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html
//
type CfnAnalysis_GeospatialLayerItemProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-layerid
	//
	LayerId *string `field:"required" json:"layerId" yaml:"layerId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-joindefinition
	//
	JoinDefinition interface{} `field:"optional" json:"joinDefinition" yaml:"joinDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-layerdefinition
	//
	LayerDefinition interface{} `field:"optional" json:"layerDefinition" yaml:"layerDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-layertype
	//
	LayerType *string `field:"optional" json:"layerType" yaml:"layerType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayeritem.html#cfn-quicksight-analysis-geospatiallayeritem-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

