package awsemrcontainers


// Container provider information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerProviderProperty := &ContainerProviderProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Info: &ContainerInfoProperty{
//   		EksInfo: &EksInfoProperty{
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html
//
type CfnSecurityConfiguration_ContainerProviderProperty struct {
	// The container provider ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html#cfn-emrcontainers-securityconfiguration-containerprovider-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The container provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html#cfn-emrcontainers-securityconfiguration-containerprovider-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Container information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html#cfn-emrcontainers-securityconfiguration-containerprovider-info
	//
	Info interface{} `field:"optional" json:"info" yaml:"info"`
}

