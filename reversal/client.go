// Package reversal provides the /transfers/reversals APIs
package reversal

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /transfers/reversals APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new transfer reversal.
func New(params *stripe.ReversalParams) (*stripe.Reversal, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.ReversalParams) (*stripe.Reversal, error) {
	path := stripe.FormatURLPath("/transfers/%s/reversals", stripe.StringValue(params.Transfer))
	reversal := &stripe.Reversal{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reversal)
	return reversal, err
}

// Get returns the details of a transfer reversal.
func Get(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Transfer must be set")
	}

	path := stripe.FormatURLPath("/transfers/%s/reversals/%s",
		stripe.StringValue(params.Transfer), id)
	reversal := &stripe.Reversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reversal)
	return reversal, err
}

// Update updates a transfer reversal's properties.
func Update(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	path := stripe.FormatURLPath("/transfers/%s/reversals/%s",
		stripe.StringValue(params.Transfer), id)
	reversal := &stripe.Reversal{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reversal)
	return reversal, err
}

// List returns a list of transfer reversals.
func List(params *stripe.ReversalListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.ReversalListParams) *Iter {
	path := stripe.FormatURLPath("/transfers/%s/reversals", stripe.StringValue(listParams.Transfer))

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ReversalList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Reversals.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Refund returns the most recent Reversals
// visited by a call to Next.
func (i *Iter) Reversal() *stripe.Reversal {
	return i.Current().(*stripe.Reversal)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
