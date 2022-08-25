package awsimagebuilder


// Contains a key/value pair that sets the named component parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentParameterProperty := &componentParameterProperty{
//   	name: jsii.String("name"),
//   	value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
type CfnImageRecipe_ComponentParameterProperty struct {
	// The name of the component parameter to set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

