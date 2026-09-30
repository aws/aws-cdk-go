package awsscn


// The DataIntegrationFlow transformation parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataIntegrationFlowTransformationProperty := &DataIntegrationFlowTransformationProperty{
//   	TransformationType: jsii.String("transformationType"),
//
//   	// the properties below are optional
//   	SqlTransformation: &SqlTransformationConfigurationProperty{
//   		Query: jsii.String("query"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtransformation.html
//
type CfnDataIntegrationFlow_DataIntegrationFlowTransformationProperty struct {
	// The transformation type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtransformation.html#cfn-scn-dataintegrationflow-dataintegrationflowtransformation-transformationtype
	//
	TransformationType *string `field:"required" json:"transformationType" yaml:"transformationType"`
	// The SQL transformation configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtransformation.html#cfn-scn-dataintegrationflow-dataintegrationflowtransformation-sqltransformation
	//
	SqlTransformation interface{} `field:"optional" json:"sqlTransformation" yaml:"sqlTransformation"`
}

