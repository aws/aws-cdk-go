package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueProperty := &ValueProperty{
//   	AttributeId: jsii.String("attributeId"),
//   	AttributeValue: jsii.String("attributeValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-value.html
//
type CfnDataTableRecord_ValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-value.html#cfn-connect-datatablerecord-value-attributeid
	//
	AttributeId *string `field:"optional" json:"attributeId" yaml:"attributeId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatablerecord-value.html#cfn-connect-datatablerecord-value-attributevalue
	//
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
}

