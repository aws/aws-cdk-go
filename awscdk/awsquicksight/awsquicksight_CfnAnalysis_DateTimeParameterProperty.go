package awsquicksight


// A date-time parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeParameterProperty := &dateTimeParameterProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnAnalysis_DateTimeParameterProperty struct {
	// A display name for the date-time parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the date-time parameter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

