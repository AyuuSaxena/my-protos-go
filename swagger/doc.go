package swagger

import (
	"embed"
	_ "embed"
)

// FS embeds all swagger JSON files in this directory.
//
//go:embed content/*
var FS embed.FS

// ContentRomServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for RomService.
//
//go:embed content/v1/rom/romService.swagger.json
var ContentRomServiceSwaggerJSON []byte

// IdentityServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for IdentityService.
//
//go:embed identity/v1/service.swagger.json
var IdentityServiceSwaggerJSON []byte
