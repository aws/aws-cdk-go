package interfacesawsdevopsguru


// A reference to a LogAnomalyDetectionIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logAnomalyDetectionIntegrationReference := &LogAnomalyDetectionIntegrationReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type LogAnomalyDetectionIntegrationReference struct {
	// The AccountId of the LogAnomalyDetectionIntegration resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

