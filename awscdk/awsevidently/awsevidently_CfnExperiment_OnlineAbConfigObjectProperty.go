package awsevidently


// A structure that contains the configuration of which variation to use as the "control" version.
//
// The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineAbConfigObjectProperty := &onlineAbConfigObjectProperty{
//   	controlTreatmentName: jsii.String("controlTreatmentName"),
//   	treatmentWeights: []interface{}{
//   		&treatmentToWeightProperty{
//   			splitWeight: jsii.Number(123),
//   			treatment: jsii.String("treatment"),
//   		},
//   	},
//   }
//
type CfnExperiment_OnlineAbConfigObjectProperty struct {
	// The name of the variation that is to be the default variation that the other variations are compared to.
	ControlTreatmentName *string `field:"optional" json:"controlTreatmentName" yaml:"controlTreatmentName"`
	// A set of key-value pairs.
	//
	// The keys are treatment names, and the values are the portion of experiment traffic to be assigned to that treatment. Specify the traffic portion in thousandths of a percent, so 20,000 for a variation would allocate 20% of the experiment traffic to that variation.
	TreatmentWeights interface{} `field:"optional" json:"treatmentWeights" yaml:"treatmentWeights"`
}

