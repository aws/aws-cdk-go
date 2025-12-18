package awsconnect


// Properties for defining a `CfnDataTableRecord`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataTableRecordProps := &CfnDataTableRecordProps{
//   	DataTableArn: jsii.String("dataTableArn"),
//   	DataTableRecord: &DataTableRecordProperty{
//   		Values: []interface{}{
//   			&ValueProperty{
//   				AttributeId: jsii.String("attributeId"),
//   				AttributeValue: jsii.String("attributeValue"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		PrimaryValues: []interface{}{
//   			&ValueProperty{
//   				AttributeId: jsii.String("attributeId"),
//   				AttributeValue: jsii.String("attributeValue"),
//   			},
//   		},
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatablerecord.html
//
type CfnDataTableRecordProps struct {
	// The Amazon Resource Name (ARN) for the data table.
	//
	// Does not include version aliases.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatablerecord.html#cfn-connect-datatablerecord-datatablearn
	//
	DataTableArn *string `field:"optional" json:"dataTableArn" yaml:"dataTableArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatablerecord.html#cfn-connect-datatablerecord-datatablerecord
	//
	DataTableRecord interface{} `field:"optional" json:"dataTableRecord" yaml:"dataTableRecord"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatablerecord.html#cfn-connect-datatablerecord-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
}

