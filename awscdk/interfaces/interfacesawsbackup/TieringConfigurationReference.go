package interfacesawsbackup


// A reference to a TieringConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tieringConfigurationReference := &TieringConfigurationReference{
//   	TieringConfigurationArn: jsii.String("tieringConfigurationArn"),
//   	TieringConfigurationName: jsii.String("tieringConfigurationName"),
//   }
//
type TieringConfigurationReference struct {
	// The ARN of the TieringConfiguration resource.
	TieringConfigurationArn *string `field:"required" json:"tieringConfigurationArn" yaml:"tieringConfigurationArn"`
	// The TieringConfigurationName of the TieringConfiguration resource.
	TieringConfigurationName *string `field:"required" json:"tieringConfigurationName" yaml:"tieringConfigurationName"`
}

