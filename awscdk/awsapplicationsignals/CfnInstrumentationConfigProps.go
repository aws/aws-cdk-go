package awsapplicationsignals

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstrumentationConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstrumentationConfigProps := &CfnInstrumentationConfigProps{
//   	CaptureConfiguration: &CaptureConfigurationProperty{
//   		CodeCapture: &CodeCaptureConfigurationProperty{
//   			CaptureLimits: &CaptureLimitsConfigProperty{
//   				MaxCollectionDepth: jsii.Number(123),
//   				MaxCollectionWidth: jsii.Number(123),
//   				MaxFieldsPerObject: jsii.Number(123),
//   				MaxHits: jsii.Number(123),
//   				MaxObjectDepth: jsii.Number(123),
//   				MaxStackFrames: jsii.Number(123),
//   				MaxStackTraceSize: jsii.Number(123),
//   				MaxStringLength: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			CaptureArguments: []*string{
//   				jsii.String("captureArguments"),
//   			},
//   			CaptureLocals: []*string{
//   				jsii.String("captureLocals"),
//   			},
//   			CaptureReturn: jsii.Boolean(false),
//   			CaptureStackTrace: jsii.Boolean(false),
//   		},
//   	},
//   	Environment: jsii.String("environment"),
//   	InstrumentationType: jsii.String("instrumentationType"),
//   	Location: &LocationProperty{
//   		CodeLocation: &CodeLocationProperty{
//   			FilePath: jsii.String("filePath"),
//   			Language: jsii.String("language"),
//
//   			// the properties below are optional
//   			ClassName: jsii.String("className"),
//   			CodeUnit: jsii.String("codeUnit"),
//   			LineNumber: jsii.Number(123),
//   			MethodName: jsii.String("methodName"),
//   		},
//   	},
//   	Service: jsii.String("service"),
//   	SignalType: jsii.String("signalType"),
//
//   	// the properties below are optional
//   	AttributeFilters: []interface{}{
//   		map[string]*string{
//   			"attributeFiltersKey": jsii.String("attributeFilters"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExpiresAt: jsii.String("expiresAt"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html
//
type CfnInstrumentationConfigProps struct {
	// Specifies what to capture when the instrumentation point is hit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-captureconfiguration
	//
	CaptureConfiguration interface{} `field:"required" json:"captureConfiguration" yaml:"captureConfiguration"`
	// The environment that the service is running in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-environment
	//
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Type of instrumentation: BREAKPOINT (temporary, expires after 24 hours) or PROBE (permanent, persists until deleted).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-instrumentationtype
	//
	InstrumentationType *string `field:"required" json:"instrumentationType" yaml:"instrumentationType"`
	// The location where instrumentation should be applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-location
	//
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The name of the service to instrument.
	//
	// This should match the service.name resource attribute reported by the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-service
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// The telemetry signal type to emit for this instrumentation.
	//
	// The supported value is SNAPSHOT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-signaltype
	//
	SignalType *string `field:"required" json:"signalType" yaml:"signalType"`
	// Client-side filters that target specific instances.
	//
	// Each object is AND-matched on keys, multiple objects are OR-matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-attributefilters
	//
	AttributeFilters interface{} `field:"optional" json:"attributeFilters" yaml:"attributeFilters"`
	// An optional short description that explains the purpose of this instrumentation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The timestamp after which this configuration is no longer served.
	//
	// For BREAKPOINT only; defaults to 24 hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-expiresat
	//
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// An optional list of key-value pairs to associate with the instrumentation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-instrumentationconfig.html#cfn-applicationsignals-instrumentationconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

