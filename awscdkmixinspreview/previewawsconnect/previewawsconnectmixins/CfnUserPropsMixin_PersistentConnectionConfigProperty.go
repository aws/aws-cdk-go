package previewawsconnectmixins


// Persistent Connection configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   persistentConnectionConfigProperty := &PersistentConnectionConfigProperty{
//   	Channel: jsii.String("channel"),
//   	PersistentConnection: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-persistentconnectionconfig.html
//
type CfnUserPropsMixin_PersistentConnectionConfigProperty struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-persistentconnectionconfig.html#cfn-connect-user-persistentconnectionconfig-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The Persistent Connection setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-persistentconnectionconfig.html#cfn-connect-user-persistentconnectionconfig-persistentconnection
	//
	PersistentConnection interface{} `field:"optional" json:"persistentConnection" yaml:"persistentConnection"`
}

