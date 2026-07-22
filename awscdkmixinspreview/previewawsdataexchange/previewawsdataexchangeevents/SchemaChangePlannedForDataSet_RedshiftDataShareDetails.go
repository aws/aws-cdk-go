package previewawsdataexchangeevents


// Type definition for RedshiftDataShareDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftDataShareDetails := &RedshiftDataShareDetails{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	Database: []*string{
//   		jsii.String("database"),
//   	},
//   	Function: []*string{
//   		jsii.String("function"),
//   	},
//   	Schema: []*string{
//   		jsii.String("schema"),
//   	},
//   	Table: []*string{
//   		jsii.String("table"),
//   	},
//   	View: []*string{
//   		jsii.String("view"),
//   	},
//   }
//
// Experimental.
type SchemaChangePlannedForDataSet_RedshiftDataShareDetails struct {
	// Arn property.
	//
	// Specify an array of string values to match this event if the actual value of Arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// Database property.
	//
	// Specify an array of string values to match this event if the actual value of Database is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Database *[]*string `field:"optional" json:"database" yaml:"database"`
	// Function property.
	//
	// Specify an array of string values to match this event if the actual value of Function is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Function *[]*string `field:"optional" json:"function" yaml:"function"`
	// Schema property.
	//
	// Specify an array of string values to match this event if the actual value of Schema is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Schema *[]*string `field:"optional" json:"schema" yaml:"schema"`
	// Table property.
	//
	// Specify an array of string values to match this event if the actual value of Table is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Table *[]*string `field:"optional" json:"table" yaml:"table"`
	// View property.
	//
	// Specify an array of string values to match this event if the actual value of View is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	View *[]*string `field:"optional" json:"view" yaml:"view"`
}

