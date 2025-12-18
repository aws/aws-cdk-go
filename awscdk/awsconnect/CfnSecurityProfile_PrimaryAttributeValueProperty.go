package awsconnect


// A primary attribute value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryAttributeValueProperty := &PrimaryAttributeValueProperty{
//   	AccessType: jsii.String("accessType"),
//   	AttributeName: jsii.String("attributeName"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html
//
type CfnSecurityProfile_PrimaryAttributeValueProperty struct {
	// The value's access type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html#cfn-connect-securityprofile-primaryattributevalue-accesstype
	//
	AccessType *string `field:"required" json:"accessType" yaml:"accessType"`
	// The value's attribute name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html#cfn-connect-securityprofile-primaryattributevalue-attributename
	//
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The value's values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html#cfn-connect-securityprofile-primaryattributevalue-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

