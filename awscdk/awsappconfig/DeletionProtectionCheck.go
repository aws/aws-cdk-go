package awsappconfig


// The deletion protection check options.
//
// Example:
//   var application application
//   var alarm alarm
//   var compositeAlarm compositeAlarm
//   var bucket bucket
//
//
//   // Environment deletion protection check
//   // Environment deletion protection check
//   appconfig.NewEnvironment(this, jsii.String("MyEnvironment"), &EnvironmentProps{
//   	Application: Application,
//   	DeletionProtectionCheck: appconfig.DeletionProtectionCheck_APPLY,
//   })
//
//   // configuration profile with deletion protection check
//   // configuration profile with deletion protection check
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfigFromFile"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromFile(jsii.String("config.json")),
//   	DeletionProtectionCheck: appconfig.DeletionProtectionCheck_BYPASS,
//   })
//
//   appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
//   	Application: Application,
//   	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
//   	DeletionProtectionCheck: appconfig.DeletionProtectionCheck_ACCOUNT_DEFAULT,
//   })
//
type DeletionProtectionCheck string

const (
	// The default setting, which uses account-level deletion protection.
	//
	// To configure account-level deletion protection, use the UpdateAccountSettings API.
	DeletionProtectionCheck_ACCOUNT_DEFAULT DeletionProtectionCheck = "ACCOUNT_DEFAULT"
	// Instructs the deletion protection check to run, even if deletion protection is disabled at the account level.
	//
	// APPLY also forces the deletion protection check to run against resources created in the past hour,
	// which are normally excluded from deletion protection checks.
	DeletionProtectionCheck_APPLY DeletionProtectionCheck = "APPLY"
	// Instructs AWS AppConfig to bypass the deletion protection check and delete an environment or a configuration profile even if deletion protection would have otherwise prevented it.
	DeletionProtectionCheck_BYPASS DeletionProtectionCheck = "BYPASS"
)

