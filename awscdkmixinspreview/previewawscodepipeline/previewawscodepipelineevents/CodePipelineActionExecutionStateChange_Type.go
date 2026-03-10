package previewawscodepipelineevents


// Type definition for Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   type := &Type{
//   	Category: []*string{
//   		jsii.String("category"),
//   	},
//   	Owner: []*string{
//   		jsii.String("owner"),
//   	},
//   	Provider: []*string{
//   		jsii.String("provider"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type CodePipelineActionExecutionStateChange_Type struct {
	// category property.
	//
	// Specify an array of string values to match this event if the actual value of category is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Category *[]*string `field:"optional" json:"category" yaml:"category"`
	// owner property.
	//
	// Specify an array of string values to match this event if the actual value of owner is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Owner *[]*string `field:"optional" json:"owner" yaml:"owner"`
	// provider property.
	//
	// Specify an array of string values to match this event if the actual value of provider is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Provider *[]*string `field:"optional" json:"provider" yaml:"provider"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

