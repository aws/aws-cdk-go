package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiEventProperty := &ApiEventProperty{
//   	Method: jsii.String("method"),
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	RestApiId: jsii.String("restApiId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-apievent.html
//
type CfnStateMachine_ApiEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-apievent.html#cfn-serverless-statemachine-apievent-method
	//
	Method *string `field:"required" json:"method" yaml:"method"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-apievent.html#cfn-serverless-statemachine-apievent-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-apievent.html#cfn-serverless-statemachine-apievent-restapiid
	//
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
}

