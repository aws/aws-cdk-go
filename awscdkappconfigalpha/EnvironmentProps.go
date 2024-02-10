package awscdkappconfigalpha


// Properties for the Environment construct.
//
// Example:
//   app := appconfig.NewApplication(this, jsii.String("MyApp"))
//   env := appconfig.NewEnvironment(this, jsii.String("MyEnv"), &EnvironmentProps{
//   	Application: app,
//   })
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfig"), &HostedConfigurationProps{
//   	Application: app,
//   	DeployTo: []iEnvironment{
//   		env,
//   	},
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   })
//
// Experimental.
type EnvironmentProps struct {
	// The description of the environment.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the environment.
	// Default: - A name is generated.
	//
	// Experimental.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The monitors for the environment.
	// Default: - No monitors.
	//
	// Experimental.
	Monitors *[]Monitor `field:"optional" json:"monitors" yaml:"monitors"`
	// The application to be associated with the environment.
	// Experimental.
	Application IApplication `field:"required" json:"application" yaml:"application"`
}

