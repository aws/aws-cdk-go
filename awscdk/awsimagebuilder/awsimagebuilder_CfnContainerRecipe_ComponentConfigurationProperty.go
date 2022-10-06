package awsimagebuilder


// Configuration details of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &componentConfigurationProperty{
//   	componentArn: jsii.String("componentArn"),
//   	parameters: []interface{}{
//   		&componentParameterProperty{
//   			name: jsii.String("name"),
//   			value: []*string{
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

