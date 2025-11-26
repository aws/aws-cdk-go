package previewawsapsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnScraperPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScraperMixinProps := &CfnScraperMixinProps{
//   	Alias: jsii.String("alias"),
//   	Destination: &DestinationProperty{
//   		AmpConfiguration: &AmpConfigurationProperty{
//   			WorkspaceArn: jsii.String("workspaceArn"),
//   		},
//   	},
//   	RoleConfiguration: &RoleConfigurationProperty{
//   		SourceRoleArn: jsii.String("sourceRoleArn"),
//   		TargetRoleArn: jsii.String("targetRoleArn"),
//   	},
//   	ScrapeConfiguration: &ScrapeConfigurationProperty{
//   		ConfigurationBlob: jsii.String("configurationBlob"),
//   	},
//   	ScraperLoggingConfiguration: &ScraperLoggingConfigurationProperty{
//   		LoggingDestination: &ScraperLoggingDestinationProperty{
//   			CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   				LogGroupArn: jsii.String("logGroupArn"),
//   			},
//   		},
//   		ScraperComponents: []interface{}{
//   			&ScraperComponentProperty{
//   				Config: &ComponentConfigProperty{
//   					Options: map[string]*string{
//   						"optionsKey": jsii.String("options"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Source: &SourceProperty{
//   		EksConfiguration: &EksConfigurationProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html
//
type CfnScraperMixinProps struct {
	// An optional user-assigned scraper alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The Amazon Managed Service for Prometheus workspace the scraper sends metrics to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The role configuration in an Amazon Managed Service for Prometheus scraper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-roleconfiguration
	//
	RoleConfiguration interface{} `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// The configuration in use by the scraper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-scrapeconfiguration
	//
	ScrapeConfiguration interface{} `field:"optional" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// The definition of logging configuration in an Amazon Managed Service for Prometheus workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-scraperloggingconfiguration
	//
	ScraperLoggingConfiguration interface{} `field:"optional" json:"scraperLoggingConfiguration" yaml:"scraperLoggingConfiguration"`
	// The Amazon EKS cluster from which the scraper collects metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// (Optional) The list of tag keys and values associated with the scraper.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-scraper.html#cfn-aps-scraper-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

