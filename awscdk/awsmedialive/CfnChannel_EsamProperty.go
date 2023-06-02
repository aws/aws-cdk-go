package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   esamProperty := &EsamProperty{
//   	AcquisitionPointId: jsii.String("acquisitionPointId"),
//   	AdAvailOffset: jsii.Number(123),
//   	PasswordParam: jsii.String("passwordParam"),
//   	PoisEndpoint: jsii.String("poisEndpoint"),
//   	Username: jsii.String("username"),
//   	ZoneIdentity: jsii.String("zoneIdentity"),
//   }
//
type CfnChannel_EsamProperty struct {
	// `CfnChannel.EsamProperty.AcquisitionPointId`.
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// `CfnChannel.EsamProperty.AdAvailOffset`.
	AdAvailOffset *float64 `field:"optional" json:"adAvailOffset" yaml:"adAvailOffset"`
	// `CfnChannel.EsamProperty.PasswordParam`.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// `CfnChannel.EsamProperty.PoisEndpoint`.
	PoisEndpoint *string `field:"optional" json:"poisEndpoint" yaml:"poisEndpoint"`
	// `CfnChannel.EsamProperty.Username`.
	Username *string `field:"optional" json:"username" yaml:"username"`
	// `CfnChannel.EsamProperty.ZoneIdentity`.
	ZoneIdentity *string `field:"optional" json:"zoneIdentity" yaml:"zoneIdentity"`
}

