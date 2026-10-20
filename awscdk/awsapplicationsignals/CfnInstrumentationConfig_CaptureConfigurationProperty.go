package awsapplicationsignals


// Specifies what to capture when the instrumentation point is hit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureConfigurationProperty := &CaptureConfigurationProperty{
//   	CodeCapture: &CodeCaptureConfigurationProperty{
//   		CaptureLimits: &CaptureLimitsConfigProperty{
//   			MaxCollectionDepth: jsii.Number(123),
//   			MaxCollectionWidth: jsii.Number(123),
//   			MaxFieldsPerObject: jsii.Number(123),
//   			MaxHits: jsii.Number(123),
//   			MaxObjectDepth: jsii.Number(123),
//   			MaxStackFrames: jsii.Number(123),
//   			MaxStackTraceSize: jsii.Number(123),
//   			MaxStringLength: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		CaptureArguments: []*string{
//   			jsii.String("captureArguments"),
//   		},
//   		CaptureLocals: []*string{
//   			jsii.String("captureLocals"),
//   		},
//   		CaptureReturn: jsii.Boolean(false),
//   		CaptureStackTrace: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-captureconfiguration.html
//
type CfnInstrumentationConfig_CaptureConfigurationProperty struct {
	// Defines what data to capture for code-level instrumentation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-captureconfiguration.html#cfn-applicationsignals-instrumentationconfig-captureconfiguration-codecapture
	//
	CodeCapture interface{} `field:"required" json:"codeCapture" yaml:"codeCapture"`
}

