package awsses


// Confidence threshold used by SES Auto Validation to decide whether an outbound recipient address should be delivered to.
//
// Example:
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	AutoValidationThreshold: ses.AutoValidationThreshold_HIGH,
//   })
//
// See: https://docs.aws.amazon.com/ses/latest/dg/email-validation-auto.html
//
type AutoValidationThreshold string

const (
	// Medium confidence threshold.
	//
	// Allow addresses with medium or higher delivery likelihood.
	AutoValidationThreshold_MEDIUM AutoValidationThreshold = "MEDIUM"
	// High confidence threshold.
	//
	// Only allow addresses with high delivery likelihood.
	AutoValidationThreshold_HIGH AutoValidationThreshold = "HIGH"
	// Amazon SES manages the threshold automatically based on sending patterns and reputation.
	AutoValidationThreshold_MANAGED AutoValidationThreshold = "MANAGED"
)

