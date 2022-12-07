package awsec2


// Options when constructing UserData for Windows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   windowsUserDataOptions := &windowsUserDataOptions{
//   	persist: jsii.Boolean(false),
//   }
//
type WindowsUserDataOptions struct {
	// Set to true to set this userdata to persist through an instance reboot;
	//
	// allowing
	// it to run on every instance start.
	// By default, UserData is run only once during the first instance launch.
	//
	// For more information, see:
	// https://aws.amazon.com/premiumsupport/knowledge-center/execute-user-data-ec2/
	// https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html#user-data-scripts
	Persist *bool `field:"optional" json:"persist" yaml:"persist"`
}

