package awsaps


// Configuration for OpenSearch exporter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   openSearchConfigurationProperty := &OpenSearchConfigurationProperty{
//   	DomainArn: jsii.String("domainArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-opensearchconfiguration.html
//
type CfnScraperPropsMixin_OpenSearchConfigurationProperty struct {
	// ARN of an OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-opensearchconfiguration.html#cfn-aps-scraper-opensearchconfiguration-domainarn
	//
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
}

