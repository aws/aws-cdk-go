package awsarcregionswitch


// The service for a cross account role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceProperty := &ServiceProperty{
//   	ClusterArn: jsii.String("clusterArn"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	ServiceArn: jsii.String("serviceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-service.html
//
type CfnPlan_ServiceProperty struct {
	// The cluster Amazon Resource Name (ARN) for a service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-service.html#cfn-arcregionswitch-plan-service-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The cross account role for a service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-service.html#cfn-arcregionswitch-plan-service-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The external ID (secret key) for the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-service.html#cfn-arcregionswitch-plan-service-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The Amazon Resource Name (ARN) for a service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-service.html#cfn-arcregionswitch-plan-service-servicearn
	//
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
}

