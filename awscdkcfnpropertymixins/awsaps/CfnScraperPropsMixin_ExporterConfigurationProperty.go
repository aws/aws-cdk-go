package awsaps


// Exporter configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   exporterConfigurationProperty := &ExporterConfigurationProperty{
//   	OpenSearchConfiguration: &OpenSearchConfigurationProperty{
//   		DomainArn: jsii.String("domainArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-exporterconfiguration.html
//
type CfnScraperPropsMixin_ExporterConfigurationProperty struct {
	// Configuration for OpenSearch exporter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-exporterconfiguration.html#cfn-aps-scraper-exporterconfiguration-opensearchconfiguration
	//
	OpenSearchConfiguration interface{} `field:"optional" json:"openSearchConfiguration" yaml:"openSearchConfiguration"`
}

