package awsglue


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
	// `CfnCrawler.DeltaTargetProperty.ConnectionName`.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// `CfnCrawler.DeltaTargetProperty.CreateNativeDeltaTable`.
	CreateNativeDeltaTable interface{} `field:"optional" json:"createNativeDeltaTable" yaml:"createNativeDeltaTable"`
	// `CfnCrawler.DeltaTargetProperty.DeltaTables`.
	DeltaTables *[]*string `field:"optional" json:"deltaTables" yaml:"deltaTables"`
	// `CfnCrawler.DeltaTargetProperty.WriteManifest`.
	WriteManifest interface{} `field:"optional" json:"writeManifest" yaml:"writeManifest"`
}

