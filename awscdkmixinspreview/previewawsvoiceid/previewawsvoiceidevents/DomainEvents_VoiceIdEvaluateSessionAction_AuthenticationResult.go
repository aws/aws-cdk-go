package previewawsvoiceidevents


// Type definition for AuthenticationResult.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticationResult := &AuthenticationResult{
//   	AudioAggregationEndedAt: []*string{
//   		jsii.String("audioAggregationEndedAt"),
//   	},
//   	AudioAggregationStartedAt: []*string{
//   		jsii.String("audioAggregationStartedAt"),
//   	},
//   	AuthenticationResultId: []*string{
//   		jsii.String("authenticationResultId"),
//   	},
//   	Configuration: &ConfigurationAuthentication{
//   		AcceptanceThreshold: []*string{
//   			jsii.String("acceptanceThreshold"),
//   		},
//   	},
//   	Decision: []*string{
//   		jsii.String("decision"),
//   	},
//   	Score: []*string{
//   		jsii.String("score"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdEvaluateSessionAction_AuthenticationResult struct {
	// audioAggregationEndedAt property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationEndedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationEndedAt *[]*string `field:"optional" json:"audioAggregationEndedAt" yaml:"audioAggregationEndedAt"`
	// audioAggregationStartedAt property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationStartedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationStartedAt *[]*string `field:"optional" json:"audioAggregationStartedAt" yaml:"audioAggregationStartedAt"`
	// authenticationResultId property.
	//
	// Specify an array of string values to match this event if the actual value of authenticationResultId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthenticationResultId *[]*string `field:"optional" json:"authenticationResultId" yaml:"authenticationResultId"`
	// configuration property.
	//
	// Specify an array of string values to match this event if the actual value of configuration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Configuration *DomainEvents_VoiceIdEvaluateSessionAction_ConfigurationAuthentication `field:"optional" json:"configuration" yaml:"configuration"`
	// decision property.
	//
	// Specify an array of string values to match this event if the actual value of decision is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Decision *[]*string `field:"optional" json:"decision" yaml:"decision"`
	// score property.
	//
	// Specify an array of string values to match this event if the actual value of score is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Score *[]*string `field:"optional" json:"score" yaml:"score"`
}

