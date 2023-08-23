package awscdk


// The type of output that a bundling operation is producing.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
//   	Path: jsii.String("/path/to/asset"),
//   	Bundling: &BundlingOptions{
//   		Image: cdk.DockerImage_FromRegistry(jsii.String("alpine")),
//   		Command: []*string{
//   			jsii.String("command-that-produces-an-archive.sh"),
//   		},
//   		OutputType: cdk.BundlingOutput_NOT_ARCHIVED,
//   	},
//   })
//
type BundlingOutput string

const (
	// The bundling output directory includes a single .zip or .jar file which will be used as the final bundle. If the output directory does not include exactly a single archive, bundling will fail.
	BundlingOutput_ARCHIVED BundlingOutput = "ARCHIVED"
	// The bundling output directory contains one or more files which will be archived and uploaded as a .zip file to S3.
	BundlingOutput_NOT_ARCHIVED BundlingOutput = "NOT_ARCHIVED"
	// If the bundling output directory contains a single archive file (zip or jar) it will be used as the bundle output as-is.
	//
	// Otherwise, all the files in the bundling output directory will be zipped.
	BundlingOutput_AUTO_DISCOVER BundlingOutput = "AUTO_DISCOVER"
	// The bundling output directory includes a single file which will be used as the final bundle.
	//
	// If the output directory does not
	// include exactly a single file, bundling will fail.
	//
	// Similar to ARCHIVED but for non-archive files.
	BundlingOutput_SINGLE_FILE BundlingOutput = "SINGLE_FILE"
)

