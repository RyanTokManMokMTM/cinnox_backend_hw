package config

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(fileName string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName(fileName)
	vp.SetConfigType("yml")
	vp.AddConfigPath("config/")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{
		vp: vp,
	}, nil
}

func (s *Setting) ReadSection(key string, v any) error {
	if err := s.vp.UnmarshalKey(key, v); err != nil {
		return err
	}
	return nil
}
