package awsapplicationsignals


// Safety limits that bound what is captured.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captureLimitsConfigProperty := &CaptureLimitsConfigProperty{
//   	MaxCollectionDepth: jsii.Number(123),
//   	MaxCollectionWidth: jsii.Number(123),
//   	MaxFieldsPerObject: jsii.Number(123),
//   	MaxHits: jsii.Number(123),
//   	MaxObjectDepth: jsii.Number(123),
//   	MaxStackFrames: jsii.Number(123),
//   	MaxStackTraceSize: jsii.Number(123),
//   	MaxStringLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html
//
type CfnInstrumentationConfig_CaptureLimitsConfigProperty struct {
	// Maximum nesting depth to traverse inside collections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxcollectiondepth
	//
	MaxCollectionDepth *float64 `field:"optional" json:"maxCollectionDepth" yaml:"maxCollectionDepth"`
	// Maximum number of items to capture from any collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxcollectionwidth
	//
	MaxCollectionWidth *float64 `field:"optional" json:"maxCollectionWidth" yaml:"maxCollectionWidth"`
	// Maximum number of fields to capture for any object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxfieldsperobject
	//
	MaxFieldsPerObject *float64 `field:"optional" json:"maxFieldsPerObject" yaml:"maxFieldsPerObject"`
	// Maximum number of times the instrumentation point can be hit before disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxhits
	//
	MaxHits *float64 `field:"optional" json:"maxHits" yaml:"maxHits"`
	// Maximum depth for nested object traversal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxobjectdepth
	//
	MaxObjectDepth *float64 `field:"optional" json:"maxObjectDepth" yaml:"maxObjectDepth"`
	// Maximum number of stack frames to capture.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxstackframes
	//
	MaxStackFrames *float64 `field:"optional" json:"maxStackFrames" yaml:"maxStackFrames"`
	// Maximum total size in bytes of a captured stack trace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxstacktracesize
	//
	MaxStackTraceSize *float64 `field:"optional" json:"maxStackTraceSize" yaml:"maxStackTraceSize"`
	// Maximum length of captured string values in characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-capturelimitsconfig.html#cfn-applicationsignals-instrumentationconfig-capturelimitsconfig-maxstringlength
	//
	MaxStringLength *float64 `field:"optional" json:"maxStringLength" yaml:"maxStringLength"`
}

