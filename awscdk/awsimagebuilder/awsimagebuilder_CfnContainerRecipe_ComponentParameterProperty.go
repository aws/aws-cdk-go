package awsimagebuilder


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
type CfnContainerRecipe_ComponentParameterProperty struct {
	// `CfnContainerRecipe.ComponentParameterProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnContainerRecipe.ComponentParameterProperty.Value`.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

