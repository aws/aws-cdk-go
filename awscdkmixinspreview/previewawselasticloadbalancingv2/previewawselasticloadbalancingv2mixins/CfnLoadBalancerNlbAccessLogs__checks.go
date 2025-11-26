//go:build !no_runtime_type_checking

package previewawselasticloadbalancingv2mixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

func (c *jsiiProxy_CfnLoadBalancerNlbAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

