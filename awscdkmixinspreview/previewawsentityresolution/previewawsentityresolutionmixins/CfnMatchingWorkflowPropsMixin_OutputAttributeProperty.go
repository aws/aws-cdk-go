package previewawsentityresolutionmixins


// A list of `OutputAttribute` objects, each of which have the fields `Name` and `Hashed` .
//
// Each of these objects selects a column to be included in the output table, and whether the values of the column should be hashed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputAttributeProperty := &OutputAttributeProperty{
//   	Hashed: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputattribute.html
//
type CfnMatchingWorkflowPropsMixin_OutputAttributeProperty struct {
	// Enables the ability to hash the column values in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputattribute.html#cfn-entityresolution-matchingworkflow-outputattribute-hashed
	//
	Hashed interface{} `field:"optional" json:"hashed" yaml:"hashed"`
	// A name of a column to be written to the output.
	//
	// This must be an `InputField` name in the schema mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputattribute.html#cfn-entityresolution-matchingworkflow-outputattribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

