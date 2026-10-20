package awsapplicationsignals


// The location where instrumentation should be applied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   locationProperty := &LocationProperty{
//   	CodeLocation: &CodeLocationProperty{
//   		ClassName: jsii.String("className"),
//   		CodeUnit: jsii.String("codeUnit"),
//   		FilePath: jsii.String("filePath"),
//   		Language: jsii.String("language"),
//   		LineNumber: jsii.Number(123),
//   		MethodName: jsii.String("methodName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-location.html
//
type CfnInstrumentationConfigPropsMixin_LocationProperty struct {
	// Identifies a code location to instrument.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-location.html#cfn-applicationsignals-instrumentationconfig-location-codelocation
	//
	CodeLocation interface{} `field:"optional" json:"codeLocation" yaml:"codeLocation"`
}

