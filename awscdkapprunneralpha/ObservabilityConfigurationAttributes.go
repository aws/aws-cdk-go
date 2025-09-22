package awscdkapprunneralpha


// Attributes for the App Runner Observability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   observabilityConfigurationAttributes := &ObservabilityConfigurationAttributes{
//   	ObservabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	ObservabilityConfigurationRevision: jsii.Number(123),
//   }
//
// Experimental.
type ObservabilityConfigurationAttributes struct {
	// The name of the Observability configuration.
	// Experimental.
	ObservabilityConfigurationName *string `field:"required" json:"observabilityConfigurationName" yaml:"observabilityConfigurationName"`
	// The revision of the Observability configuration.
	// Experimental.
	ObservabilityConfigurationRevision *float64 `field:"required" json:"observabilityConfigurationRevision" yaml:"observabilityConfigurationRevision"`
}

