package previewawsdataexchangeevents


// Type definition for Scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scope_ := &Scope{
//   	LakeFormationTagPolicies: []LakeFormationTagPolicyDetails{
//   		&LakeFormationTagPolicyDetails{
//   			Database: []*string{
//   				jsii.String("database"),
//   			},
//   			Table: []*string{
//   				jsii.String("table"),
//   			},
//   		},
//   	},
//   	RedshiftDataShares: []RedshiftDataShareDetails{
//   		&RedshiftDataShareDetails{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   			Database: []*string{
//   				jsii.String("database"),
//   			},
//   			Function: []*string{
//   				jsii.String("function"),
//   			},
//   			Schema: []*string{
//   				jsii.String("schema"),
//   			},
//   			Table: []*string{
//   				jsii.String("table"),
//   			},
//   			View: []*string{
//   				jsii.String("view"),
//   			},
//   		},
//   	},
//   	S3DataAccesses: []S3DataAccessDetails{
//   		&S3DataAccessDetails{
//   			KeyPrefixes: []*string{
//   				jsii.String("keyPrefixes"),
//   			},
//   			Keys: []*string{
//   				jsii.String("keys"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DeprecationPlannedForDataSet_Scope struct {
	// LakeFormationTagPolicies property.
	//
	// Specify an array of string values to match this event if the actual value of LakeFormationTagPolicies is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LakeFormationTagPolicies *[]*DeprecationPlannedForDataSet_LakeFormationTagPolicyDetails `field:"optional" json:"lakeFormationTagPolicies" yaml:"lakeFormationTagPolicies"`
	// RedshiftDataShares property.
	//
	// Specify an array of string values to match this event if the actual value of RedshiftDataShares is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RedshiftDataShares *[]*DeprecationPlannedForDataSet_RedshiftDataShareDetails `field:"optional" json:"redshiftDataShares" yaml:"redshiftDataShares"`
	// S3DataAccesses property.
	//
	// Specify an array of string values to match this event if the actual value of S3DataAccesses is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3DataAccesses *[]*DeprecationPlannedForDataSet_S3DataAccessDetails `field:"optional" json:"s3DataAccesses" yaml:"s3DataAccesses"`
}

