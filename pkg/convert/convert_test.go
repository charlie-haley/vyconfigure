package convert

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCommands = []string{
		"test firewall ipv6-name WAN_IN default-action",
		"test firewall ipv6-name WAN_IN rule 10 state established",
		"test firewall ipv6-name WAN_IN rule 10 state related",
		"test firewall ipv6-name WAN_IN rule 10 action",
		"test firewall ipv6-name WAN_IN rule 20 action",
		"test firewall ipv6-name WAN_IN rule 20 protocol",
		"test firewall ipv6-name WAN_IN rule 30",
	}
)

// YamlToCmds
func Test_YamlToCmds(t *testing.T) {
	yaml := `
firewall:
  ipv6-name:
    WAN_IN:
      default-action: drop
      rule:
        "10":
            action: accept
            state:
              established: enable
              related: enable
        "20":
            action: accept
            protocol: ipv6-icmp
        "30": {}
    `

	res, err := YamlToCmds([]byte(yaml), "test ")
	assert.NoError(t, err)

	assert.ElementsMatch(t, res, testCommands)
}

func Test_JsonToCmds(t *testing.T) {
	json := `
    {
        "firewall": {
            "ipv6-name": {
                "WAN_IN": {
                    "default-action": "drop",
                    "rule": {
                        "10": {
                            "action": "accept",
                            "state": {
                                "established": "enable",
                                "related": "enable"
                            }
                        },
                        "20": {
                            "action": "accept",
                            "protocol": "ipv6-icmp"
                        },
                        "30": {}
                    }
                }
            }
        }
    }	
    `

	res, err := JsonToCmds([]byte(json), "test ")
	assert.NoError(t, err)

	assert.ElementsMatch(t, res, testCommands)
}

func Test_CmdsToData(t *testing.T) {
	operation := "delete"
	res := CmdsToData(testCommands, operation)

	assert.NotEmpty(t, res)

	for i, c := range res {
		expectedPath := strings.Split(testCommands[i], " ")
		assert.ElementsMatch(t, c.Path, expectedPath)

		assert.Equal(t, operation, c.Operation)
	}
}
