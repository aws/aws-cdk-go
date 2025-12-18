package awsconnect


// Contains granular access control configuration for security profiles, including data table access permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   granularAccessControlConfigurationProperty := &GranularAccessControlConfigurationProperty{
//   	DataTableAccessControlConfiguration: &DataTableAccessControlConfigurationProperty{
//   		PrimaryAttributeAccessControlConfiguration: &PrimaryAttributeAccessControlConfigurationItemProperty{
//   			PrimaryAttributeValues: []interface{}{
//   				&PrimaryAttributeValueProperty{
//   					AccessType: jsii.String("accessType"),
//   					AttributeName: jsii.String("attributeName"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-granularaccesscontrolconfiguration.html
//
type CfnSecurityProfile_GranularAccessControlConfigurationProperty struct {
	// The access control configuration for data tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-securityprofile-granularaccesscontrolconfiguration.html#cfn-connect-securityprofile-granularaccesscontrolconfiguration-datatableaccesscontrolconfiguration
	//
	DataTableAccessControlConfiguration interface{} `field:"optional" json:"dataTableAccessControlConfiguration" yaml:"dataTableAccessControlConfiguration"`
}

