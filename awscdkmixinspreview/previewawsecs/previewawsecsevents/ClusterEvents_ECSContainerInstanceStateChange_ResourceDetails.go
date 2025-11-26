package previewawsecsevents


// Type definition for ResourceDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceDetails := &ResourceDetails{
//   	DoubleValue: []*string{
//   		jsii.String("doubleValue"),
//   	},
//   	IntegerValue: []*string{
//   		jsii.String("integerValue"),
//   	},
//   	LongValue: []*string{
//   		jsii.String("longValue"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	StringSetValue: []*string{
//   		jsii.String("stringSetValue"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSContainerInstanceStateChange_ResourceDetails struct {
	// doubleValue property.
	//
	// Specify an array of string values to match this event if the actual value of doubleValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DoubleValue *[]*string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// integerValue property.
	//
	// Specify an array of string values to match this event if the actual value of integerValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IntegerValue *[]*string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// longValue property.
	//
	// Specify an array of string values to match this event if the actual value of longValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LongValue *[]*string `field:"optional" json:"longValue" yaml:"longValue"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// stringSetValue property.
	//
	// Specify an array of string values to match this event if the actual value of stringSetValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StringSetValue *[]*string `field:"optional" json:"stringSetValue" yaml:"stringSetValue"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

