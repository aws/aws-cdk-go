package previewawsquicksightmixins


// The option that determines the data label type.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataLabelTypeProperty := &DataLabelTypeProperty{
//   	DataPathLabelType: &DataPathLabelTypeProperty{
//   		FieldId: jsii.String("fieldId"),
//   		FieldValue: jsii.String("fieldValue"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	FieldLabelType: &FieldLabelTypeProperty{
//   		FieldId: jsii.String("fieldId"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	MaximumLabelType: &MaximumLabelTypeProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	MinimumLabelType: &MinimumLabelTypeProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	RangeEndsLabelType: &RangeEndsLabelTypeProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html
//
type CfnDashboardPropsMixin_DataLabelTypeProperty struct {
	// The option that specifies individual data values for labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-datapathlabeltype
	//
	DataPathLabelType interface{} `field:"optional" json:"dataPathLabelType" yaml:"dataPathLabelType"`
	// Determines the label configuration for the entire field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-fieldlabeltype
	//
	FieldLabelType interface{} `field:"optional" json:"fieldLabelType" yaml:"fieldLabelType"`
	// Determines the label configuration for the maximum value in a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-maximumlabeltype
	//
	MaximumLabelType interface{} `field:"optional" json:"maximumLabelType" yaml:"maximumLabelType"`
	// Determines the label configuration for the minimum value in a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-minimumlabeltype
	//
	MinimumLabelType interface{} `field:"optional" json:"minimumLabelType" yaml:"minimumLabelType"`
	// Determines the label configuration for range end value in a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-rangeendslabeltype
	//
	RangeEndsLabelType interface{} `field:"optional" json:"rangeEndsLabelType" yaml:"rangeEndsLabelType"`
}

