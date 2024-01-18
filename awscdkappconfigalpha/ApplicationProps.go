package awscdkappconfigalpha


// Properties for the Application construct.
//
// Example:
//   appconfig.NewApplication(this, jsii.String("MyApplication"), &ApplicationProps{
//   	ApplicationName: jsii.String("App1"),
//   	Description: jsii.String("This is my application created through CDK."),
//   })
//
// Experimental.
type ApplicationProps struct {
	// The name of the application.
	// Default: - A name is generated.
	//
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The description for the application.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

