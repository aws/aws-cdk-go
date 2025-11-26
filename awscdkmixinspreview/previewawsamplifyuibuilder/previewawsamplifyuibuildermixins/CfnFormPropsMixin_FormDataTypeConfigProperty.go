package previewawsamplifyuibuildermixins


// The `FormDataTypeConfig` property specifies the data type configuration for the data source associated with a form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formDataTypeConfigProperty := &FormDataTypeConfigProperty{
//   	DataSourceType: jsii.String("dataSourceType"),
//   	DataTypeName: jsii.String("dataTypeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formdatatypeconfig.html
//
type CfnFormPropsMixin_FormDataTypeConfigProperty struct {
	// The data source type, either an Amplify DataStore model or a custom data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formdatatypeconfig.html#cfn-amplifyuibuilder-form-formdatatypeconfig-datasourcetype
	//
	DataSourceType *string `field:"optional" json:"dataSourceType" yaml:"dataSourceType"`
	// The unique name of the data type you are using as the data source for the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-formdatatypeconfig.html#cfn-amplifyuibuilder-form-formdatatypeconfig-datatypename
	//
	DataTypeName *string `field:"optional" json:"dataTypeName" yaml:"dataTypeName"`
}

