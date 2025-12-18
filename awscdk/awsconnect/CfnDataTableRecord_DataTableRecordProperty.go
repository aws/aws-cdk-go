package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataTableRecordProperty := &DataTableRecordProperty{
//   	Values: []interface{}{
//   		&ValueProperty{
//   			AttributeId: jsii.String("attributeId"),
//   			AttributeValue: jsii.String("attributeValue"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	PrimaryValues: []interface{}{
//   		&ValueProperty{
//   			AttributeId: jsii.String("attributeId"),
//   			AttributeValue: jsii.String("attributeValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-datatablerecord.html
//
type CfnDataTableRecord_DataTableRecordProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-datatablerecord.html#cfn-connect-datatablerecord-datatablerecord-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-datatablerecord.html#cfn-connect-datatablerecord-datatablerecord-primaryvalues
	//
	PrimaryValues interface{} `field:"optional" json:"primaryValues" yaml:"primaryValues"`
}

