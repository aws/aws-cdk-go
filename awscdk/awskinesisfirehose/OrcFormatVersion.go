package awskinesisfirehose


// The available WriterVersions for ORC output format.
//
// Example:
//   outputFormat := firehose.NewOrcOutputFormat(&OrcOutputFormatProps{
//   	FormatVersion: firehose.OrcFormatVersion_V0_11,
//   	BlockSize: awscdk.Size_Mebibytes(jsii.Number(256)),
//   	Compression: firehose.OrcCompression_NONE(),
//   	BloomFilterColumns: []*string{
//   		jsii.String("columnA"),
//   	},
//   	BloomFilterFalsePositiveProbability: jsii.Number(0.1),
//   	DictionaryKeyThreshold: jsii.Number(0.7),
//   	EnablePadding: jsii.Boolean(true),
//   	PaddingTolerance: jsii.Number(0.2),
//   	RowIndexStride: jsii.Number(9000),
//   	StripeSize: awscdk.Size_*Mebibytes(jsii.Number(32)),
//   })
//
type OrcFormatVersion string

const (
	// Use V0_11 ORC writer version when writing the output of the record transformation.
	OrcFormatVersion_V0_11 OrcFormatVersion = "V0_11"
	// Use V0_12 ORC writer version when writing the output of the record transformation.
	OrcFormatVersion_V0_12 OrcFormatVersion = "V0_12"
)

