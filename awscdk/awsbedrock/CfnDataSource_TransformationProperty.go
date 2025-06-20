package awsbedrock


// A custom processing step for documents moving through a data source ingestion pipeline.
//
// To process documents after they have been converted into chunks, set the step to apply to `POST_CHUNKING` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformationProperty := &TransformationProperty{
//   	StepToApply: jsii.String("stepToApply"),
//   	TransformationFunction: &TransformationFunctionProperty{
//   		TransformationLambdaConfiguration: &TransformationLambdaConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformation.html
//
type CfnDataSource_TransformationProperty struct {
	// When the service applies the transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformation.html#cfn-bedrock-datasource-transformation-steptoapply
	//
	StepToApply *string `field:"required" json:"stepToApply" yaml:"stepToApply"`
	// A Lambda function that processes documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformation.html#cfn-bedrock-datasource-transformation-transformationfunction
	//
	TransformationFunction interface{} `field:"required" json:"transformationFunction" yaml:"transformationFunction"`
}

