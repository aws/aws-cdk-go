package awssagemaker


// Settings for the capacity reservation for the compute instances that SageMaker AI reserves for an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   capacityReservationConfigProperty := &CapacityReservationConfigProperty{
//   	CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   	MlReservationArn: jsii.String("mlReservationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capacityreservationconfig.html
//
type CfnEndpointConfigPropsMixin_CapacityReservationConfigProperty struct {
	// Options that you can choose for the capacity reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capacityreservationconfig.html#cfn-sagemaker-endpointconfig-capacityreservationconfig-capacityreservationpreference
	//
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// The Amazon Resource Name (ARN) that uniquely identifies the ML capacity reservation that SageMaker AI applies when it deploys the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-capacityreservationconfig.html#cfn-sagemaker-endpointconfig-capacityreservationconfig-mlreservationarn
	//
	MlReservationArn *string `field:"optional" json:"mlReservationArn" yaml:"mlReservationArn"`
}

