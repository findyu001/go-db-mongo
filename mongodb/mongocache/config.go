package mongocache

import (
  "crypto/sha1"
  "encoding/hex"
)

type Config struct {
  URI      string `conf:"uri"`
  User     string `conf:"user"`
  Password string `conf:"passwd"`
  MaxConn  uint64 `conf:"maxconn"`
}

func (c *Config) id() string {
  sha := sha1.New()
  sha.Write([]byte(c.User))
  sha.Write([]byte(c.Password))
  sha.Write([]byte(c.URI))
  return hex.EncodeToString(sha.Sum([]byte{}))
}
