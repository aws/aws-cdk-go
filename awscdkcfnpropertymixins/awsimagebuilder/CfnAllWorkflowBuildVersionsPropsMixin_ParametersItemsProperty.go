package awsimagebuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   parametersItemsProperty := &ParametersItemsProperty{
//   	DefaultValue: []*string{
//   		jsii.String("defaultValue"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-parametersitems.html
//
type CfnAllWorkflowBuildVersionsPropsMixin_ParametersItemsProperty struct {
	// The default value of this parameter if no input is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-parametersitems.html#cfn-imagebuilder-allworkflowbuildversions-parametersitems-defaultvalue
	//
	DefaultValue *[]*string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Describes this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-parametersitems.html#cfn-imagebuilder-allworkflowbuildversions-parametersitems-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of this input parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-parametersitems.html#cfn-imagebuilder-allworkflowbuildversions-parametersitems-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of input this parameter provides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-parametersitems.html#cfn-imagebuilder-allworkflowbuildversions-parametersitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

