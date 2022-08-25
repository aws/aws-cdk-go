package awsstepfunctionstasks


// Properties for the EMR Cluster Applications.
//
// Applies to Amazon EMR releases 4.0 and later. A case-insensitive list of applications for Amazon EMR to install and configure when launching
// the cluster.
//
// See the RunJobFlow API for complete documentation on input parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationConfigProperty := &applicationConfigProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	additionalInfo: map[string]*string{
//   		"additionalInfoKey": jsii.String("additionalInfo"),
//   	},
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	version: jsii.String("version"),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_Application.html
//
type EmrCreateCluster_ApplicationConfigProperty struct {
	// The name of the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// This option is for advanced users only.
	//
	// This is meta information about third-party applications that third-party vendors use
	// for testing purposes.
	AdditionalInfo *map[string]*string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Arguments for Amazon EMR to pass to the application.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The version of the application.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

