package awsecs


// The properties to import from the service using the Fargate launch type.
//
// Example:
//   var cluster Cluster
//
//
//   // Import service from EC2 service attributes
//   ec2ServiceFromAttributes := ecs.Ec2Service_FromEc2ServiceAttributes(this, jsii.String("Ec2ServiceFromAttributes"), &Ec2ServiceAttributes{
//   	ServiceArn: jsii.String("arn:aws:ecs:us-west-2:123456789012:service/my-http-service"),
//   	Cluster: Cluster,
//   })
//
//   // Import service from EC2 service ARN
//   ec2ServiceFromArn := ecs.Ec2Service_FromEc2ServiceArn(this, jsii.String("Ec2ServiceFromArn"), jsii.String("arn:aws:ecs:us-west-2:123456789012:service/my-http-service"))
//
//   // Import service from Fargate service attributes
//   fargateServiceFromAttributes := ecs.FargateService_FromFargateServiceAttributes(this, jsii.String("FargateServiceFromAttributes"), &FargateServiceAttributes{
//   	ServiceArn: jsii.String("arn:aws:ecs:us-west-2:123456789012:service/my-http-service"),
//   	Cluster: Cluster,
//   })
//
//   // Import service from Fargate service ARN
//   fargateServiceFromArn := ecs.FargateService_FromFargateServiceArn(this, jsii.String("FargateServiceFromArn"), jsii.String("arn:aws:ecs:us-west-2:123456789012:service/my-http-service"))
//
type FargateServiceAttributes struct {
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

