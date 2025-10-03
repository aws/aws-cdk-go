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
//   	RoleConfiguration: &RoleConfigurationProperty{
//   		SourceRoleArn: jsii.String("sourceRoleArn"),
//   		TargetRoleArn: jsii.String("targetRoleArn"),
//   	},
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
	// The Amazon Managed Service for Prometheus workspace the scraper sends metrics to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The configuration in use by the scraper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-scrapeconfiguration
	//
	ScrapeConfiguration interface{} `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// The Amazon EKS cluster from which the scraper collects metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// An optional user-assigned scraper alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The role configuration in an Amazon Managed Service for Prometheus scraper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-roleconfiguration
	//
	RoleConfiguration interface{} `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// (Optional) The list of tag keys and values associated with the scraper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

