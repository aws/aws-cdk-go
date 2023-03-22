package awsmedialive


// The configuration information for this output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputDestinationSettingsProperty := &outputDestinationSettingsProperty{
//   	passwordParam: jsii.String("passwordParam"),
//   	streamName: jsii.String("streamName"),
//   	url: jsii.String("url"),
//   	username: jsii.String("username"),
//   }
//
type CfnChannel_OutputDestinationSettingsProperty struct {
	// The password parameter that holds the password for accessing the downstream system.
	//
	// This password parameter applies only if the downstream system requires credentials.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// The stream name for the content.
	//
	// This applies only to RTMP outputs.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// The URL for the destination.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name to connect to the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

