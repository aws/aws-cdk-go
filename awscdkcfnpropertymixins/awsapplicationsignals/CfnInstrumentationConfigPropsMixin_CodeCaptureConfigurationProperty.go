package awsapplicationsignals


// Defines what data to capture for code-level instrumentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeCaptureConfigurationProperty := &CodeCaptureConfigurationProperty{
//   	CaptureArguments: []*string{
//   		jsii.String("captureArguments"),
//   	},
//   	CaptureLimits: &CaptureLimitsConfigProperty{
//   		MaxCollectionDepth: jsii.Number(123),
//   		MaxCollectionWidth: jsii.Number(123),
//   		MaxFieldsPerObject: jsii.Number(123),
//   		MaxHits: jsii.Number(123),
//   		MaxObjectDepth: jsii.Number(123),
//   		MaxStackFrames: jsii.Number(123),
//   		MaxStackTraceSize: jsii.Number(123),
//   		MaxStringLength: jsii.Number(123),
//   	},
//   	CaptureLocals: []*string{
//   		jsii.String("captureLocals"),
//   	},
//   	CaptureReturn: jsii.Boolean(false),
//   	CaptureStackTrace: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codecaptureconfiguration.html
//
type CfnInstrumentationConfigPropsMixin_CodeCaptureConfigurationProperty struct {
	// The function arguments to capture.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codecaptureconfiguration.html#cfn-applicationsignals-instrumentationconfig-codecaptureconfiguration-capturearguments
	//
	CaptureArguments *[]*string `field:"optional" json:"captureArguments" yaml:"captureArguments"`
	// Safety limits that bound what is captured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codecaptureconfiguration.html#cfn-applicationsignals-instrumentationconfig-codecaptureconfiguration-capturelimits
	//
	CaptureLimits interface{} `field:"optional" json:"captureLimits" yaml:"captureLimits"`
	// The local variables to capture by name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codecaptureconfiguration.html#cfn-applicationsignals-instrumentationconfig-codecaptureconfiguration-capturelocals
	//
	CaptureLocals *[]*string `field:"optional" json:"captureLocals" yaml:"captureLocals"`
	// Whether to capture the return value.
	//
	// Defaults to false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codecaptureconfiguration.html#cfn-applicationsignals-instrumentationconfig-codecaptureconfiguration-capturereturn
	//
	CaptureReturn interface{} `field:"optional" json:"captureReturn" yaml:"captureReturn"`
	// Whether to capture a stack trace.
	//
	// Defaults to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codecaptureconfiguration.html#cfn-applicationsignals-instrumentationconfig-codecaptureconfiguration-capturestacktrace
	//
	CaptureStackTrace interface{} `field:"optional" json:"captureStackTrace" yaml:"captureStackTrace"`
}

