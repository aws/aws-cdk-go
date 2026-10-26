package awsmedialive


// Information about the MediaConnect Router association for this input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   routerSettingsProperty := &RouterSettingsProperty{
//   	Destinations: []interface{}{
//   		&RouterDestinationSettingsProperty{
//   			AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   			RouterOutputArn: jsii.String("routerOutputArn"),
//   		},
//   	},
//   	EncryptionType: jsii.String("encryptionType"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html
//
type CfnInputPropsMixin_RouterSettingsProperty struct {
	// MediaConnect Router destinations associated with this input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html#cfn-medialive-input-routersettings-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Encryption configuration for the MediaConnect Router association.
	//
	// AUTOMATIC uses a service-managed secret; SECRETS_MANAGER requires SecretArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html#cfn-medialive-input-routersettings-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The ARN of the Secrets Manager secret used to encrypt this input.
	//
	// Required when EncryptionType is SECRETS_MANAGER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-routersettings.html#cfn-medialive-input-routersettings-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

