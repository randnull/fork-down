package utils

import (
	"errors"
	"fork-down/custom_errors"
	"os"
	"strings"
)

func ValidateInput(filePath *string, manPath *string, oldMan *string) error {
	if *filePath == "" {
		return custom_errors.FileNotProvide
	}

	if *manPath == "" || *oldMan == "" {
		return custom_errors.ManifestNotProvide
	}

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		return custom_errors.FileNotFound
	}

	if _, err := os.Stat(*manPath); errors.Is(err, os.ErrNotExist) {
		return custom_errors.ManifestNotFound
	}

	if _, err := os.Stat(*oldMan); errors.Is(err, os.ErrNotExist) {
		return custom_errors.ManifestNotFound
	}

	if !strings.HasSuffix(*filePath, "bin") {
		return custom_errors.FileFormatError
	}

	//if !strings.HasSuffix(*manPath, "json") && !strings.HasSuffix(*manPath, "rdx") {
	//	return custom_errors.ManifestFormatError
	//}

	//if !strings.HasSuffix(*oldMan, "json") && !strings.HasSuffix(*oldMan, "rdx") {
	//	return custom_errors.ManifestFormatError
	//}

	return nil
}
