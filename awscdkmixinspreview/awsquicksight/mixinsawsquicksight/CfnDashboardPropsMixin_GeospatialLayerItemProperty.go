package mixinsawsquicksight


// The properties for a single geospatial layer.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html
//
type CfnDashboardPropsMixin_GeospatialLayerItemProperty struct {
	// A list of custom actions for a layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The data source for the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The join definition properties for a layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-joindefinition
	//
	JoinDefinition interface{} `field:"optional" json:"joinDefinition" yaml:"joinDefinition"`
	// The label that is displayed for the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The definition properties for a layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-layerdefinition
	//
	LayerDefinition interface{} `field:"optional" json:"layerDefinition" yaml:"layerDefinition"`
	// The ID of the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-layerid
	//
	LayerId *string `field:"optional" json:"layerId" yaml:"layerId"`
	// The layer type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-layertype
	//
	LayerType *string `field:"optional" json:"layerType" yaml:"layerType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The state of visibility for the layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayeritem.html#cfn-quicksight-dashboard-geospatiallayeritem-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

