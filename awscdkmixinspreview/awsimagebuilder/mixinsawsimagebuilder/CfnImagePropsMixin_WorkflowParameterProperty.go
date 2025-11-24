package mixinsawsimagebuilder


// Contains a key/value pair that sets the named workflow parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workflowParameterProperty := &WorkflowParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-workflowparameter.html
//
type CfnImagePropsMixin_WorkflowParameterProperty struct {
	// The name of the workflow parameter to set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-workflowparameter.html#cfn-imagebuilder-image-workflowparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Sets the value for the named workflow parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-workflowparameter.html#cfn-imagebuilder-image-workflowparameter-value
	//
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

