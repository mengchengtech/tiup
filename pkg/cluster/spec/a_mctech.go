// add by zhangbing

package spec

import perrs "github.com/pingcap/errors"

// Labels returns the labels of TIDB
func (s *TiDBSpec) Labels() (map[string]string, error) {
	lbs := make(map[string]string)

	if serverLabels := GetValueFromPath(s.Config, "labels"); serverLabels != nil {
		m := map[any]any{}
		if sm, ok := serverLabels.(map[string]any); ok {
			for k, v := range sm {
				m[k] = v
			}
		} else if im, ok := serverLabels.(map[any]any); ok {
			m = im
		}
		for k, v := range m {
			key, ok := k.(string)
			if !ok {
				return nil, perrs.Errorf("TiKV label name %v is not a string, check the instance: %s:%d", k, s.Host, s.GetMainPort())
			}
			value, ok := v.(string)
			if !ok {
				return nil, perrs.Errorf("TiKV label value %v is not a string, check the instance: %s:%d", v, s.Host, s.GetMainPort())
			}

			lbs[key] = value
		}
	}

	return lbs, nil
}

// Labels returns the labels of TiKV
func (s *TiFlashSpec) Labels() (map[string]string, error) {
	lbs := make(map[string]string)

	if serverLabels := GetValueFromPath(s.Config, "server.labels"); serverLabels != nil {
		m := map[any]any{}
		if sm, ok := serverLabels.(map[string]any); ok {
			for k, v := range sm {
				m[k] = v
			}
		} else if im, ok := serverLabels.(map[any]any); ok {
			m = im
		}
		for k, v := range m {
			key, ok := k.(string)
			if !ok {
				return nil, perrs.Errorf("TiKV label name %v is not a string, check the instance: %s:%d", k, s.Host, s.GetMainPort())
			}
			value, ok := v.(string)
			if !ok {
				return nil, perrs.Errorf("TiKV label value %v is not a string, check the instance: %s:%d", v, s.Host, s.GetMainPort())
			}

			lbs[key] = value
		}
	}

	return lbs, nil
}
