package awsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnScraper`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScraperProps := &CfnScraperProps{
//   	Destination: &DestinationProperty{
//   		AmpConfiguration: &AmpConfigurationProperty{
//   			WorkspaceArn: jsii.String("workspaceArn"),
//   		},
//   	},
//   	ScrapeConfiguration: &ScrapeConfigurationProperty{
//   		ConfigurationBlob: jsii.String("configurationBlob"),
//   	},
//   	Source: &SourceProperty{
//   		EksConfiguration: &EksConfigurationProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//
//   			// the properties below are optional
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Alias: jsii.String("alias"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html
//
type CfnScraperProps struct {
	// Scraper metrics destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Scraper configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-scrapeconfiguration
	//
	ScrapeConfiguration interface{} `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// Scraper metrics source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Scraper alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

