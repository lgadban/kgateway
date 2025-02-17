//go:build ignore

package assertions

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kgateway-dev/kgateway/v2/test/kubernetes/testutils/cluster"
	"github.com/kgateway-dev/kgateway/v2/test/kubernetes/testutils/kgateway"
)

// Provider is the entity that provides methods which assert behaviors of a Kubernetes Cluster
// These assertions occur against a running instance of kgateway, within a Kubernetes Cluster.
type Provider struct {
	Assert  *assert.Assertions
	Require *require.Assertions

	// Gomega is well-used around the codebase, so we also add support here
	// NOTE TO DEVELOPERS: We recommend relying on testify assertions where possible
	Gomega gomega.Gomega

	clusterContext  *cluster.Context
	kgatewayContext *kgateway.Context
}

// NewProvider returns a Provider that will provide Assertions that can be executed against an
// installation of kgateway
func NewProvider(t *testing.T) *Provider {
	gomega.RegisterTestingT(t)
	return &Provider{
		Assert:  assert.New(t),
		Require: require.New(t),
		Gomega:  gomega.NewWithT(t),

		clusterContext:  nil,
		kgatewayContext: nil,
	}
}

// WithClusterContext sets the provider to point to the provided cluster
func (p *Provider) WithClusterContext(clusterContext *cluster.Context) *Provider {
	p.clusterContext = clusterContext
	return p
}

// WithKgatewayContext sets the providers to point to a particular installation of kgateway
func (p *Provider) WithKgatewayContext(kgatewayCtx *kgateway.Context) *Provider {
	p.kgatewayContext = kgatewayCtx
	return p
}

// expectKgatewayContextDefined is invoked by methods on the Provider that can only be invoked
// if the provider has been configured to point to a kgateway installation
// There are certain Assertions that can be invoked that do not require that kgateway be installed for them to be invoked
func (p *Provider) expectKgatewayContextDefined() {
	p.Require.NotNil(p.kgatewayContext, "Provider attempted to create an Assertion that requires a kgateway installation, but none was configured")
}
