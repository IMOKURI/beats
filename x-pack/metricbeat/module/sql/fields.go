// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sql

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "sql", asset.ModuleFieldsPri, AssetSql); err != nil {
		panic(err)
	}
}

// AssetSql returns asset data.
// This is the base64 encoded gzipped contents of module/sql.
func AssetSql() string {
	return "eJykkk2O4jAUhPc5RYnlSHCALGY1yxESmgMMjl2AGycm9jPduX3L+UEQpZGaruWruOqL/dY4sysRW1cAYsWxxCq2blUAgY4qskRFUQVgGHWwF7G+KfG7AJDPofYmORbAwdKZWPbGGo2qOQVnSXdhiWPw6TJOFvKy9rF1exwo+sSImhKsjjgEX0Ph3+4vjBJVqUiMR+5777vbxNDdpksEWfO/nPQFXdY8eA5wD2GCvTI8WBPJmd27D2bmPenN+tPnIUUaiAc/qJMQcuKAtVmkmBP/EGKX46bunkN756hleq5lislsUs1g9ebXIpGv3qhlZg3D/8MXxqeq37hvIG+Hzts+jcA0z1GjBNscXyZ96Xa3vlmPV4SrcokPtJ8BAAD//zTX/E8="
}
