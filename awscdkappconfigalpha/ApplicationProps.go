package awscdkappconfigalpha


// Example:
//   appconfig.NewApplication(this, jsii.String("MyApplication"), &ApplicationProps{
//   	Name: jsii.String("App1"),
//   	Description: jsii.String("This is my application created through CDK."),
//   })
//
// Experimental.
type ApplicationProps struct {
	// The description for the application.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the application.
	// Default: - A name is generated.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

