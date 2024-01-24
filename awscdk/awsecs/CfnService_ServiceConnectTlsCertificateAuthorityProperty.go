package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectTlsCertificateAuthorityProperty := &ServiceConnectTlsCertificateAuthorityProperty{
//   	AwsPcaAuthorityArn: jsii.String("awsPcaAuthorityArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttlscertificateauthority.html
//
type CfnService_ServiceConnectTlsCertificateAuthorityProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttlscertificateauthority.html#cfn-ecs-service-serviceconnecttlscertificateauthority-awspcaauthorityarn
	//
	AwsPcaAuthorityArn *string `field:"optional" json:"awsPcaAuthorityArn" yaml:"awsPcaAuthorityArn"`
}

