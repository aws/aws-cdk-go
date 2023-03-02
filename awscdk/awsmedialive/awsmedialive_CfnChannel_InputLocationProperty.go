package awsmedialive


// The input location.
//
// The parent of this entity is InputLossBehavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLocationProperty := &inputLocationProperty{
//   	passwordParam: jsii.String("passwordParam"),
//   	uri: jsii.String("uri"),
//   	username: jsii.String("username"),
//   }
//
type CfnChannel_InputLocationProperty struct {
	// The password parameter that holds the password for accessing the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// The URI should be a path to a file that is accessible to the Live system (for example, an http:// URI) depending on the output type.
	//
	// For example, an RTMP destination should have a URI similar to rtmp://fmsserver/live.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The user name to connect to the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

