package previewawsvoiceidevents


// Type definition for Configuration_Fraud.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configurationFraud := &ConfigurationFraud{
//   	RiskThreshold: []*string{
//   		jsii.String("riskThreshold"),
//   	},
//   	WatchlistId: []*string{
//   		jsii.String("watchlistId"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_ConfigurationFraud struct {
	// riskThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of riskThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RiskThreshold *[]*string `field:"optional" json:"riskThreshold" yaml:"riskThreshold"`
	// watchlistId property.
	//
	// Specify an array of string values to match this event if the actual value of watchlistId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WatchlistId *[]*string `field:"optional" json:"watchlistId" yaml:"watchlistId"`
}

