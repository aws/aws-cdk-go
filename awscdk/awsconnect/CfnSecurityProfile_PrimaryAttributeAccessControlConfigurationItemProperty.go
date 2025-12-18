package awsconnect


// A primary attribute access control configuration item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryAttributeAccessControlConfigurationItemProperty := &PrimaryAttributeAccessControlConfigurationItemProperty{
//   	PrimaryAttributeValues: []interface{}{
//   		&PrimaryAttributeValueProperty{
//   			AccessType: jsii.String("accessType"),
//   			AttributeName: jsii.String("attributeName"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributeaccesscontrolconfigurationitem.html
//
type CfnSecurityProfile_PrimaryAttributeAccessControlConfigurationItemProperty struct {
	// The item's primary attribute values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-primaryattributeaccesscontrolconfigurationitem.html#cfn-connect-securityprofile-primaryattributeaccesscontrolconfigurationitem-primaryattributevalues
	//
	PrimaryAttributeValues interface{} `field:"required" json:"primaryAttributeValues" yaml:"primaryAttributeValues"`
}

