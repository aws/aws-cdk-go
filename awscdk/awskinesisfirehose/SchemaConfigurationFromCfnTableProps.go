package awskinesisfirehose


// Options for creating a Schema for record format conversion from a `glue.CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaConfigurationFromCfnTableProps := &SchemaConfigurationFromCfnTableProps{
//   	Region: jsii.String("region"),
//   	VersionId: jsii.String("versionId"),
//   }
//
type SchemaConfigurationFromCfnTableProps struct {
	// The region of the database the table is in.
	// Default: the region of the stack that contains the table reference is used.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Specifies the table version for the output data schema.
	//
	// if set to `LATEST`, Firehose uses the most recent table version. This means that any updates to the table are automatically picked up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-schemaconfiguration.html#cfn-kinesisfirehose-deliverystream-schemaconfiguration-versionid
	//
	// Default: `LATEST`.
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

