package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automaticFailConfigurationProperty := &AutomaticFailConfigurationProperty{
//   	TargetSection: jsii.String("targetSection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-automaticfailconfiguration.html
//
type CfnEvaluationForm_AutomaticFailConfigurationProperty struct {
	// The target section refId to control failure propagation boundary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-automaticfailconfiguration.html#cfn-connect-evaluationform-automaticfailconfiguration-targetsection
	//
	TargetSection *string `field:"optional" json:"targetSection" yaml:"targetSection"`
}

