package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formDataTypeConfigProperty := &formDataTypeConfigProperty{
//   	dataSourceType: jsii.String("dataSourceType"),
//   	dataTypeName: jsii.String("dataTypeName"),
//   }
//
type CfnForm_FormDataTypeConfigProperty struct {
	// `CfnForm.FormDataTypeConfigProperty.DataSourceType`.
	DataSourceType *string `field:"required" json:"dataSourceType" yaml:"dataSourceType"`
	// `CfnForm.FormDataTypeConfigProperty.DataTypeName`.
	DataTypeName *string `field:"required" json:"dataTypeName" yaml:"dataTypeName"`
}

