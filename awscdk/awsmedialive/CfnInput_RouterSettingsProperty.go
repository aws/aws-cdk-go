package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routerSettingsProperty := &RouterSettingsProperty{
//   	Destinations: []interface{}{
//   		&RouterDestinationSettingsProperty{
//   			AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   		},
//   	},
//   	EncryptionType: jsii.String("encryptionType"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html
//
type CfnInput_RouterSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html#cfn-medialive-input-routersettings-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html#cfn-medialive-input-routersettings-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html#cfn-medialive-input-routersettings-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

