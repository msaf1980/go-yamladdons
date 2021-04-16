package yamladdons

import (
	"time"

	"fmt"
)

// YDuration is a simple wrappper for time.Duration (for work with YAML decoder)
type YDuration time.Duration

func (t *YDuration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tm string
	if err := unmarshal(&tm); err != nil {
		return err
	}

	td, err := time.ParseDuration(tm)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to time.Duration: %v", tm, err)
	}

	*t = YDuration(td)
	return nil
}

func (t *YDuration) Duration() time.Duration {
	return time.Duration(*t)
}
