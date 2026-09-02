package previewawsdataexchangeevents


// Type definition for LakeFormationTagPolicyDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lakeFormationTagPolicyDetails := &LakeFormationTagPolicyDetails{
//   	Database: []*string{
//   		jsii.String("database"),
//   	},
//   	Table: []*string{
//   		jsii.String("table"),
//   	},
//   }
//
// Experimental.
type DeprecationPlannedForDataSet_LakeFormationTagPolicyDetails struct {
	// Database property.
	//
	// Specify an array of string values to match this event if the actual value of Database is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Database *[]*string `field:"optional" json:"database" yaml:"database"`
	// Table property.
	//
	// Specify an array of string values to match this event if the actual value of Table is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Table *[]*string `field:"optional" json:"table" yaml:"table"`
}

