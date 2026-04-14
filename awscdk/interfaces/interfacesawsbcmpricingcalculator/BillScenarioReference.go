package interfacesawsbcmpricingcalculator


// A reference to a BillScenario resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billScenarioReference := &BillScenarioReference{
//   	BillScenarioArn: jsii.String("billScenarioArn"),
//   }
//
type BillScenarioReference struct {
	// The Arn of the BillScenario resource.
	BillScenarioArn *string `field:"required" json:"billScenarioArn" yaml:"billScenarioArn"`
}

