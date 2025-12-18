package previewawsconnectmixins


// A data table access control configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnSecurityProfilePropsMixin_DataTableAccessControlConfigurationProperty struct {
	// The configuration's primary attribute access control configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-datatableaccesscontrolconfiguration.html#cfn-connect-securityprofile-datatableaccesscontrolconfiguration-primaryattributeaccesscontrolconfiguration
	//
	PrimaryAttributeAccessControlConfiguration interface{} `field:"optional" json:"primaryAttributeAccessControlConfiguration" yaml:"primaryAttributeAccessControlConfiguration"`
}

