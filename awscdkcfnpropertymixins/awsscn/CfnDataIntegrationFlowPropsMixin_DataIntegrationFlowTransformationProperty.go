package awsscn


// The DataIntegrationFlow transformation parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataIntegrationFlowTransformationProperty := &DataIntegrationFlowTransformationProperty{
//   	SqlTransformation: &SqlTransformationConfigurationProperty{
//   		Query: jsii.String("query"),
//   	},
//   	TransformationType: jsii.String("transformationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtransformation.html
//
type CfnDataIntegrationFlowPropsMixin_DataIntegrationFlowTransformationProperty struct {
	// The SQL transformation configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtransformation.html#cfn-scn-dataintegrationflow-dataintegrationflowtransformation-sqltransformation
	//
	SqlTransformation interface{} `field:"optional" json:"sqlTransformation" yaml:"sqlTransformation"`
	// The transformation type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowtransformation.html#cfn-scn-dataintegrationflow-dataintegrationflowtransformation-transformationtype
	//
	TransformationType *string `field:"optional" json:"transformationType" yaml:"transformationType"`
}

