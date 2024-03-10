package econfig

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gotomicro/ego/core/constant"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/eflag"
	"github.com/gotomicro/ego/core/elog"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Init 初始化配置
func Init() (err error) {
	var configAddr = eflag.String("config")
	if len(configAddr) <= 0 {
		configAddr = "./configs/server/local.toml"
	}
	return LoadConfig(configAddr)
}

// LoadConfig 启动导入配置，导入失败panic
func LoadConfig(configAddr string) error {
	ext := filepath.Ext(configAddr)
	if ext == "" { // 如果配置文件没有扩展名，尝试从环境变量获取配置文件的扩展名
		ext = os.Getenv(constant.EgoDefaultConfigExt)
	}
	var unmarshalFunc econf.Unmarshaller
	switch ext {
	case ".zj", ".json":
		unmarshalFunc = json.Unmarshal
	case ".zt", ".toml":
		unmarshalFunc = toml.Unmarshal
	case ".zy", ".yaml", ".yml":
		unmarshalFunc = yaml.Unmarshal
	default:
		elog.Error("data source: invalid configuration type", zap.String("ext", ext))
		return fmt.Errorf("data source: invalid ext: %s", ext)
	}

	file, err := os.Open(configAddr)
	if err != nil {
		elog.Error("open Config Error", zap.String("configAddr", configAddr), zap.Error(err))
		return err
	}

	err = econf.LoadFromReader(file, unmarshalFunc)
	if err != nil {
		elog.Error("load Config from file error", zap.String("configAddr", configAddr), zap.Error(err))
		return err
	}
	return nil
}
