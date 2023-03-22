package awsquicksight


// A string parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringParameterProperty := &stringParameterProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnDashboard_StringParameterProperty struct {
	// A display name for a string parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of a string parameter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

