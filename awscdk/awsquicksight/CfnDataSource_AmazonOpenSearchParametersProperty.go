package awsquicksight


// The parameters for OpenSearch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonOpenSearchParametersProperty := &AmazonOpenSearchParametersProperty{
//   	Domain: jsii.String("domain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-amazonopensearchparameters.html
//
type CfnDataSource_AmazonOpenSearchParametersProperty struct {
	// The OpenSearch domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-amazonopensearchparameters.html#cfn-quicksight-datasource-amazonopensearchparameters-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

