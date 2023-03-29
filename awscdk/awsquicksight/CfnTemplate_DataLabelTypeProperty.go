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
type CfnTemplate_DataLabelTypeProperty struct {
	// `CfnTemplate.DataLabelTypeProperty.DataPathLabelType`.
	DataPathLabelType interface{} `field:"optional" json:"dataPathLabelType" yaml:"dataPathLabelType"`
	// `CfnTemplate.DataLabelTypeProperty.FieldLabelType`.
	FieldLabelType interface{} `field:"optional" json:"fieldLabelType" yaml:"fieldLabelType"`
	// `CfnTemplate.DataLabelTypeProperty.MaximumLabelType`.
	MaximumLabelType interface{} `field:"optional" json:"maximumLabelType" yaml:"maximumLabelType"`
	// `CfnTemplate.DataLabelTypeProperty.MinimumLabelType`.
	MinimumLabelType interface{} `field:"optional" json:"minimumLabelType" yaml:"minimumLabelType"`
	// `CfnTemplate.DataLabelTypeProperty.RangeEndsLabelType`.
	RangeEndsLabelType interface{} `field:"optional" json:"rangeEndsLabelType" yaml:"rangeEndsLabelType"`
}

