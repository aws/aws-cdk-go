package awscognito


// The user pool feature plan, or tier.
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-sign-in-feature-plans.html
//
type FeaturePlan string

const (
	// Lite feature plan.
	FeaturePlan_LITE FeaturePlan = "LITE"
	// Essentials feature plan.
	FeaturePlan_ESSENTIALS FeaturePlan = "ESSENTIALS"
	// Plus feature plan.
	FeaturePlan_PLUS FeaturePlan = "PLUS"
)

