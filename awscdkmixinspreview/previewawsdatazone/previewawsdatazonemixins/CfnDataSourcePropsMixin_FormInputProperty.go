package previewawsdatazonemixins


// The details of a metadata form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formInputProperty := &FormInputProperty{
//   	Content: jsii.String("content"),
//   	FormName: jsii.String("formName"),
//   	TypeIdentifier: jsii.String("typeIdentifier"),
//   	TypeRevision: jsii.String("typeRevision"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html
//
type CfnDataSourcePropsMixin_FormInputProperty struct {
	// The content of the metadata form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The name of the metadata form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-formname
	//
	FormName *string `field:"optional" json:"formName" yaml:"formName"`
	// The ID of the metadata form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-typeidentifier
	//
	TypeIdentifier *string `field:"optional" json:"typeIdentifier" yaml:"typeIdentifier"`
	// The revision of the metadata form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-typerevision
	//
	TypeRevision *string `field:"optional" json:"typeRevision" yaml:"typeRevision"`
}

