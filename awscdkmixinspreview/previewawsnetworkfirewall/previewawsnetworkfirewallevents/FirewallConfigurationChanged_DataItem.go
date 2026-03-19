package previewawsnetworkfirewallevents


// Type definition for DataItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataItem := &DataItem{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	ConfigurationResourceArn: []*string{
//   		jsii.String("configurationResourceArn"),
//   	},
//   	CurrentConfigurationSyncStatus: []*string{
//   		jsii.String("currentConfigurationSyncStatus"),
//   	},
//   	CurrentConfigurationUpdateToken: []*string{
//   		jsii.String("currentConfigurationUpdateToken"),
//   	},
//   	PreviousConfigurationSyncStatus: []*string{
//   		jsii.String("previousConfigurationSyncStatus"),
//   	},
//   	PreviousConfigurationUpdateToken: []*string{
//   		jsii.String("previousConfigurationUpdateToken"),
//   	},
//   }
//
// Experimental.
type FirewallConfigurationChanged_DataItem struct {
	// Availability Zone property.
	//
	// Specify an array of string values to match this event if the actual value of Availability Zone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Configuration Resource ARN property.
	//
	// Specify an array of string values to match this event if the actual value of Configuration Resource ARN is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConfigurationResourceArn *[]*string `field:"optional" json:"configurationResourceArn" yaml:"configurationResourceArn"`
	// Current Configuration Sync Status property.
	//
	// Specify an array of string values to match this event if the actual value of Current Configuration Sync Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentConfigurationSyncStatus *[]*string `field:"optional" json:"currentConfigurationSyncStatus" yaml:"currentConfigurationSyncStatus"`
	// Current Configuration Update Token property.
	//
	// Specify an array of string values to match this event if the actual value of Current Configuration Update Token is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentConfigurationUpdateToken *[]*string `field:"optional" json:"currentConfigurationUpdateToken" yaml:"currentConfigurationUpdateToken"`
	// Previous Configuration Sync Status property.
	//
	// Specify an array of string values to match this event if the actual value of Previous Configuration Sync Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousConfigurationSyncStatus *[]*string `field:"optional" json:"previousConfigurationSyncStatus" yaml:"previousConfigurationSyncStatus"`
	// Previous Configuration Update Token property.
	//
	// Specify an array of string values to match this event if the actual value of Previous Configuration Update Token is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousConfigurationUpdateToken *[]*string `field:"optional" json:"previousConfigurationUpdateToken" yaml:"previousConfigurationUpdateToken"`
}

