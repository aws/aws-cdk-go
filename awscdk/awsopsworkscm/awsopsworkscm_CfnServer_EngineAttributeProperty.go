package awsopsworkscm


// The `EngineAttribute` property type specifies administrator credentials for an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise server.
//
// `EngineAttribute` is a property of the `AWS::OpsWorksCM::Server` resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   engineAttributeProperty := &engineAttributeProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnServer_EngineAttributeProperty struct {
	// The name of the engine attribute.
	//
	// *Attribute name for Chef Automate servers:*
	//
	// - `CHEF_AUTOMATE_ADMIN_PASSWORD`
	//
	// *Attribute names for Puppet Enterprise servers:*
	//
	// - `PUPPET_ADMIN_PASSWORD`
	// - `PUPPET_R10K_REMOTE`
	// - `PUPPET_R10K_PRIVATE_KEY`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the engine attribute.
	//
	// *Attribute value for Chef Automate servers:*
	//
	// - `CHEF_AUTOMATE_PIVOTAL_KEY` : A base64-encoded RSA public key. The corresponding private key is required to access the Chef API. You can generate this key by running the following [OpenSSL](https://docs.aws.amazon.com/https://www.openssl.org/) command on Linux-based computers.
	//
	// `openssl genrsa -out *pivotal_key_file_name* .pem 2048`
	//
	// On Windows-based computers, you can use the PuTTYgen utility to generate a base64-encoded RSA private key. For more information, see [PuTTYgen - Key Generator for PuTTY on Windows](https://docs.aws.amazon.com/https://www.ssh.com/ssh/putty/windows/puttygen) on SSH.com.
	//
	// *Attribute values for Puppet Enterprise servers:*
	//
	// - `PUPPET_ADMIN_PASSWORD` : An administrator password that you can use to sign in to the Puppet Enterprise console webpage after the server is online. The password must use between 8 and 32 ASCII characters.
	// - `PUPPET_R10K_REMOTE` : The r10k remote is the URL of your control repository (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an r10k remote opens TCP port 8170.
	// - `PUPPET_R10K_PRIVATE_KEY` : If you are using a private Git repository, add `PUPPET_R10K_PRIVATE_KEY` to specify a PEM-encoded private SSH key.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

