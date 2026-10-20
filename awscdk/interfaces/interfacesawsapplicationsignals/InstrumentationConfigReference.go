package interfacesawsapplicationsignals


// A reference to a InstrumentationConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instrumentationConfigReference := &InstrumentationConfigReference{
//   	Environment: jsii.String("environment"),
//   	InstrumentationConfigArn: jsii.String("instrumentationConfigArn"),
//   	InstrumentationType: jsii.String("instrumentationType"),
//   	LocationHash: jsii.String("locationHash"),
//   	Service: jsii.String("service"),
//   	SignalType: jsii.String("signalType"),
//   }
//
type InstrumentationConfigReference struct {
	// The Environment of the InstrumentationConfig resource.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The ARN of the InstrumentationConfig resource.
	InstrumentationConfigArn *string `field:"required" json:"instrumentationConfigArn" yaml:"instrumentationConfigArn"`
	// The InstrumentationType of the InstrumentationConfig resource.
	InstrumentationType *string `field:"required" json:"instrumentationType" yaml:"instrumentationType"`
	// The LocationHash of the InstrumentationConfig resource.
	LocationHash *string `field:"required" json:"locationHash" yaml:"locationHash"`
	// The Service of the InstrumentationConfig resource.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The SignalType of the InstrumentationConfig resource.
	SignalType *string `field:"required" json:"signalType" yaml:"signalType"`
}

