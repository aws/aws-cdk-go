package awsquicksight


// The configuration of a tree map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html
//
type CfnTemplate_TreeMapConfigurationProperty struct {
	// The label options (label text, label visibility) for the colors displayed in a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-colorlabeloptions
	//
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// The color options (gradient color, point of divergence) of a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-colorscale
	//
	ColorScale interface{} `field:"optional" json:"colorScale" yaml:"colorScale"`
	// The options that determine if visual data labels are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The label options (label text, label visibility) of the groups that are displayed in a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-grouplabeloptions
	//
	GroupLabelOptions interface{} `field:"optional" json:"groupLabelOptions" yaml:"groupLabelOptions"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The label options (label text, label visibility) of the sizes that are displayed in a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-sizelabeloptions
	//
	SizeLabelOptions interface{} `field:"optional" json:"sizeLabelOptions" yaml:"sizeLabelOptions"`
	// The sort configuration of a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapconfiguration.html#cfn-quicksight-template-treemapconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
}

