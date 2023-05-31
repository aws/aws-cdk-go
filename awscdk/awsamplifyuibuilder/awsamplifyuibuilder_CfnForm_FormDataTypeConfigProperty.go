package awsamplifyuibuilder


// The `FormDataTypeConfig` property specifies the data type configuration for the data source associated with a form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formDataTypeConfigProperty := &FormDataTypeConfigProperty{
//   	DataSourceType: jsii.String("dataSourceType"),
//   	DataTypeName: jsii.String("dataTypeName"),
//   }
//
type CfnForm_FormDataTypeConfigProperty struct {
	// The data source type, either an Amplify DataStore model or a custom data type.
	DataSourceType *string `field:"required" json:"dataSourceType" yaml:"dataSourceType"`
	// The unique name of the data type you are using as the data source for the form.
	DataTypeName *string `field:"required" json:"dataTypeName" yaml:"dataTypeName"`
}

