package awsdatazone


// The details of a metadata form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formInputProperty := &FormInputProperty{
//   	FormName: jsii.String("formName"),
//
//   	// the properties below are optional
//   	Content: jsii.String("content"),
//   	TypeIdentifier: jsii.String("typeIdentifier"),
//   	TypeRevision: jsii.String("typeRevision"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html
//
type CfnDataSource_FormInputProperty struct {
	// The name of the metadata form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-formname
	//
	FormName *string `field:"required" json:"formName" yaml:"formName"`
	// The content of the metadata form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The ID of the metadata form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-typeidentifier
	//
	TypeIdentifier *string `field:"optional" json:"typeIdentifier" yaml:"typeIdentifier"`
	// The revision of the metadata form type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-forminput.html#cfn-datazone-datasource-forminput-typerevision
	//
	TypeRevision *string `field:"optional" json:"typeRevision" yaml:"typeRevision"`
}

