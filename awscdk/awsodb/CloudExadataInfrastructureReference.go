package awsodb


// A reference to a CloudExadataInfrastructure resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudExadataInfrastructureReference := &CloudExadataInfrastructureReference{
//   	CloudExadataInfrastructureArn: jsii.String("cloudExadataInfrastructureArn"),
//   }
//
type CloudExadataInfrastructureReference struct {
	// The CloudExadataInfrastructureArn of the CloudExadataInfrastructure resource.
	CloudExadataInfrastructureArn *string `field:"required" json:"cloudExadataInfrastructureArn" yaml:"cloudExadataInfrastructureArn"`
}

