package awselasticbeanstalk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageSourceProperty := &ImageSourceProperty{
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagesource.html
//
type CfnApplicationVersion_ImageSourceProperty struct {
	// The URI of the container image, e.g. an ECR image URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-imagesource.html#cfn-elasticbeanstalk-applicationversion-imagesource-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

