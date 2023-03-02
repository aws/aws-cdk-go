package awsappflow


// A map used to store task-related information.
//
// The execution service looks for particular information based on the `TaskType` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskPropertiesObjectProperty := &taskPropertiesObjectProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFlow_TaskPropertiesObjectProperty struct {
	// The task property key.
	//
	// *Allowed Values* : `VALUE | VALUES | DATA_TYPE | UPPER_BOUND | LOWER_BOUND | SOURCE_DATA_TYPE | DESTINATION_DATA_TYPE | VALIDATION_ACTION | MASK_VALUE | MASK_LENGTH | TRUNCATE_LENGTH | MATH_OPERATION_FIELDS_ORDER | CONCAT_FORMAT | SUBFIELD_CATEGORY_MAP` | `EXCLUDE_SOURCE_FIELDS_LIST`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The task property value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

