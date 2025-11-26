package previewawsdatabrewmixins


// Represents a transformation and associated parameters that are used to apply a change to an AWS Glue DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	Operation: jsii.String("operation"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-action.html
//
type CfnRecipePropsMixin_ActionProperty struct {
	// The name of a valid DataBrew transformation to be performed on the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-action.html#cfn-databrew-recipe-action-operation
	//
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Contextual parameters for the transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-action.html#cfn-databrew-recipe-action-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

