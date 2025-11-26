package previewawsguarddutyevents


// Type definition for UnusualBehavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   unusualBehavior := &UnusualBehavior{
//   	IsUnusualUserIdentity: []*string{
//   		jsii.String("isUnusualUserIdentity"),
//   	},
//   	NumberOfPast24HoursApIsBucketProfiling: []*string{
//   		jsii.String("numberOfPast24HoursApIsBucketProfiling"),
//   	},
//   	NumberOfPast24HoursApIsBucketUserIdentityProfiling: []*string{
//   		jsii.String("numberOfPast24HoursApIsBucketUserIdentityProfiling"),
//   	},
//   	NumberOfPast24HoursApIsUserIdentityProfiling: []*string{
//   		jsii.String("numberOfPast24HoursApIsUserIdentityProfiling"),
//   	},
//   	UnusualApIsAccountProfiling: []*string{
//   		jsii.String("unusualApIsAccountProfiling"),
//   	},
//   	UnusualApIsUserIdentityProfiling: []*string{
//   		jsii.String("unusualApIsUserIdentityProfiling"),
//   	},
//   	UnusualAsNsAccountProfiling: []*string{
//   		jsii.String("unusualAsNsAccountProfiling"),
//   	},
//   	UnusualAsNsBucketProfiling: []*string{
//   		jsii.String("unusualAsNsBucketProfiling"),
//   	},
//   	UnusualAsNsUserIdentityProfiling: []*string{
//   		jsii.String("unusualAsNsUserIdentityProfiling"),
//   	},
//   	UnusualBucketsAccountProfiling: []*string{
//   		jsii.String("unusualBucketsAccountProfiling"),
//   	},
//   	UnusualBucketsUserIdentityProfiling: []*string{
//   		jsii.String("unusualBucketsUserIdentityProfiling"),
//   	},
//   	UnusualUserAgentsAccountProfiling: []*string{
//   		jsii.String("unusualUserAgentsAccountProfiling"),
//   	},
//   	UnusualUserAgentsUserIdentityProfiling: []*string{
//   		jsii.String("unusualUserAgentsUserIdentityProfiling"),
//   	},
//   	UnusualUserNamesAccountProfiling: []*string{
//   		jsii.String("unusualUserNamesAccountProfiling"),
//   	},
//   	UnusualUserNamesBucketProfiling: []*string{
//   		jsii.String("unusualUserNamesBucketProfiling"),
//   	},
//   	UnusualUserTypesAccountProfiling: []*string{
//   		jsii.String("unusualUserTypesAccountProfiling"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_UnusualBehavior struct {
	// isUnusualUserIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of isUnusualUserIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IsUnusualUserIdentity *[]*string `field:"optional" json:"isUnusualUserIdentity" yaml:"isUnusualUserIdentity"`
	// numberOfPast24HoursAPIsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfPast24HoursAPIsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfPast24HoursApIsBucketProfiling *[]*string `field:"optional" json:"numberOfPast24HoursApIsBucketProfiling" yaml:"numberOfPast24HoursApIsBucketProfiling"`
	// numberOfPast24HoursAPIsBucketUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfPast24HoursAPIsBucketUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfPast24HoursApIsBucketUserIdentityProfiling *[]*string `field:"optional" json:"numberOfPast24HoursApIsBucketUserIdentityProfiling" yaml:"numberOfPast24HoursApIsBucketUserIdentityProfiling"`
	// numberOfPast24HoursAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfPast24HoursAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfPast24HoursApIsUserIdentityProfiling *[]*string `field:"optional" json:"numberOfPast24HoursApIsUserIdentityProfiling" yaml:"numberOfPast24HoursApIsUserIdentityProfiling"`
	// unusualAPIsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualAPIsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualApIsAccountProfiling *[]*string `field:"optional" json:"unusualApIsAccountProfiling" yaml:"unusualApIsAccountProfiling"`
	// unusualAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualApIsUserIdentityProfiling *[]*string `field:"optional" json:"unusualApIsUserIdentityProfiling" yaml:"unusualApIsUserIdentityProfiling"`
	// unusualASNsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualASNsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualAsNsAccountProfiling *[]*string `field:"optional" json:"unusualAsNsAccountProfiling" yaml:"unusualAsNsAccountProfiling"`
	// unusualASNsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualASNsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualAsNsBucketProfiling *[]*string `field:"optional" json:"unusualAsNsBucketProfiling" yaml:"unusualAsNsBucketProfiling"`
	// unusualASNsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualASNsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualAsNsUserIdentityProfiling *[]*string `field:"optional" json:"unusualAsNsUserIdentityProfiling" yaml:"unusualAsNsUserIdentityProfiling"`
	// unusualBucketsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualBucketsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualBucketsAccountProfiling *[]*string `field:"optional" json:"unusualBucketsAccountProfiling" yaml:"unusualBucketsAccountProfiling"`
	// unusualBucketsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualBucketsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualBucketsUserIdentityProfiling *[]*string `field:"optional" json:"unusualBucketsUserIdentityProfiling" yaml:"unusualBucketsUserIdentityProfiling"`
	// unusualUserAgentsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualUserAgentsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualUserAgentsAccountProfiling *[]*string `field:"optional" json:"unusualUserAgentsAccountProfiling" yaml:"unusualUserAgentsAccountProfiling"`
	// unusualUserAgentsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualUserAgentsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualUserAgentsUserIdentityProfiling *[]*string `field:"optional" json:"unusualUserAgentsUserIdentityProfiling" yaml:"unusualUserAgentsUserIdentityProfiling"`
	// unusualUserNamesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualUserNamesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualUserNamesAccountProfiling *[]*string `field:"optional" json:"unusualUserNamesAccountProfiling" yaml:"unusualUserNamesAccountProfiling"`
	// unusualUserNamesBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualUserNamesBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualUserNamesBucketProfiling *[]*string `field:"optional" json:"unusualUserNamesBucketProfiling" yaml:"unusualUserNamesBucketProfiling"`
	// unusualUserTypesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of unusualUserTypesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UnusualUserTypesAccountProfiling *[]*string `field:"optional" json:"unusualUserTypesAccountProfiling" yaml:"unusualUserTypesAccountProfiling"`
}

