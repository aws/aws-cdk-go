package awsmedialive


// Settings that apply only if the input is a pull type of input.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSourceRequestProperty := &inputSourceRequestProperty{
//   	passwordParam: jsii.String("passwordParam"),
//   	url: jsii.String("url"),
//   	username: jsii.String("username"),
//   }
//
type CfnInput_InputSourceRequestProperty struct {
	// The password parameter that holds the password for accessing the upstream system.
	//
	// The password parameter applies only if the upstream system requires credentials.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// For a pull input, the URL where MediaLive pulls the source content from.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name to connect to the upstream system.
	//
	// The user name applies only if the upstream system requires credentials.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

