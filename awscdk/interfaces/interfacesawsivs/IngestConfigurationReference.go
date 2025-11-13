package interfacesawsivs


// A reference to a IngestConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestConfigurationReference := &IngestConfigurationReference{
//   	IngestConfigurationArn: jsii.String("ingestConfigurationArn"),
//   }
//
type IngestConfigurationReference struct {
	// The Arn of the IngestConfiguration resource.
	IngestConfigurationArn *string `field:"required" json:"ingestConfigurationArn" yaml:"ingestConfigurationArn"`
}

