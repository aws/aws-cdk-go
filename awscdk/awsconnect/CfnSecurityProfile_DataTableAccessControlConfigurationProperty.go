package awsconnect


// Defines the access control configuration for data tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataTableAccessControlConfigurationProperty := &DataTableAccessControlConfigurationProperty{
//   	PrimaryAttributeAccessControlConfiguration: &PrimaryAttributeAccessControlConfigurationItemProperty{
//   		PrimaryAttributeValues: []interface{}{
//   			&PrimaryAttributeValueProperty{
//   				AccessType: jsii.String("accessType"),
//   				AttributeName: jsii.String("attributeName"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-datatableaccesscontrolconfiguration.html
//
type CfnSecurityProfile_DataTableAccessControlConfigurationProperty struct {
	// Contains the configuration for record-based access control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-datatableaccesscontrolconfiguration.html#cfn-connect-securityprofile-datatableaccesscontrolconfiguration-primaryattributeaccesscontrolconfiguration
	//
	PrimaryAttributeAccessControlConfiguration interface{} `field:"optional" json:"primaryAttributeAccessControlConfiguration" yaml:"primaryAttributeAccessControlConfiguration"`
}

