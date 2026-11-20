package awsglue


// Configuration for synchronizing lineage from the external data source.
//
// Omitting this property on an update leaves the stored configuration unchanged; set Status to DISABLED to stop synchronizing lineage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   lineageSyncConfigurationProperty := &LineageSyncConfigurationProperty{
//   	ManagedCrawlerDetails: &ManagedCrawlerDetailsProperty{
//   		CrawlerArn: jsii.String("crawlerArn"),
//   		Error: &ErrorDetailProperty{
//   			ErrorCode: jsii.String("errorCode"),
//   			ErrorMessage: jsii.String("errorMessage"),
//   		},
//   		Status: jsii.String("status"),
//   	},
//   	RefreshIntervalMinutes: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-lineagesyncconfiguration.html
//
type CfnCatalogPropsMixin_LineageSyncConfigurationProperty struct {
	// The state of the crawler Glue manages on your behalf for a synchronization mode.
	//
	// Returned by the service; Glue rejects it as input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-lineagesyncconfiguration.html#cfn-glue-catalog-lineagesyncconfiguration-managedcrawlerdetails
	//
	ManagedCrawlerDetails interface{} `field:"optional" json:"managedCrawlerDetails" yaml:"managedCrawlerDetails"`
	// How often the crawler synchronizes lineage, in minutes.
	//
	// Glue enforces a service-side minimum that may be higher than the minimum accepted here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-lineagesyncconfiguration.html#cfn-glue-catalog-lineagesyncconfiguration-refreshintervalminutes
	//
	RefreshIntervalMinutes *float64 `field:"optional" json:"refreshIntervalMinutes" yaml:"refreshIntervalMinutes"`
	// Whether lineage synchronization is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-lineagesyncconfiguration.html#cfn-glue-catalog-lineagesyncconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

