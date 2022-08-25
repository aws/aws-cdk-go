package awsconfig


// The maximum frequency at which the AWS Config rule runs evaluations.
//
// Example:
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   config.NewManagedRule(this, jsii.String("AccessKeysRotated"), &managedRuleProps{
//   	identifier: config.managedRuleIdentifiers_ACCESS_KEYS_ROTATED(),
//   	inputParameters: map[string]interface{}{
//   		"maxAccessKeyAge": jsii.Number(60),
//   	},
//
//   	// default is 24 hours
//   	maximumExecutionFrequency: config.maximumExecutionFrequency_TWELVE_HOURS,
//   })
//
type MaximumExecutionFrequency string

const (
	// 1 hour.
	MaximumExecutionFrequency_ONE_HOUR MaximumExecutionFrequency = "ONE_HOUR"
	// 3 hours.
	MaximumExecutionFrequency_THREE_HOURS MaximumExecutionFrequency = "THREE_HOURS"
	// 6 hours.
	MaximumExecutionFrequency_SIX_HOURS MaximumExecutionFrequency = "SIX_HOURS"
	// 12 hours.
	MaximumExecutionFrequency_TWELVE_HOURS MaximumExecutionFrequency = "TWELVE_HOURS"
	// 24 hours.
	MaximumExecutionFrequency_TWENTY_FOUR_HOURS MaximumExecutionFrequency = "TWENTY_FOUR_HOURS"
)

