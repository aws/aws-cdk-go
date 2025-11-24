package mixinsawsarczonalshift


// A control condition is an alarm that you specify for a practice run.
//
// When you configure practice runs with zonal autoshift for a resource, you specify Amazon CloudWatch alarms, which you create in CloudWatch to use with the practice run. The alarms that you specify are an *outcome alarm* , to monitor application health during practice runs and, optionally, a *blocking alarm* , to block practice runs from starting or to interrupt a practice run in progress.
//
// Control condition alarms do not apply for autoshifts.
//
// For more information, see [Considerations when you configure zonal autoshift](https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.considerations.html) in the ARC Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   controlConditionProperty := &ControlConditionProperty{
//   	AlarmIdentifier: jsii.String("alarmIdentifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-controlcondition.html
//
type CfnZonalAutoshiftConfigurationPropsMixin_ControlConditionProperty struct {
	// The Amazon Resource Name (ARN) for an Amazon CloudWatch alarm that you specify as a control condition for a practice run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-controlcondition.html#cfn-arczonalshift-zonalautoshiftconfiguration-controlcondition-alarmidentifier
	//
	AlarmIdentifier *string `field:"optional" json:"alarmIdentifier" yaml:"alarmIdentifier"`
	// The type of alarm specified for a practice run.
	//
	// You can only specify Amazon CloudWatch alarms for practice runs, so the only valid value is `CLOUDWATCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-controlcondition.html#cfn-arczonalshift-zonalautoshiftconfiguration-controlcondition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

