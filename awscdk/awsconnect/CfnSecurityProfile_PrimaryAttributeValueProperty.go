package awsconnect


// An object defining the access control for a specific attribute and its values.
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
	// Specifies the type of access granted.
	//
	// Currently, only "ALLOW" is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html#cfn-connect-securityprofile-primaryattributevalue-accesstype
	//
	AccessType *string `field:"required" json:"accessType" yaml:"accessType"`
	// The name of the primary attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html#cfn-connect-securityprofile-primaryattributevalue-attributename
	//
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// An array of allowed primary values for the specified primary attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributevalue.html#cfn-connect-securityprofile-primaryattributevalue-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

