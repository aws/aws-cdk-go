package awsentityresolution


// A list of `OutputAttribute` objects, each of which have the fields `Name` and `Hashed` .
//
// Each of these objects selects a column to be included in the output table, and whether the values of the column should be hashed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputSourceProperty := &OutputSourceProperty{
//   	Output: []interface{}{
//   		&OutputAttributeProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Hashed: jsii.Boolean(false),
//   		},
//   	},
//   	OutputS3Path: jsii.String("outputS3Path"),
//
//   	// the properties below are optional
//   	ApplyNormalization: jsii.Boolean(false),
//   	KmsArn: jsii.String("kmsArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html
//
type CfnMatchingWorkflow_OutputSourceProperty struct {
	// A list of `OutputAttribute` objects, each of which have the fields `Name` and `Hashed` .
	//
	// Each of these objects selects a column to be included in the output table, and whether the values of the column should be hashed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-output
	//
	Output interface{} `field:"required" json:"output" yaml:"output"`
	// The S3 path to which AWS Entity Resolution will write the output table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-outputs3path
	//
	OutputS3Path *string `field:"required" json:"outputS3Path" yaml:"outputS3Path"`
	// Normalizes the attributes defined in the schema in the input data.
	//
	// For example, if an attribute has an `AttributeType` of `PHONE_NUMBER` , and the data in the input table is in a format of 1234567890, AWS Entity Resolution will normalize this field in the output to (123)-456-7890.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-applynormalization
	//
	ApplyNormalization interface{} `field:"optional" json:"applyNormalization" yaml:"applyNormalization"`
	// Customer KMS ARN for encryption at rest.
	//
	// If not provided, system will use an AWS Entity Resolution managed KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-outputsource.html#cfn-entityresolution-matchingworkflow-outputsource-kmsarn
	//
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

