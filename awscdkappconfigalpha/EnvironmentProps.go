package awscdkappconfigalpha


// Properties for the Environment construct.
//
// Example:
//   var application application
//   var alarm alarm
//   var compositeAlarm compositeAlarm
//
//
//   appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
//   	Application: Application,
//   	Monitors: []monitor{
//   		appconfig.*monitor_FromCloudWatchAlarm(alarm),
//   		appconfig.*monitor_*FromCloudWatchAlarm(compositeAlarm),
//   	},
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

