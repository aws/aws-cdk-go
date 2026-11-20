package awsglue


// A FederatedCatalog structure that references an entity outside the Glue Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   federatedCatalogProperty := &FederatedCatalogProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	ConnectionType: jsii.String("connectionType"),
//   	Identifier: jsii.String("identifier"),
//   	SyncConfiguration: &SyncConfigurationProperty{
//   		AssetsSyncConfiguration: &AssetsSyncConfigurationProperty{
//   			ManagedCrawlerDetails: &ManagedCrawlerDetailsProperty{
//   				CrawlerArn: jsii.String("crawlerArn"),
//   				Error: &ErrorDetailProperty{
//   					ErrorCode: jsii.String("errorCode"),
//   					ErrorMessage: jsii.String("errorMessage"),
//   				},
//   				Status: jsii.String("status"),
//   			},
//   			RefreshIntervalMinutes: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		LineageSyncConfiguration: &LineageSyncConfigurationProperty{
//   			ManagedCrawlerDetails: &ManagedCrawlerDetailsProperty{
//   				CrawlerArn: jsii.String("crawlerArn"),
//   				Error: &ErrorDetailProperty{
//   					ErrorCode: jsii.String("errorCode"),
//   					ErrorMessage: jsii.String("errorMessage"),
//   				},
//   				Status: jsii.String("status"),
//   			},
//   			RefreshIntervalMinutes: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html
//
type CfnCatalogPropsMixin_FederatedCatalogProperty struct {
	// The name of the connection to an external data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html#cfn-glue-catalog-federatedcatalog-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The type of connection used to access the federated catalog.
	//
	// Synchronization is only supported for certain connection types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html#cfn-glue-catalog-federatedcatalog-connectiontype
	//
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// A unique identifier for the federated catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html#cfn-glue-catalog-federatedcatalog-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Configuration for synchronizing a federated catalog with its external data source.
	//
	// Removing this property does not disable synchronization: Glue keeps the stored configuration for anything omitted from an update. To stop synchronizing, set the Status of each mode to DISABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html#cfn-glue-catalog-federatedcatalog-syncconfiguration
	//
	SyncConfiguration interface{} `field:"optional" json:"syncConfiguration" yaml:"syncConfiguration"`
}

