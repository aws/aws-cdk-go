package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDashboard_DataLabelTypeProperty struct {
	// `CfnDashboard.DataLabelTypeProperty.DataPathLabelType`.
	DataPathLabelType interface{} `field:"optional" json:"dataPathLabelType" yaml:"dataPathLabelType"`
	// `CfnDashboard.DataLabelTypeProperty.FieldLabelType`.
	FieldLabelType interface{} `field:"optional" json:"fieldLabelType" yaml:"fieldLabelType"`
	// `CfnDashboard.DataLabelTypeProperty.MaximumLabelType`.
	MaximumLabelType interface{} `field:"optional" json:"maximumLabelType" yaml:"maximumLabelType"`
	// `CfnDashboard.DataLabelTypeProperty.MinimumLabelType`.
	MinimumLabelType interface{} `field:"optional" json:"minimumLabelType" yaml:"minimumLabelType"`
	// `CfnDashboard.DataLabelTypeProperty.RangeEndsLabelType`.
	RangeEndsLabelType interface{} `field:"optional" json:"rangeEndsLabelType" yaml:"rangeEndsLabelType"`
}

