package convert

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/charlie-haley/vyconfigure/pkg/api"
	"sigs.k8s.io/yaml"
)

// YamlToCmds
func YamlToCmds(config []byte, prefix string) ([]string, error) {
	var cmds []string

	j, _ := yaml.YAMLToJSON(config)

	var nestedMap map[string]interface{}
	err := json.Unmarshal(j, &nestedMap)
	if err != nil {
		return nil, err
	}

	err = mapToCmds(true, &cmds, nestedMap, prefix)
	if err != nil {
		return nil, err
	}

	return cmds, nil
}

// JsonToCmds
func JsonToCmds(config []byte, prefix string) ([]string, error) {
	var cmds []string

	var nestedMap map[string]interface{}
	err := json.Unmarshal(config, &nestedMap)
	if err != nil {
		return nil, err
	}

	err = mapToCmds(true, &cmds, nestedMap, prefix)
	if err != nil {
		return nil, err
	}

	return cmds, nil
}

// CmdsToData converts a list of commands into a format the HTTP API understands
func CmdsToData(cmds []string, op string) []api.Cmd {
	var res []api.Cmd
	for _, c := range cmds {
		cmd := api.Cmd{
			Operation: op,
			Path:      strings.Split(c, " "),
		}
		res = append(res, cmd)
	}

	return res
}

// mapToCmds maps an interface to an array of VyOS commands
func mapToCmds(top bool, cmds *[]string, nm interface{}, prefix string) error {
	assign := func(cmd string, v interface{}, array bool, ignoreValue bool) error {
		switch v := v.(type) {
		case map[string]interface{}, []interface{}:
			if err := mapToCmds(false, cmds, v, cmd); err != nil {
				return err
			}
		case string:
			if array || !ignoreValue {
				*cmds = append(*cmds, cmd+" "+v)
			} else {
				*cmds = append(*cmds, cmd)
			}
		default:
			*cmds = append(*cmds, cmd+" "+v.(string))
		}

		return nil
	}

	switch nm := nm.(type) {
	case map[string]interface{}:
		for k, v := range nm {
			cmd := buildCmd(top, prefix, k)

			// this is pretty ugly, basically when building the cmds we only care about the key if the value is {}
			r, _ := json.Marshal(v)
			res := string(r)

			if res == "{}" {
				if err := assign(cmd, k, false, true); err != nil {
					return err
				}
				continue
			}
			// again, very crude but check if we're looking at an array of string
			if strings.HasPrefix(res, "[") && strings.HasSuffix(res, "]") {
				for _, val := range v.([]interface{}) {
					if err := assign(cmd, val, true, false); err != nil {
						return err
					}
				}
				continue
			}

			if err := assign(cmd, v, false, false); err != nil {
				return err
			}
		}
	case []interface{}:
		for _, v := range nm {
			cmd := buildCmd(true, prefix, "")
			if err := assign(cmd, v, false, false); err != nil {
				return err
			}
		}
	default:
		return errors.New("invalid input, must be a map or slice of interface")
	}

	return nil
}

// buildCmd
func buildCmd(array bool, prefix, value string) string {
	if array {
		prefix += value
	} else {
		prefix += " " + value
	}

	return prefix
}
