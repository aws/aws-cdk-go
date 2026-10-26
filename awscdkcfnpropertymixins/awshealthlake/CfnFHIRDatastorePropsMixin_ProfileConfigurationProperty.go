package awshealthlake


// The profile configuration for the Data Store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   profileConfigurationProperty := &ProfileConfigurationProperty{
//   	DefaultProfiles: []*string{
//   		jsii.String("defaultProfiles"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-profileconfiguration.html
//
type CfnFHIRDatastorePropsMixin_ProfileConfigurationProperty struct {
	// The FHIR profiles the Data Store validates resources against by default, for example us-core-3.1.1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-profileconfiguration.html#cfn-healthlake-fhirdatastore-profileconfiguration-defaultprofiles
	//
	DefaultProfiles *[]*string `field:"optional" json:"defaultProfiles" yaml:"defaultProfiles"`
}

