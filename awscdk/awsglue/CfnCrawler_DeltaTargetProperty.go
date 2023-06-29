package awsglue


// Specifies a Delta data store to crawl one or more Delta tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaTargetProperty := &DeltaTargetProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	CreateNativeDeltaTable: jsii.Boolean(false),
//   	DeltaTables: []*string{
//   		jsii.String("deltaTables"),
//   	},
//   	WriteManifest: jsii.Boolean(false),
//   }
//
type CfnCrawler_DeltaTargetProperty struct {
	// The name of the connection to use to connect to the Delta table target.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.
	CreateNativeDeltaTable interface{} `field:"optional" json:"createNativeDeltaTable" yaml:"createNativeDeltaTable"`
	// A list of the Amazon S3 paths to the Delta tables.
	DeltaTables *[]*string `field:"optional" json:"deltaTables" yaml:"deltaTables"`
	// Specifies whether to write the manifest files to the Delta table path.
	WriteManifest interface{} `field:"optional" json:"writeManifest" yaml:"writeManifest"`
}

