package awsglue


// Configuration for synchronizing a federated catalog with its external data source.
//
// Removing this property does not disable synchronization: Glue keeps the stored configuration for anything omitted from an update. To stop synchronizing, set the Status of each mode to DISABLED.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   syncConfigurationProperty := &SyncConfigurationProperty{
//   	AssetsSyncConfiguration: &AssetsSyncConfigurationProperty{
//   		ManagedCrawlerDetails: &ManagedCrawlerDetailsProperty{
//   			CrawlerArn: jsii.String("crawlerArn"),
//   			Error: &ErrorDetailProperty{
//   				ErrorCode: jsii.String("errorCode"),
//   				ErrorMessage: jsii.String("errorMessage"),
//   			},
//   			Status: jsii.String("status"),
//   		},
//   		RefreshIntervalMinutes: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	LineageSyncConfiguration: &LineageSyncConfigurationProperty{
//   		ManagedCrawlerDetails: &ManagedCrawlerDetailsProperty{
//   			CrawlerArn: jsii.String("crawlerArn"),
//   			Error: &ErrorDetailProperty{
//   				ErrorCode: jsii.String("errorCode"),
//   				ErrorMessage: jsii.String("errorMessage"),
//   			},
//   			Status: jsii.String("status"),
//   		},
//   		RefreshIntervalMinutes: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-syncconfiguration.html
//
type CfnCatalogPropsMixin_SyncConfigurationProperty struct {
	// Configuration for synchronizing assets from the external data source.
	//
	// Omitting this property on an update leaves the stored configuration unchanged; set Status to DISABLED to stop synchronizing assets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-syncconfiguration.html#cfn-glue-catalog-syncconfiguration-assetssyncconfiguration
	//
	AssetsSyncConfiguration interface{} `field:"optional" json:"assetsSyncConfiguration" yaml:"assetsSyncConfiguration"`
	// Configuration for synchronizing lineage from the external data source.
	//
	// Omitting this property on an update leaves the stored configuration unchanged; set Status to DISABLED to stop synchronizing lineage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-syncconfiguration.html#cfn-glue-catalog-syncconfiguration-lineagesyncconfiguration
	//
	LineageSyncConfiguration interface{} `field:"optional" json:"lineageSyncConfiguration" yaml:"lineageSyncConfiguration"`
	// The IAM role, in the catalog owner's account, that the crawlers Glue manages on your behalf assume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-syncconfiguration.html#cfn-glue-catalog-syncconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

