package awsapplicationsignals


// Identifies a code location to instrument.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeLocationProperty := &CodeLocationProperty{
//   	FilePath: jsii.String("filePath"),
//   	Language: jsii.String("language"),
//
//   	// the properties below are optional
//   	ClassName: jsii.String("className"),
//   	CodeUnit: jsii.String("codeUnit"),
//   	LineNumber: jsii.Number(123),
//   	MethodName: jsii.String("methodName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html
//
type CfnInstrumentationConfig_CodeLocationProperty struct {
	// The source file path relative to the project or source root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html#cfn-applicationsignals-instrumentationconfig-codelocation-filepath
	//
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The programming language for this instrumentation point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html#cfn-applicationsignals-instrumentationconfig-codelocation-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
	// The class or type name that contains the method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html#cfn-applicationsignals-instrumentationconfig-codelocation-classname
	//
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The package, module, or namespace that contains the target code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html#cfn-applicationsignals-instrumentationconfig-codelocation-codeunit
	//
	CodeUnit *string `field:"optional" json:"codeUnit" yaml:"codeUnit"`
	// The line number to instrument.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html#cfn-applicationsignals-instrumentationconfig-codelocation-linenumber
	//
	LineNumber *float64 `field:"optional" json:"lineNumber" yaml:"lineNumber"`
	// The method or function name to instrument.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-instrumentationconfig-codelocation.html#cfn-applicationsignals-instrumentationconfig-codelocation-methodname
	//
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
}

