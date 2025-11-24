package mixinsawsopensearchserverless


// Index settings for the OpenSearch Serverless index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   indexSettingsProperty := &IndexSettingsProperty{
//   	Index: &IndexProperty{
//   		Knn: jsii.Boolean(false),
//   		KnnAlgoParamEfSearch: jsii.Number(123),
//   		RefreshInterval: jsii.String("refreshInterval"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-indexsettings.html
//
type CfnIndexPropsMixin_IndexSettingsProperty struct {
	// Index settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-index-indexsettings.html#cfn-opensearchserverless-index-indexsettings-index
	//
	Index interface{} `field:"optional" json:"index" yaml:"index"`
}

