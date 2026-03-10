package previewawssagemakerevents


// Type definition for AlgorithmSpecificationItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   algorithmSpecificationItem := &AlgorithmSpecificationItem{
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Regex: []*string{
//   		jsii.String("regex"),
//   	},
//   }
//
// Experimental.
type SageMakerHyperParameterTuningJobStateChange_AlgorithmSpecificationItem struct {
	// Name property.
	//
	// Specify an array of string values to match this event if the actual value of Name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// Regex property.
	//
	// Specify an array of string values to match this event if the actual value of Regex is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Regex *[]*string `field:"optional" json:"regex" yaml:"regex"`
}

