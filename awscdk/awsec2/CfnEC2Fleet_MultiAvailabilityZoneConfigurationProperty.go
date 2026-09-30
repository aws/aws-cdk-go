package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiAvailabilityZoneConfigurationProperty := &MultiAvailabilityZoneConfigurationProperty{
//   	ConfigurationType: jsii.String("configurationType"),
//   	StandbyAvailabilityZones: []interface{}{
//   		&StandbyAvailabilityZoneProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-multiavailabilityzoneconfiguration.html
//
type CfnEC2Fleet_MultiAvailabilityZoneConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-multiavailabilityzoneconfiguration.html#cfn-ec2-ec2fleet-multiavailabilityzoneconfiguration-configurationtype
	//
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-multiavailabilityzoneconfiguration.html#cfn-ec2-ec2fleet-multiavailabilityzoneconfiguration-standbyavailabilityzones
	//
	StandbyAvailabilityZones interface{} `field:"optional" json:"standbyAvailabilityZones" yaml:"standbyAvailabilityZones"`
}

