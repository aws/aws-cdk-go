package awsscn


// The SQL transformation configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sqlTransformationConfigurationProperty := &SqlTransformationConfigurationProperty{
//   	Query: jsii.String("query"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-sqltransformationconfiguration.html
//
type CfnDataIntegrationFlowPropsMixin_SqlTransformationConfigurationProperty struct {
	// The transformation SQL query body based on SparkSQL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-sqltransformationconfiguration.html#cfn-scn-dataintegrationflow-sqltransformationconfiguration-query
	//
	Query *string `field:"optional" json:"query" yaml:"query"`
}

