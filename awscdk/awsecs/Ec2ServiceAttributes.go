package awsecs


// The properties to import from the service using the EC2 launch type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   ec2ServiceAttributes := &Ec2ServiceAttributes{
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	ServiceArn: jsii.String("serviceArn"),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
type Ec2ServiceAttributes struct {
	// The cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The service ARN.
	// Default: - either this, or `serviceName`, is required.
	//
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	// Default: - either this, or `serviceArn`, is required.
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

