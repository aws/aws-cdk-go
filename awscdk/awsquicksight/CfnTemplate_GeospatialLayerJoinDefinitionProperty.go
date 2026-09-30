package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerjoindefinition.html
//
type CfnTemplate_GeospatialLayerJoinDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerjoindefinition.html#cfn-quicksight-template-geospatiallayerjoindefinition-colorfield
	//
	ColorField interface{} `field:"optional" json:"colorField" yaml:"colorField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerjoindefinition.html#cfn-quicksight-template-geospatiallayerjoindefinition-datasetkeyfield
	//
	DatasetKeyField interface{} `field:"optional" json:"datasetKeyField" yaml:"datasetKeyField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerjoindefinition.html#cfn-quicksight-template-geospatiallayerjoindefinition-shapekeyfield
	//
	ShapeKeyField *string `field:"optional" json:"shapeKeyField" yaml:"shapeKeyField"`
}

