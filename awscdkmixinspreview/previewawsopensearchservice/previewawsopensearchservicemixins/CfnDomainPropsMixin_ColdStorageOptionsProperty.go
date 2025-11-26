package previewawsopensearchservicemixins


// Container for the parameters required to enable cold storage for an OpenSearch Service domain.
//
// For more information, see [Cold storage for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cold-storage.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   coldStorageOptionsProperty := &ColdStorageOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-coldstorageoptions.html
//
type CfnDomainPropsMixin_ColdStorageOptionsProperty struct {
	// Whether to enable or disable cold storage on the domain.
	//
	// You must enable UltraWarm storage to enable cold storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-coldstorageoptions.html#cfn-opensearchservice-domain-coldstorageoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

