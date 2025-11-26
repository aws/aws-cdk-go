package previewawsguarddutyevents


// Type definition for ProfiledBehavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   profiledBehavior := &ProfiledBehavior{
//   	FrequentProfiledApIsAccountProfiling: []*string{
//   		jsii.String("frequentProfiledApIsAccountProfiling"),
//   	},
//   	FrequentProfiledApIsUserIdentityProfiling: []*string{
//   		jsii.String("frequentProfiledApIsUserIdentityProfiling"),
//   	},
//   	FrequentProfiledAsNsAccountProfiling: []*string{
//   		jsii.String("frequentProfiledAsNsAccountProfiling"),
//   	},
//   	FrequentProfiledAsNsBucketProfiling: []*string{
//   		jsii.String("frequentProfiledAsNsBucketProfiling"),
//   	},
//   	FrequentProfiledAsNsUserIdentityProfiling: []*string{
//   		jsii.String("frequentProfiledAsNsUserIdentityProfiling"),
//   	},
//   	FrequentProfiledBucketsAccountProfiling: []*string{
//   		jsii.String("frequentProfiledBucketsAccountProfiling"),
//   	},
//   	FrequentProfiledBucketsUserIdentityProfiling: []*string{
//   		jsii.String("frequentProfiledBucketsUserIdentityProfiling"),
//   	},
//   	FrequentProfiledUserAgentsAccountProfiling: []*string{
//   		jsii.String("frequentProfiledUserAgentsAccountProfiling"),
//   	},
//   	FrequentProfiledUserAgentsUserIdentityProfiling: []*string{
//   		jsii.String("frequentProfiledUserAgentsUserIdentityProfiling"),
//   	},
//   	FrequentProfiledUserNamesAccountProfiling: []*string{
//   		jsii.String("frequentProfiledUserNamesAccountProfiling"),
//   	},
//   	FrequentProfiledUserNamesBucketProfiling: []*string{
//   		jsii.String("frequentProfiledUserNamesBucketProfiling"),
//   	},
//   	FrequentProfiledUserTypesAccountProfiling: []*string{
//   		jsii.String("frequentProfiledUserTypesAccountProfiling"),
//   	},
//   	InfrequentProfiledApIsAccountProfiling: []*string{
//   		jsii.String("infrequentProfiledApIsAccountProfiling"),
//   	},
//   	InfrequentProfiledApIsUserIdentityProfiling: []*string{
//   		jsii.String("infrequentProfiledApIsUserIdentityProfiling"),
//   	},
//   	InfrequentProfiledAsNsAccountProfiling: []*string{
//   		jsii.String("infrequentProfiledAsNsAccountProfiling"),
//   	},
//   	InfrequentProfiledAsNsBucketProfiling: []*string{
//   		jsii.String("infrequentProfiledAsNsBucketProfiling"),
//   	},
//   	InfrequentProfiledAsNsUserIdentityProfiling: []*string{
//   		jsii.String("infrequentProfiledAsNsUserIdentityProfiling"),
//   	},
//   	InfrequentProfiledBucketsAccountProfiling: []*string{
//   		jsii.String("infrequentProfiledBucketsAccountProfiling"),
//   	},
//   	InfrequentProfiledBucketsUserIdentityProfiling: []*string{
//   		jsii.String("infrequentProfiledBucketsUserIdentityProfiling"),
//   	},
//   	InfrequentProfiledUserAgentsAccountProfiling: []*string{
//   		jsii.String("infrequentProfiledUserAgentsAccountProfiling"),
//   	},
//   	InfrequentProfiledUserAgentsUserIdentityProfiling: []*string{
//   		jsii.String("infrequentProfiledUserAgentsUserIdentityProfiling"),
//   	},
//   	InfrequentProfiledUserNamesAccountProfiling: []*string{
//   		jsii.String("infrequentProfiledUserNamesAccountProfiling"),
//   	},
//   	InfrequentProfiledUserNamesBucketProfiling: []*string{
//   		jsii.String("infrequentProfiledUserNamesBucketProfiling"),
//   	},
//   	InfrequentProfiledUserTypesAccountProfiling: []*string{
//   		jsii.String("infrequentProfiledUserTypesAccountProfiling"),
//   	},
//   	NumberOfHistoricalDailyAvgApIsBucketProfiling: []*string{
//   		jsii.String("numberOfHistoricalDailyAvgApIsBucketProfiling"),
//   	},
//   	NumberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling: []*string{
//   		jsii.String("numberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling"),
//   	},
//   	NumberOfHistoricalDailyAvgApIsUserIdentityProfiling: []*string{
//   		jsii.String("numberOfHistoricalDailyAvgApIsUserIdentityProfiling"),
//   	},
//   	NumberOfHistoricalDailyMaxApIsBucketProfiling: []*string{
//   		jsii.String("numberOfHistoricalDailyMaxApIsBucketProfiling"),
//   	},
//   	NumberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling: []*string{
//   		jsii.String("numberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling"),
//   	},
//   	NumberOfHistoricalDailyMaxApIsUserIdentityProfiling: []*string{
//   		jsii.String("numberOfHistoricalDailyMaxApIsUserIdentityProfiling"),
//   	},
//   	RareProfiledApIsAccountProfiling: []*string{
//   		jsii.String("rareProfiledApIsAccountProfiling"),
//   	},
//   	RareProfiledApIsUserIdentityProfiling: []*string{
//   		jsii.String("rareProfiledApIsUserIdentityProfiling"),
//   	},
//   	RareProfiledAsNsAccountProfiling: []*string{
//   		jsii.String("rareProfiledAsNsAccountProfiling"),
//   	},
//   	RareProfiledAsNsBucketProfiling: []*string{
//   		jsii.String("rareProfiledAsNsBucketProfiling"),
//   	},
//   	RareProfiledAsNsUserIdentityProfiling: []*string{
//   		jsii.String("rareProfiledAsNsUserIdentityProfiling"),
//   	},
//   	RareProfiledBucketsAccountProfiling: []*string{
//   		jsii.String("rareProfiledBucketsAccountProfiling"),
//   	},
//   	RareProfiledBucketsUserIdentityProfiling: []*string{
//   		jsii.String("rareProfiledBucketsUserIdentityProfiling"),
//   	},
//   	RareProfiledUserAgentsAccountProfiling: []*string{
//   		jsii.String("rareProfiledUserAgentsAccountProfiling"),
//   	},
//   	RareProfiledUserAgentsUserIdentityProfiling: []*string{
//   		jsii.String("rareProfiledUserAgentsUserIdentityProfiling"),
//   	},
//   	RareProfiledUserNamesAccountProfiling: []*string{
//   		jsii.String("rareProfiledUserNamesAccountProfiling"),
//   	},
//   	RareProfiledUserNamesBucketProfiling: []*string{
//   		jsii.String("rareProfiledUserNamesBucketProfiling"),
//   	},
//   	RareProfiledUserTypesAccountProfiling: []*string{
//   		jsii.String("rareProfiledUserTypesAccountProfiling"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_ProfiledBehavior struct {
	// frequentProfiledAPIsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledAPIsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledApIsAccountProfiling *[]*string `field:"optional" json:"frequentProfiledApIsAccountProfiling" yaml:"frequentProfiledApIsAccountProfiling"`
	// frequentProfiledAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledApIsUserIdentityProfiling *[]*string `field:"optional" json:"frequentProfiledApIsUserIdentityProfiling" yaml:"frequentProfiledApIsUserIdentityProfiling"`
	// frequentProfiledASNsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledASNsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledAsNsAccountProfiling *[]*string `field:"optional" json:"frequentProfiledAsNsAccountProfiling" yaml:"frequentProfiledAsNsAccountProfiling"`
	// frequentProfiledASNsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledASNsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledAsNsBucketProfiling *[]*string `field:"optional" json:"frequentProfiledAsNsBucketProfiling" yaml:"frequentProfiledAsNsBucketProfiling"`
	// frequentProfiledASNsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledASNsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledAsNsUserIdentityProfiling *[]*string `field:"optional" json:"frequentProfiledAsNsUserIdentityProfiling" yaml:"frequentProfiledAsNsUserIdentityProfiling"`
	// frequentProfiledBucketsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledBucketsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledBucketsAccountProfiling *[]*string `field:"optional" json:"frequentProfiledBucketsAccountProfiling" yaml:"frequentProfiledBucketsAccountProfiling"`
	// frequentProfiledBucketsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledBucketsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledBucketsUserIdentityProfiling *[]*string `field:"optional" json:"frequentProfiledBucketsUserIdentityProfiling" yaml:"frequentProfiledBucketsUserIdentityProfiling"`
	// frequentProfiledUserAgentsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledUserAgentsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledUserAgentsAccountProfiling *[]*string `field:"optional" json:"frequentProfiledUserAgentsAccountProfiling" yaml:"frequentProfiledUserAgentsAccountProfiling"`
	// frequentProfiledUserAgentsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledUserAgentsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledUserAgentsUserIdentityProfiling *[]*string `field:"optional" json:"frequentProfiledUserAgentsUserIdentityProfiling" yaml:"frequentProfiledUserAgentsUserIdentityProfiling"`
	// frequentProfiledUserNamesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledUserNamesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledUserNamesAccountProfiling *[]*string `field:"optional" json:"frequentProfiledUserNamesAccountProfiling" yaml:"frequentProfiledUserNamesAccountProfiling"`
	// frequentProfiledUserNamesBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledUserNamesBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledUserNamesBucketProfiling *[]*string `field:"optional" json:"frequentProfiledUserNamesBucketProfiling" yaml:"frequentProfiledUserNamesBucketProfiling"`
	// frequentProfiledUserTypesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of frequentProfiledUserTypesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FrequentProfiledUserTypesAccountProfiling *[]*string `field:"optional" json:"frequentProfiledUserTypesAccountProfiling" yaml:"frequentProfiledUserTypesAccountProfiling"`
	// infrequentProfiledAPIsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledAPIsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledApIsAccountProfiling *[]*string `field:"optional" json:"infrequentProfiledApIsAccountProfiling" yaml:"infrequentProfiledApIsAccountProfiling"`
	// infrequentProfiledAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledApIsUserIdentityProfiling *[]*string `field:"optional" json:"infrequentProfiledApIsUserIdentityProfiling" yaml:"infrequentProfiledApIsUserIdentityProfiling"`
	// infrequentProfiledASNsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledASNsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledAsNsAccountProfiling *[]*string `field:"optional" json:"infrequentProfiledAsNsAccountProfiling" yaml:"infrequentProfiledAsNsAccountProfiling"`
	// infrequentProfiledASNsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledASNsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledAsNsBucketProfiling *[]*string `field:"optional" json:"infrequentProfiledAsNsBucketProfiling" yaml:"infrequentProfiledAsNsBucketProfiling"`
	// infrequentProfiledASNsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledASNsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledAsNsUserIdentityProfiling *[]*string `field:"optional" json:"infrequentProfiledAsNsUserIdentityProfiling" yaml:"infrequentProfiledAsNsUserIdentityProfiling"`
	// infrequentProfiledBucketsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledBucketsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledBucketsAccountProfiling *[]*string `field:"optional" json:"infrequentProfiledBucketsAccountProfiling" yaml:"infrequentProfiledBucketsAccountProfiling"`
	// infrequentProfiledBucketsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledBucketsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledBucketsUserIdentityProfiling *[]*string `field:"optional" json:"infrequentProfiledBucketsUserIdentityProfiling" yaml:"infrequentProfiledBucketsUserIdentityProfiling"`
	// infrequentProfiledUserAgentsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledUserAgentsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledUserAgentsAccountProfiling *[]*string `field:"optional" json:"infrequentProfiledUserAgentsAccountProfiling" yaml:"infrequentProfiledUserAgentsAccountProfiling"`
	// infrequentProfiledUserAgentsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledUserAgentsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledUserAgentsUserIdentityProfiling *[]*string `field:"optional" json:"infrequentProfiledUserAgentsUserIdentityProfiling" yaml:"infrequentProfiledUserAgentsUserIdentityProfiling"`
	// infrequentProfiledUserNamesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledUserNamesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledUserNamesAccountProfiling *[]*string `field:"optional" json:"infrequentProfiledUserNamesAccountProfiling" yaml:"infrequentProfiledUserNamesAccountProfiling"`
	// infrequentProfiledUserNamesBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledUserNamesBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledUserNamesBucketProfiling *[]*string `field:"optional" json:"infrequentProfiledUserNamesBucketProfiling" yaml:"infrequentProfiledUserNamesBucketProfiling"`
	// infrequentProfiledUserTypesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of infrequentProfiledUserTypesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InfrequentProfiledUserTypesAccountProfiling *[]*string `field:"optional" json:"infrequentProfiledUserTypesAccountProfiling" yaml:"infrequentProfiledUserTypesAccountProfiling"`
	// numberOfHistoricalDailyAvgAPIsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfHistoricalDailyAvgAPIsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfHistoricalDailyAvgApIsBucketProfiling *[]*string `field:"optional" json:"numberOfHistoricalDailyAvgApIsBucketProfiling" yaml:"numberOfHistoricalDailyAvgApIsBucketProfiling"`
	// numberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling *[]*string `field:"optional" json:"numberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling" yaml:"numberOfHistoricalDailyAvgApIsBucketUserIdentityProfiling"`
	// numberOfHistoricalDailyAvgAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfHistoricalDailyAvgAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfHistoricalDailyAvgApIsUserIdentityProfiling *[]*string `field:"optional" json:"numberOfHistoricalDailyAvgApIsUserIdentityProfiling" yaml:"numberOfHistoricalDailyAvgApIsUserIdentityProfiling"`
	// numberOfHistoricalDailyMaxAPIsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfHistoricalDailyMaxAPIsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfHistoricalDailyMaxApIsBucketProfiling *[]*string `field:"optional" json:"numberOfHistoricalDailyMaxApIsBucketProfiling" yaml:"numberOfHistoricalDailyMaxApIsBucketProfiling"`
	// numberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling *[]*string `field:"optional" json:"numberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling" yaml:"numberOfHistoricalDailyMaxApIsBucketUserIdentityProfiling"`
	// numberOfHistoricalDailyMaxAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of numberOfHistoricalDailyMaxAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NumberOfHistoricalDailyMaxApIsUserIdentityProfiling *[]*string `field:"optional" json:"numberOfHistoricalDailyMaxApIsUserIdentityProfiling" yaml:"numberOfHistoricalDailyMaxApIsUserIdentityProfiling"`
	// rareProfiledAPIsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledAPIsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledApIsAccountProfiling *[]*string `field:"optional" json:"rareProfiledApIsAccountProfiling" yaml:"rareProfiledApIsAccountProfiling"`
	// rareProfiledAPIsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledAPIsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledApIsUserIdentityProfiling *[]*string `field:"optional" json:"rareProfiledApIsUserIdentityProfiling" yaml:"rareProfiledApIsUserIdentityProfiling"`
	// rareProfiledASNsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledASNsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledAsNsAccountProfiling *[]*string `field:"optional" json:"rareProfiledAsNsAccountProfiling" yaml:"rareProfiledAsNsAccountProfiling"`
	// rareProfiledASNsBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledASNsBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledAsNsBucketProfiling *[]*string `field:"optional" json:"rareProfiledAsNsBucketProfiling" yaml:"rareProfiledAsNsBucketProfiling"`
	// rareProfiledASNsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledASNsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledAsNsUserIdentityProfiling *[]*string `field:"optional" json:"rareProfiledAsNsUserIdentityProfiling" yaml:"rareProfiledAsNsUserIdentityProfiling"`
	// rareProfiledBucketsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledBucketsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledBucketsAccountProfiling *[]*string `field:"optional" json:"rareProfiledBucketsAccountProfiling" yaml:"rareProfiledBucketsAccountProfiling"`
	// rareProfiledBucketsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledBucketsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledBucketsUserIdentityProfiling *[]*string `field:"optional" json:"rareProfiledBucketsUserIdentityProfiling" yaml:"rareProfiledBucketsUserIdentityProfiling"`
	// rareProfiledUserAgentsAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledUserAgentsAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledUserAgentsAccountProfiling *[]*string `field:"optional" json:"rareProfiledUserAgentsAccountProfiling" yaml:"rareProfiledUserAgentsAccountProfiling"`
	// rareProfiledUserAgentsUserIdentityProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledUserAgentsUserIdentityProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledUserAgentsUserIdentityProfiling *[]*string `field:"optional" json:"rareProfiledUserAgentsUserIdentityProfiling" yaml:"rareProfiledUserAgentsUserIdentityProfiling"`
	// rareProfiledUserNamesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledUserNamesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledUserNamesAccountProfiling *[]*string `field:"optional" json:"rareProfiledUserNamesAccountProfiling" yaml:"rareProfiledUserNamesAccountProfiling"`
	// rareProfiledUserNamesBucketProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledUserNamesBucketProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledUserNamesBucketProfiling *[]*string `field:"optional" json:"rareProfiledUserNamesBucketProfiling" yaml:"rareProfiledUserNamesBucketProfiling"`
	// rareProfiledUserTypesAccountProfiling property.
	//
	// Specify an array of string values to match this event if the actual value of rareProfiledUserTypesAccountProfiling is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RareProfiledUserTypesAccountProfiling *[]*string `field:"optional" json:"rareProfiledUserTypesAccountProfiling" yaml:"rareProfiledUserTypesAccountProfiling"`
}

