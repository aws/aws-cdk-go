package mixinsawsbedrock


// Settings for customizing steps in the data source content ingestion pipeline.
//
// You can configure the data source to process documents with a Lambda function after they are parsed and converted into chunks. When you add a post-chunking transformation, the service stores chunked documents in an S3 bucket and invokes a Lambda function to process them.
//
// To process chunked documents with a Lambda function, define an S3 bucket path for input and output objects, and a transformation that specifies the Lambda function to invoke. You can use the Lambda function to customize how chunks are split, and the metadata for each chunk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customTransformationConfigurationProperty := &CustomTransformationConfigurationProperty{
//   	IntermediateStorage: &IntermediateStorageProperty{
//   		S3Location: &S3LocationProperty{
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	Transformations: []interface{}{
//   		&TransformationProperty{
//   			StepToApply: jsii.String("stepToApply"),
//   			TransformationFunction: &TransformationFunctionProperty{
//   				TransformationLambdaConfiguration: &TransformationLambdaConfigurationProperty{
//   					LambdaArn: jsii.String("lambdaArn"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-customtransformationconfiguration.html
//
type CfnDataSourcePropsMixin_CustomTransformationConfigurationProperty struct {
	// An S3 bucket path for input and output objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-customtransformationconfiguration.html#cfn-bedrock-datasource-customtransformationconfiguration-intermediatestorage
	//
	IntermediateStorage interface{} `field:"optional" json:"intermediateStorage" yaml:"intermediateStorage"`
	// A Lambda function that processes documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-customtransformationconfiguration.html#cfn-bedrock-datasource-customtransformationconfiguration-transformations
	//
	Transformations interface{} `field:"optional" json:"transformations" yaml:"transformations"`
}

