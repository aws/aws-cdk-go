package awscodebuild


// Environment type for Windows Docker images.
//
// Example:
//   var ecrRepository repository
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.WindowsBuildImage_FromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.WindowsImageType_SERVER_2019),
//   		// optional certificate to include in the build image
//   		Certificate: &BuildEnvironmentCertificate{
//   			Bucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
//   			ObjectKey: jsii.String("path/to/cert.pem"),
//   		},
//   	},
//   })
//
type WindowsImageType string

const (
	// The standard environment type, WINDOWS_CONTAINER.
	WindowsImageType_STANDARD WindowsImageType = "STANDARD"
	// The WINDOWS_SERVER_2019_CONTAINER environment type.
	WindowsImageType_SERVER_2019 WindowsImageType = "SERVER_2019"
)

