package previewawsimagebuildermixins


// Configuration details of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html
//
type CfnImageRecipePropsMixin_ComponentConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-componentarn
	//
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// A group of parameter settings that Image Builder uses to configure the component for a specific recipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

