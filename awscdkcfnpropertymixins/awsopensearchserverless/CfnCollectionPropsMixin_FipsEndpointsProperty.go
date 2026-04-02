package awsopensearchserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fipsEndpointsProperty := &FipsEndpointsProperty{
//   	CollectionEndpoint: jsii.String("collectionEndpoint"),
//   	DashboardEndpoint: jsii.String("dashboardEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-fipsendpoints.html
//
type CfnCollectionPropsMixin_FipsEndpointsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-fipsendpoints.html#cfn-opensearchserverless-collection-fipsendpoints-collectionendpoint
	//
	CollectionEndpoint *string `field:"optional" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collection-fipsendpoints.html#cfn-opensearchserverless-collection-fipsendpoints-dashboardendpoint
	//
	DashboardEndpoint *string `field:"optional" json:"dashboardEndpoint" yaml:"dashboardEndpoint"`
}

