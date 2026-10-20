package awsapplicationsignals


// The location where instrumentation should be applied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &LocationProperty{
//   	CodeLocation: &CodeLocationProperty{
//   		FilePath: jsii.String("filePath"),
//   		Language: jsii.String("language"),
//
//   		// the properties below are optional
//   		ClassName: jsii.String("className"),
//   		CodeUnit: jsii.String("codeUnit"),
//   		LineNumber: jsii.Number(123),
//   		MethodName: jsii.String("methodName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-location.html
//
type CfnInstrumentationConfig_LocationProperty struct {
	// Identifies a code location to instrument.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-location.html#cfn-applicationsignals-instrumentationconfig-location-codelocation
	//
	CodeLocation interface{} `field:"required" json:"codeLocation" yaml:"codeLocation"`
}

