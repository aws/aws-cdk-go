package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityReservationConfigProperty := &CapacityReservationConfigProperty{
//   	CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   	MlReservationArn: jsii.String("mlReservationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capacityreservationconfig.html
//
type CfnEndpointConfig_CapacityReservationConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capacityreservationconfig.html#cfn-sagemaker-endpointconfig-capacityreservationconfig-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capacityreservationconfig.html#cfn-sagemaker-endpointconfig-capacityreservationconfig-mlreservationarn
	//
	MlReservationArn *string `field:"optional" json:"mlReservationArn" yaml:"mlReservationArn"`
}

