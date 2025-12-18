package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataTableRecordProperty := &DataTableRecordProperty{
//   	PrimaryValues: []interface{}{
//   		&ValueProperty{
//   			AttributeId: jsii.String("attributeId"),
//   			AttributeValue: jsii.String("attributeValue"),
//   		},
//   	},
//   	Values: []interface{}{
//   		&ValueProperty{
//   			AttributeId: jsii.String("attributeId"),
//   			AttributeValue: jsii.String("attributeValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-datatablerecord.html
//
type CfnDataTableRecordPropsMixin_DataTableRecordProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-datatablerecord.html#cfn-connect-datatablerecord-datatablerecord-primaryvalues
	//
	PrimaryValues interface{} `field:"optional" json:"primaryValues" yaml:"primaryValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-datatablerecord.html#cfn-connect-datatablerecord-datatablerecord-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

