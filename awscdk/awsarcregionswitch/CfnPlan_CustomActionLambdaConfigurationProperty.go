package awsarcregionswitch


// Configuration for AWS Lambda functions that perform custom actions during a Region switch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionLambdaConfigurationProperty := &CustomActionLambdaConfigurationProperty{
//   	Lambdas: []interface{}{
//   		&LambdasProperty{
//   			Arn: jsii.String("arn"),
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   		},
//   	},
//   	RegionToRun: jsii.String("regionToRun"),
//   	RetryIntervalMinutes: jsii.Number(123),
//
//   	// the properties below are optional
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &LambdaUngracefulProperty{
//   		Behavior: jsii.String("behavior"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.html
//
type CfnPlan_CustomActionLambdaConfigurationProperty struct {
	// The AWS Lambda functions for the execution block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.html#cfn-arcregionswitch-plan-customactionlambdaconfiguration-lambdas
	//
	Lambdas interface{} `field:"required" json:"lambdas" yaml:"lambdas"`
	// The AWS Region for the function to run in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.html#cfn-arcregionswitch-plan-customactionlambdaconfiguration-regiontorun
	//
	RegionToRun *string `field:"required" json:"regionToRun" yaml:"regionToRun"`
	// The retry interval specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.html#cfn-arcregionswitch-plan-customactionlambdaconfiguration-retryintervalminutes
	//
	RetryIntervalMinutes *float64 `field:"required" json:"retryIntervalMinutes" yaml:"retryIntervalMinutes"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.html#cfn-arcregionswitch-plan-customactionlambdaconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.html#cfn-arcregionswitch-plan-customactionlambdaconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

