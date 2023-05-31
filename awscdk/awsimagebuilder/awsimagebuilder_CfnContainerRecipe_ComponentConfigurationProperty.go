package awsimagebuilder


// Configuration details of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &ComponentConfigurationProperty{
//   	ComponentArn: jsii.String("componentArn"),
//   	Parameters: []interface{}{
//   		&ComponentParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnContainerRecipe_ComponentConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// `CfnContainerRecipe.ComponentConfigurationProperty.Parameters`.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

