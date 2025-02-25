package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/huaweicloud"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`huaweicloud {
			access_key_id thekey
			secret_access_key itsasecret
			project_id 123456
			region cn-south-1
		}`,`huaweicloud {
			access_key_id thekey1
			secret_access_key itsasecret1
			project_id 789012
			region cn-east-3
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&huaweicloud.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedAccessKeyId := "thekey"
			actualAccessKeyId := p.Provider.AccessKeyId
			if expectedAccessKeyId != actualAccessKeyId {
				t.Errorf("Expected AccessKeyId to be '%s' but got '%s'", expectedAccessKeyId, actualAccessKeyId)
			}

			expectedSecretAccessKey := "itsasecret"
			actualSecretAccessKey := p.Provider.SecretAccessKey
			if expectedSecretAccessKey != actualSecretAccessKey {
				t.Errorf("Expected SecretAccessKey to be '%s' but got '%s'", expectedSecretAccessKey, actualSecretAccessKey)
			}

			expectedProjectId := "123456"
			actualProjectId := p.Provider.ProjectID
			if expectedProjectId != actualProjectId {
				t.Errorf("Expected ProjectId to be '%s' but got '%s'", expectedProjectId, actualProjectId)
			}

			expectedRegion := "cn-south-1"
			actualRegion := p.Provider.Region
			if expectedRegion != actualRegion {
				t.Errorf("Expected Region to be '%s' but got '%s'", expectedRegion, actualRegion)
			}
		})
	}
}