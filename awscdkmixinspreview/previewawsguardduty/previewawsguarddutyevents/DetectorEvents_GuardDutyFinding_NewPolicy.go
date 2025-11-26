package previewawsguarddutyevents


// Type definition for NewPolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   newPolicy := &NewPolicy{
//   	AllowUsersToChangePassword: []*string{
//   		jsii.String("allowUsersToChangePassword"),
//   	},
//   	HardExpiry: []*string{
//   		jsii.String("hardExpiry"),
//   	},
//   	MaxPasswordAge: []*string{
//   		jsii.String("maxPasswordAge"),
//   	},
//   	MinimumPasswordLength: []*string{
//   		jsii.String("minimumPasswordLength"),
//   	},
//   	PasswordReusePrevention: []*string{
//   		jsii.String("passwordReusePrevention"),
//   	},
//   	RequireLowercaseCharacters: []*string{
//   		jsii.String("requireLowercaseCharacters"),
//   	},
//   	RequireNumbers: []*string{
//   		jsii.String("requireNumbers"),
//   	},
//   	RequireSymbols: []*string{
//   		jsii.String("requireSymbols"),
//   	},
//   	RequireUppercaseCharacters: []*string{
//   		jsii.String("requireUppercaseCharacters"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_NewPolicy struct {
	// allowUsersToChangePassword property.
	//
	// Specify an array of string values to match this event if the actual value of allowUsersToChangePassword is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AllowUsersToChangePassword *[]*string `field:"optional" json:"allowUsersToChangePassword" yaml:"allowUsersToChangePassword"`
	// hardExpiry property.
	//
	// Specify an array of string values to match this event if the actual value of hardExpiry is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HardExpiry *[]*string `field:"optional" json:"hardExpiry" yaml:"hardExpiry"`
	// maxPasswordAge property.
	//
	// Specify an array of string values to match this event if the actual value of maxPasswordAge is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxPasswordAge *[]*string `field:"optional" json:"maxPasswordAge" yaml:"maxPasswordAge"`
	// minimumPasswordLength property.
	//
	// Specify an array of string values to match this event if the actual value of minimumPasswordLength is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MinimumPasswordLength *[]*string `field:"optional" json:"minimumPasswordLength" yaml:"minimumPasswordLength"`
	// passwordReusePrevention property.
	//
	// Specify an array of string values to match this event if the actual value of passwordReusePrevention is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PasswordReusePrevention *[]*string `field:"optional" json:"passwordReusePrevention" yaml:"passwordReusePrevention"`
	// requireLowercaseCharacters property.
	//
	// Specify an array of string values to match this event if the actual value of requireLowercaseCharacters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequireLowercaseCharacters *[]*string `field:"optional" json:"requireLowercaseCharacters" yaml:"requireLowercaseCharacters"`
	// requireNumbers property.
	//
	// Specify an array of string values to match this event if the actual value of requireNumbers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequireNumbers *[]*string `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// requireSymbols property.
	//
	// Specify an array of string values to match this event if the actual value of requireSymbols is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequireSymbols *[]*string `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// requireUppercaseCharacters property.
	//
	// Specify an array of string values to match this event if the actual value of requireUppercaseCharacters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequireUppercaseCharacters *[]*string `field:"optional" json:"requireUppercaseCharacters" yaml:"requireUppercaseCharacters"`
}

