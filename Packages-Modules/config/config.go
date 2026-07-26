// ============================================================
//  PACKAGE CONFIG — demonstrates init() for setup
// ============================================================
//  This package simulates loading configuration when the
//  application starts.
//
//  Python equivalent:
//    config/
//    ├── __init__.py
//    └── settings.py
//
//    # config/__init__.py
//    from .settings import Config
//    config = Config()
//    print("config loaded")  # ← runs at import time

package config

import "fmt"

// =============================================================================
// PACKAGE-LEVEL STATE (initialized in init())
// =============================================================================

// AppConfig holds the application configuration
type AppConfig struct {
	AppName string
	Port    int
	Debug   bool
}

// Config is the package-level config instance
// Python: config = Config() at module level in __init__.py
var Config AppConfig

// =============================================================================
// init() — Configuration Loader
// =============================================================================
// This runs AUTOMATICALLY when the config package is first imported.
// Go resolves init() execution in DEPENDENCY ORDER:
//   1. Dependencies' init() run first
//   2. Then the importing package's init()
//   3. Finally main()
//
// init() is perfect for:
//   - Loading config files / environment variables
//   - Opening database connections
//   - Registering things (HTTP handlers, DB drivers)
//   - Validating setup
//
// In Python, you'd do this at module level or in __init__.py:
//   config = Config()
//   config.load_from_env()
//
// The difference: Go init() has GUARANTEED ORDER and runs exactly ONCE.

func init() {
	fmt.Println("[config] loading configuration...")

	Config = AppConfig{
		AppName: "GoLan App",
		Port:    8080,
		Debug:   false,
	}

	fmt.Println("[config] app name:", Config.AppName)
	fmt.Println("[config] port:", Config.Port)
	fmt.Println("[config] debug mode:", Config.Debug)
	fmt.Println("[config] configuration loaded successfully")
}

// GetConfig returns the current configuration
// Python: just reference config directly
func GetConfig() AppConfig {
	return Config
}
