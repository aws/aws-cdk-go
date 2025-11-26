package previewawsquicksightmixins


// The parameters for OpenSearch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   amazonElasticsearchParametersProperty := &AmazonElasticsearchParametersProperty{
//   	Domain: jsii.String("domain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-amazonelasticsearchparameters.html
//
type CfnDataSourcePropsMixin_AmazonElasticsearchParametersProperty struct {
	// The OpenSearch domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-amazonelasticsearchparameters.html#cfn-quicksight-datasource-amazonelasticsearchparameters-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

