package domain

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBundle_UnmarshalJSON(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/bundle.json")
	require.NoError(t, err)

	var bundle Bundle
	err = json.Unmarshal(bytes, &bundle)
	require.NoError(t, err)
	assert.Equal(t, bundle.ID, "0c3151bd-1cbf-4d64-b04d-cd9187a4c6e0")
	assert.Equal(t, bundle.Type, BundleResType)
	assert.Len(t, bundle.Entries, 4)
}
