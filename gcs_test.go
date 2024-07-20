package cloudcourier

import "testing"

type FakeProviderConfig struct{}

func (f *FakeProviderConfig) GetProvider() Provider {
	return PROVIDER_UNSPECIFIED
}

func TestNewGCSProviderClient_WrongProviderConfigFails(t *testing.T) {
	config := FakeProviderConfig{}
	_, err := newGCSProviderClient(&config)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
