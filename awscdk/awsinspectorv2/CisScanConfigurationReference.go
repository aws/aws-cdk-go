package awsinspectorv2


// A reference to a CisScanConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cisScanConfigurationReference := &CisScanConfigurationReference{
//   	CisScanConfigurationArn: jsii.String("cisScanConfigurationArn"),
//   }
//
type CisScanConfigurationReference struct {
	// The Arn of the CisScanConfiguration resource.
	CisScanConfigurationArn *string `field:"required" json:"cisScanConfigurationArn" yaml:"cisScanConfigurationArn"`
}

